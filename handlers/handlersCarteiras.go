package handlers

import (
	//	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateCarteiraHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Carteira")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		nome := r.FormValue("Nome")
		sqlStatement := "INSERT INTO carteiras(nome) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome).Scan(&id)
		log.Println(sqlStatement + " :: " + nome)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome)
		http.Redirect(w, r, route.CarteirasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateCarteiraHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Carteira")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		sqlStatement := "UPDATE carteiras SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome)
		http.Redirect(w, r, route.CarteirasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteCarteiraHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Carteira")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM carteiras WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.CarteirasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListCarteirasHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Carteiras")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, nome FROM carteiras order by id asc")
		sec.CheckInternalServerError(err, w)
		var carteiras []mdl.Carteira
		var carteira mdl.Carteira
		var i = 1
		for rows.Next() {
			err = rows.Scan(&carteira.Id, &carteira.Nome)
			sec.CheckInternalServerError(err, w)
			carteira.Order = i
			i++
			carteiras = append(carteiras, carteira)
		}
		var page mdl.PageCarteiras
		page.Carteiras = carteiras
		page.AppName = mdl.AppName
		page.Title = "Carteiras"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/carteiras/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Carteiras", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
