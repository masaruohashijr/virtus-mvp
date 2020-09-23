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

func CreateEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Titulo")
		sqlStatement := "INSERT INTO entidades(titulo) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Titulo")
		sqlStatement := "UPDATE entidades SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Título: " + titulo)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM entidades WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListEntidadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Entidades")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, titulo FROM entidades order by id asc")
		sec.CheckInternalServerError(err, w)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			err = rows.Scan(&entidade.Id, &entidade.Titulo)
			sec.CheckInternalServerError(err, w)
			entidade.Order = i
			i++
			entidades = append(entidades, entidade)
		}
		var page mdl.PageEntidades
		page.Entidades = entidades
		page.AppName = mdl.AppName
		page.Title = "Entidades"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/entidades/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
