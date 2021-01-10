package handlers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User")
	log.Println(r.Method)
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		name := r.FormValue("Name")
		log.Println(name)
		username := r.FormValue("Username")
		log.Println(username)
		password := r.FormValue("Password")
		log.Println(password)
		email := r.FormValue("Email")
		log.Println(email)
		mobile := r.FormValue("Mobile")
		log.Println(mobile)
		role := r.FormValue("Role")
		log.Println(role)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		log.Println(hash)
		statusUserId := GetStartStatus("user")
		sqlStatement := "INSERT INTO Users(name, username, password, email, mobile, role_id, author_id, criado_em, status_id) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
		id := 0
		err = Db.QueryRow(sqlStatement, name, username, hash, email, mobile, role, currentUser.Id, time.Now(), statusUserId).Scan(&id)
		if err != nil {
			log.Println(err.Error())
			errMsg := "Erro ao criar Usuário."
			if role == "" {
				errMsg = errMsg + " Faltou informar o Perfil do Usuário."
			}
			http.Redirect(w, r, route.UsersRoute+"?errMsg="+errMsg, 301)
		} else {
			log.Println("INSERT: Id: " + strconv.Itoa(id) +
				" | Name: " + name + " | Username: " + username +
				" | Password: " + password + " | Email: " + email +
				" | Mobile: " + mobile + " | Role: " + role)
			http.Redirect(w, r, route.UsersRoute+"?msg=Usuário criado com sucesso.", 301)
		}
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update User")
	log.Println(r.Method)
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		role := r.FormValue("RoleForUpdate")
		log.Println("Role: " + role)
		sqlStatement := " UPDATE Users SET name=$1, " +
			" username=$2, email=$3, mobile=$4, role_id=$5 " +
			" WHERE id=$6 "
		log.Println(sqlStatement)
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		updtForm.Exec(name, username, email, mobile, role, id)
		log.Println("UPDATE: Id: " +
			id + " | Name: " +
			name + " | Username: " +
			username + " | E-mail: " +
			email + " | Mobile: " +
			mobile + " | Role: " +
			role)
		http.Redirect(w, r, route.UsersRoute+"?msg=Usuário atualizado com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete User")
	log.Println(r.Method)
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		errMsg := "Usuário vinculado a registro não pode ser removido."
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Users WHERE id=$1"
		deleteForm, _ := Db.Prepare(sqlStatement)
		_, err := deleteForm.Exec(id)
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			http.Redirect(w, r, route.UsersRoute+"?errMsg="+errMsg, 301)
		} else {
			http.Redirect(w, r, route.UsersRoute+"?msg=Usuário removido com sucesso.", 301)
		}
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func SendPasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Send Password")
	log.Println(r.Method)
	// Sender data.
	from := "virtusimpavidas@gmail.com"
	emailPassword := "v1rtudesemmedo"
	subject := "Sistema Virtus"
	emailTo := r.FormValue("Email")
	// Receiver email address.
	to := []string{
		emailTo,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	log.Println(password)
	user := loadUserByEmail(emailTo)
	if user.Id == 0 {
		errMsg := "?errMsg=Email não encontrado!"
		log.Println(errMsg)
		http.Redirect(w, r, "/logout"+errMsg, 301)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	atualizarSenha(hash, strconv.FormatInt(user.Id, 10))
	// Message.
	message := []byte("To: " + emailTo + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		"Username: " + user.Username + "\r\n" +
		"Nova senha: " + password + "\r\n")
	log.Println(message)

	// Authentication.
	auth := smtp.PlainAuth("", from, emailPassword, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "?msg=Email enviado com Sucesso!"
	log.Println(msg)
	http.Redirect(w, r, "/logout"+msg, 301)
}

func loadUserByEmail(emailTo string) mdl.User {
	rows, _ := Db.Query("SELECT id, username FROM users WHERE email = $1", emailTo)
	var user mdl.User
	if rows.Next() {
		rows.Scan(&user.Id, &user.Username)
	}
	return user
}

func RecoverPasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recover Password")
	log.Println(r.Method)
	http.ServeFile(w, r, "tiles/users/Recover-Passwd.html")
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Change Password")
	log.Println(r.Method)
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		//		currentUser := GetUserInCookie(w, r)
		id := r.FormValue("Id")
		log.Println(id)
		username := r.FormValue("Username")
		log.Println(username)
		password := r.FormValue("Password")
		log.Println(password)
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		atualizarSenha(hash, id)
		http.Redirect(w, r, route.UsersRoute+"?msg=Nova senha atualizada com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func atualizarSenha(hash []byte, id string) {
	sqlStatement := "UPDATE Users SET password = '" + string(hash[:]) +
		"' WHERE id = " + id
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func RegisterNewUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Register New User")
	log.Println(r.Method)
	name := r.FormValue("Name")
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	//	log.Println(password)
	email := r.FormValue("Email")
	mobile := r.FormValue("Mobile")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//	log.Println(hash)
	sqlStatement := "INSERT INTO Users(name, " +
		" username, " +
		" password, " +
		" email, " +
		" mobile, " +
		" role_id, " +
		" criado_em) " +
		" VALUES ( '" + name + "', " +
		" '" + username + "', " +
		" '" + string(hash[:]) + "', " +
		" '" + email + "', " +
		" '" + mobile + "', " +
		" 1, " +
		" now()::timestamp ) " +
		" RETURNING id"
	log.Println(sqlStatement)
	id := 0
	err = Db.QueryRow(sqlStatement).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		errMsg := "Erro ao criar Usuário."
		http.Redirect(w, r, route.UsersRoute+"?errMsg="+errMsg, 301)
	} else {
		log.Println("INSERT: Id: " + strconv.Itoa(id) +
			" | Name: " + name + " | Username: " + username +
			" | Password: " + password + " | Email: " + email +
			" | Mobile: " + mobile)
		msg := "?msg=Seu registro foi efetuado com sucesso.\n Após nossa análise, você receberá um e-mail de confirmação."
		log.Println(msg)
		http.Redirect(w, r, "/logout"+msg, 301)
	}
}

func SignUpUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign-Up Users")
	log.Println(r.Method)
	http.ServeFile(w, r, "tiles/users/Sign-Up-User.html")
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Users")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listUsers") {
		msg := r.FormValue("msg")
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, a.name, a.username, a.password, " +
			" a.email, a.mobile, " +
			" COALESCE(a.role_id, 0), COALESCE(b.name,'') as role_name, " +
			" a.author_id, " +
			" e.name as author_name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(d.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM users a " +
			" LEFT JOIN roles b ON a.role_id = b.id " +
			" LEFT JOIN status d ON a.status_id = d.id " +
			" LEFT JOIN users e ON a.author_id = e.id " +
			" ORDER BY a.name ASC "
		log.Println("SQL: " + sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var users []mdl.User
		var user mdl.User
		var i = 1
		for rows.Next() {
			rows.Scan(&user.Id,
				&user.Name,
				&user.Username,
				&user.Password,
				&user.Email,
				&user.Mobile,
				&user.Role,
				&user.RoleName,
				&user.AuthorId,
				&user.AuthorName,
				&user.C_CriadoEm,
				&user.CStatus,
				&user.StatusId,
				&user.IdVersaoOrigem)
			user.Order = i
			i++
			log.Println(user)
			users = append(users, user)
		}
		sql = "SELECT id, name FROM roles ORDER BY name asc"
		log.Println("SQL Roles: " + sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var roles []mdl.Role
		var role mdl.Role
		i = 1
		for rows.Next() {
			rows.Scan(&role.Id,
				&role.Name)
			role.Order = i
			i++
			roles = append(roles, role)
		}
		var page mdl.PageUsers
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		if msg != "" {
			page.Msg = msg
		}
		page.Users = users
		page.Roles = roles
		page.AppName = mdl.AppName
		page.Title = "Usuários"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/users/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Users", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
