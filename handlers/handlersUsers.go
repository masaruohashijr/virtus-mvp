package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create User")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		role := r.FormValue("RoleForInsert")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		sqlStatement := "INSERT INTO Users(name, username, password, email, mobile, role_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
		id := 0
		err = Db.QueryRow(sqlStatement, name, username, hash, email, mobile, role).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) +
			" | Name: " + name + " | Username: " + username +
			" | Password: " + password + " | Email: " + email +
			" | Mobile: " + mobile + " | Role: " + role)
	}
	http.Redirect(w, r, route.UsersRoute, 301)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update User")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		role := r.FormValue("RoleForUpdate")
		log.Println("Role: " + role)
		sqlStatement := "UPDATE Users SET name=$1, " +
			"username=$2, email=$3, mobile=$4, role_id=$5 " +
			"WHERE id=$6"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, username, email, mobile, role, id)
		log.Println("UPDATE: Id: " +
			id + " | Name: " +
			name + " | Username: " +
			username + " | E-mail: " +
			email + " | Mobile: " +
			mobile + " | Role: " +
			role)
	}
	http.Redirect(w, r, route.UsersRoute, 301)
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
	sec.IsAuthenticated(w, r)
	query := "SELECT " +
		"a.id, a.name, a.username, a.password, " +
		"a.email, a.mobile, COALESCE(a.role_id, 0), COALESCE(b.name,'') " +
		"FROM users a LEFT JOIN roles b ON a.role_id = b.id"
	log.Println("Query: " + query)
	rows, err := Db.Query(query)
	sec.CheckInternalServerError(err, w)
	var users []mdl.User
	var user mdl.User
	var i = 1
	for rows.Next() {
		err = rows.Scan(&user.Id,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Mobile,
			&user.Role,
			&user.RoleName)
		user.Order = i
		i++
		sec.CheckInternalServerError(err, w)
		users = append(users, user)
	}
	query = "SELECT id, name FROM roles ORDER BY name asc"
	log.Println("Query Roles: " + query)
	rows, err = Db.Query(query)
	sec.CheckInternalServerError(err, w)
	var roles []mdl.Role
	var role mdl.Role
	i = 1
	for rows.Next() {
		err = rows.Scan(&role.Id,
			&role.Name)
		role.Order = i
		i++
		sec.CheckInternalServerError(err, w)
		roles = append(roles, role)
	}
	var page mdl.PageUsers
	page.Users = users
	page.Roles = roles
	page.AppName = mdl.AppName
	page.Title = "Usuários"
	page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
	var tmpl = template.Must(template.ParseGlob("tiles/users/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Users", page)
	sec.CheckInternalServerError(err, w)
}
