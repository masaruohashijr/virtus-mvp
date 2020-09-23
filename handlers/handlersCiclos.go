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

func CreateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Titulo")
		sqlStatement := "INSERT INTO ciclos(titulo) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Titulo")
		sqlStatement := "UPDATE ciclos SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Título: " + titulo)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM ciclos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListCiclosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Ciclos")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, titulo FROM ciclos order by id asc")
		sec.CheckInternalServerError(err, w)
		var ciclos []mdl.Ciclo
		var componente mdl.Ciclo
		var i = 1
		for rows.Next() {
			err = rows.Scan(&componente.Id, &componente.Titulo)
			sec.CheckInternalServerError(err, w)
			componente.Order = i
			i++
			ciclos = append(ciclos, componente)
		}
		var page mdl.PageCiclos
		page.Ciclos = ciclos
		page.AppName = mdl.AppName
		page.Title = "Ciclos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/ciclos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Ciclos", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
