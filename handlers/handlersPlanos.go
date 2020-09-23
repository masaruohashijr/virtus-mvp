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

func CreatePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Titulo")
		sqlStatement := "INSERT INTO planos(titulo) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdatePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Titulo")
		sqlStatement := "UPDATE planos SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Título: " + titulo)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeletePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM planos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListPlanosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Planos")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, titulo FROM planos order by id asc")
		sec.CheckInternalServerError(err, w)
		var planos []mdl.Plano
		var componente mdl.Plano
		var i = 1
		for rows.Next() {
			err = rows.Scan(&componente.Id, &componente.Titulo)
			sec.CheckInternalServerError(err, w)
			componente.Order = i
			i++
			planos = append(planos, componente)
		}
		var page mdl.PagePlanos
		page.Planos = planos
		page.AppName = mdl.AppName
		page.Title = "Planos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/planos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Planos", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
