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

func CreateMatrizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Matriz")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Nome")
		sqlStatement := "INSERT INTO matrizes(nome) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + titulo)
		http.Redirect(w, r, route.MatrizesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateMatrizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Matriz")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Nome")
		sqlStatement := "UPDATE matrizes SET nome=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + titulo)
		http.Redirect(w, r, route.MatrizesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteMatrizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Matriz")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM matrizes WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.MatrizesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListMatrizesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Matrizes")
	if sec.IsAuthenticated(w, r) {
		rows, _ := Db.Query("SELECT id, nome FROM matrizes order by id asc")
		var matrizes []mdl.Matriz
		var componente mdl.Matriz
		var i = 1
		for rows.Next() {
			rows.Scan(&componente.Id, &componente.Nome)
			componente.Order = i
			i++
			matrizes = append(matrizes, componente)
		}
		var page mdl.PageMatrizes
		page.Matrizes = matrizes
		page.AppName = mdl.AppName
		page.Title = "Matrizes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Matrizes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
