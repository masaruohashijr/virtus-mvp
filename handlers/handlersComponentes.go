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

func CreateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Titulo")
		sqlStatement := "INSERT INTO componentes(titulo) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		http.Redirect(w, r, route.ComponentesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Titulo")
		sqlStatement := "UPDATE componentes SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Título: " + titulo)
	}
	http.Redirect(w, r, route.ComponentesRoute, 301)
}

func DeleteComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM componentes WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ComponentesRoute, 301)
}

func ListComponentesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Componentes")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, titulo FROM componentes order by id asc")
		sec.CheckInternalServerError(err, w)
		var componentes []mdl.Componente
		var componente mdl.Componente
		var i = 1
		for rows.Next() {
			err = rows.Scan(&componente.Id, &componente.Titulo)
			sec.CheckInternalServerError(err, w)
			componente.Order = i
			i++
			componentes = append(componentes, componente)
		}
		var page mdl.PageComponentes
		page.Componentes = componentes
		page.AppName = mdl.AppName
		page.Title = "Componentes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/componentes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Componentes", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
