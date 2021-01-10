package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if sec.IsAuthenticated(w, r) {
		http.Redirect(w, r, route.EntidadesRoute, 200)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
	log.Println("IndexHandler")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout Handler")
	session, _ := sec.Store.Get(r, sec.CookieName)
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	var page mdl.PageLogin
	var msg = r.FormValue("msg")
	if msg != "" {
		page.Msg = msg
	}
	var errMsg = r.FormValue("errMsg")
	if errMsg != "" {
		page.ErrMsg = errMsg
	}
	var tmpl = template.Must(template.ParseFiles("tiles/login.html"))
	tmpl.ExecuteTemplate(w, "Login", page)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Register User Handler")
	session, _ := sec.Store.Get(r, sec.CookieName)
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.ServeFile(w, r, "tiles/users/Register-User.html")
}

func RecoverUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recover User Password Handler")
	session, _ := sec.Store.Get(r, sec.CookieName)
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.ServeFile(w, r, "tiles/users/Recover-Passwd.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var page mdl.PageLogin
	if r.Method != "POST" {
		var tmpl = template.Must(template.ParseFiles("tiles/login.html"))
		tmpl.ExecuteTemplate(w, "Login", page)
		return
	}
	username := r.FormValue("usrname")
	log.Println(username)
	password := r.FormValue("psw")
	var user mdl.User
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Println(string(bytes))
	sql := "SELECT u.id, " +
		" u.name, " +
		" u.username, " +
		" u.password, " +
		" COALESCE(u.role_id, 0), " +
		" coalesce(r.name,'') as role_name " +
		" FROM users u " +
		" LEFT JOIN roles r ON u.role_id = r.id " +
		" WHERE username='" + username + "'"
	log.Println(sql)
	Db.QueryRow(sql).Scan(
		&user.Id, &user.Name,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.RoleName)
	// validate password
	log.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err.Error())
		page.ErrMsg = "Usuário/Senha inválido."
		var tmpl = template.Must(template.ParseFiles("tiles/login.html"))
		tmpl.ExecuteTemplate(w, "Login", page)
		return
	} else {
		AddUserInCookie(w, r, user)
		// Abrindo o Cookie
		savedUser := GetUserInCookie(w, r)
		log.Println("MAIN Saved User is " + savedUser.Username)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	}
}

func GetUserInCookie(w http.ResponseWriter, r *http.Request) mdl.User {
	var savedUser mdl.User
	session, _ := sec.Store.Get(r, sec.CookieName)
	sessionUser := session.Values["user"]
	if sessionUser != nil {
		strUser := sessionUser.(string)
		json.Unmarshal([]byte(strUser), &savedUser)
	}
	log.Println("Saved User is " + savedUser.Name)
	return savedUser
}

func AddUserInCookie(w http.ResponseWriter, r *http.Request, user mdl.User) {
	sec.Store.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session, _ := sec.Store.Get(r, sec.CookieName)
	bytesUser, _ := json.Marshal(&user)
	session.Values["user"] = string(bytesUser)
	sec.Store.Save(r, w, session)
	session.Save(r, w)
}

func BuildLoggedUser(user mdl.User) mdl.LoggedUser {
	var loggedUser mdl.LoggedUser
	loggedUser.User = user
	loggedUser.HasPermission = func(permission string) bool {
		//log.Println(permission)
		query := "SELECT " +
			"A.feature_id, B.code FROM features_roles A, features B " +
			"WHERE A.feature_id = B.id AND A.role_id = $1"
		rows, _ := Db.Query(query, user.Role)
		defer rows.Close()
		var features []mdl.Feature
		var feature mdl.Feature
		for rows.Next() {
			rows.Scan(&feature.Id, &feature.Code)
			features = append(features, feature)
		}
		for _, value := range features {
			//log.Println(value.Code)
			if value.Code == permission {
				//log.Println(permission + " encontrada!!!")
				return true
			}
		}
		return false
	}
	return loggedUser
}

func HasPermission(currentUser mdl.User, permission string) bool {
	log.Println("HasPermission")
	query := "SELECT " +
		"A.feature_id, B.code FROM features_roles A, features B " +
		"WHERE A.feature_id = B.id AND A.role_id = $1"
	rows, _ := Db.Query(query, currentUser.Role)
	defer rows.Close()
	var featuresCurrentUser []mdl.Feature
	var feature mdl.Feature
	for rows.Next() {
		rows.Scan(&feature.Id, &feature.Code)
		featuresCurrentUser = append(featuresCurrentUser, feature)
	}
	for _, value := range featuresCurrentUser {
		//log.Println("Feature CurUser: " + value.Code)
		if value.Code == permission {
			//log.Println("Permission: " + permission)
			//log.Println("PASSEI NO VESTIBULAR")
			return true
		}
	}
	//log.Println("NÃO PASSEI ")
	return false
}
