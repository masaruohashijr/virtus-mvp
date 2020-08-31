package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Client")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		sqlStatement := "INSERT INTO Clients(name, username, password, email, mobile) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		id := 0
		err = Db.QueryRow(sqlStatement, name, username, hash, email, mobile).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Username: " + username + " | Password: " + password + " | Email: " + email + " | Mobile: " + mobile)
	}
	http.Redirect(w, r, route.ClientsRoute, 301)
}

func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Client")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		username := r.FormValue("Username")
		email := r.FormValue("Email")
		mobile := r.FormValue("Mobile")
		sqlStatement := "UPDATE Clients SET name=$1, username=$2, email=$3, mobile=$4 WHERE id=$5"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, username, email, mobile, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Username: " + username + " | E-mail: " + email + " | Mobile: " + mobile)
	}
	http.Redirect(w, r, route.ClientsRoute, 301)
}

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Client")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Clients WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ClientsRoute, 301)
}

func ListClientsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Clients")
	sec.IsAuthenticated(w, r)
	query := "SELECT " +
		"id, name, username, password, " +
		"email, mobile " +
		"FROM clients"
	log.Println("Query: " + query)
	rows, err := Db.Query(query)
	sec.CheckInternalServerError(err, w)
	var clients []mdl.Client
	var client mdl.Client
	var i = 1
	for rows.Next() {
		err = rows.Scan(&client.Id, &client.Name, &client.Username, &client.Password, &client.Email, &client.Mobile)
		client.Order = i
		i++
		sec.CheckInternalServerError(err, w)
		clients = append(clients, client)
	}
	var page mdl.PageClients
	page.Clients = clients
	page.Title = "Clientes"
	var tmpl = template.Must(template.ParseGlob("tiles/clients/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Clients", page)
	sec.CheckInternalServerError(err, w)
}
