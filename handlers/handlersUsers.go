package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create User")
	if r.Method == "POST" {
		currentUser := GetUserInCookie(w, r)
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		role := r.FormValue("RoleForInsert")
		escritorio := r.FormValue("EscritorioForInsert")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		sqlStatement := "INSERT INTO Users(name, username, password, email, mobile, role_id, escritorio_id, author_id, criado_em) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
		id := 0
		err = Db.QueryRow(sqlStatement, name, username, hash, email, mobile, role, escritorio, currentUser.Id, time.Now()).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) +
			" | Name: " + name + " | Username: " + username +
			" | Password: " + password + " | Email: " + email +
			" | Mobile: " + mobile + " | Role: " + role + " | Escritório: " + escritorio)
	}
	http.Redirect(w, r, route.UsersRoute, 301)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update User")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		role := r.FormValue("RoleForUpdate")
		escritorio := r.FormValue("EscritorioForUpdate")
		if escritorio == "" {
			escritorio = "0"
		}
		log.Println("Role: " + role)
		sqlStatement := " UPDATE Users SET name=$1, " +
			" username=$2, email=$3, mobile=$4, role_id=$5 , escritorio_id=$6 " +
			" WHERE id=$7 "
		log.Println(sqlStatement)
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(name, username, email, mobile, role, escritorio, id)
		log.Println("UPDATE: Id: " +
			id + " | Name: " +
			name + " | Username: " +
			username + " | E-mail: " +
			email + " | Mobile: " +
			mobile + " | Role: " +
			role + " | Escritorio: " +
			escritorio)
		http.Redirect(w, r, route.UsersRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete User")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Users WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.UsersRoute, 301)
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Users")
	if sec.IsAuthenticated(w, r) {
		sql := "SELECT " +
			" a.id, a.name, a.username, a.password, " +
			" a.email, a.mobile, COALESCE(a.role_id, 0), COALESCE(b.name,'') as role_name, " +
			" COALESCE(f.id, 0), COALESCE(f.abreviatura,'') as escritorio_nome, " +
			" a.author_id, " +
			" e.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(d.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM users a " +
			" LEFT JOIN roles b ON a.role_id = b.id " +
			" LEFT JOIN membros c ON a.id = c.usuario_id " +
			" LEFT JOIN escritorios f ON c.escritorio_id = f.id " +
			" LEFT JOIN status d ON a.status_id = c.id " +
			" LEFT JOIN users e ON a.author_id = e.id " +
			" ORDER BY a.name ASC "
		log.Println("SQL: " + sql)
		rows, _ := Db.Query(sql)
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
				&user.Escritorio,
				&user.EscritorioNome,
				&user.AuthorId,
				&user.AuthorName,
				&user.C_CriadoEm,
				&user.CStatus,
				&user.StatusId,
				&user.IdVersaoOrigem)
			user.Order = i
			i++
			users = append(users, user)
		}
		sql = "SELECT id, name FROM roles ORDER BY name asc"
		log.Println("SQL Roles: " + sql)
		rows, _ = Db.Query(sql)
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
		sql = "SELECT id, nome FROM escritorios ORDER BY nome asc"
		log.Println("SQL Escritorios: " + sql)
		rows, _ = Db.Query(sql)
		var escritorios []mdl.Escritorio
		var escritorio mdl.Escritorio
		i = 1
		for rows.Next() {
			rows.Scan(&escritorio.Id,
				&escritorio.Nome)
			escritorio.Order = i
			i++
			escritorios = append(escritorios, escritorio)
		}
		var page mdl.PageUsers
		page.Users = users
		page.Roles = roles
		page.Escritorios = escritorios
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
