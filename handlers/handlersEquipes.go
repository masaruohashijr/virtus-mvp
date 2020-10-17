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

func CreateEquipeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Equipe")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		nome := r.FormValue("Nome")
		sqlStatement := "INSERT INTO equipes(nome) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome).Scan(&id)
		log.Println(sqlStatement + " :: " + nome)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome)
		http.Redirect(w, r, route.EquipesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateEquipeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Equipe")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		sqlStatement := "UPDATE equipes SET nome=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(nome, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome)
		http.Redirect(w, r, route.EquipesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteEquipeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Equipe")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM equipes WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.EquipesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListEquipesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Equipes")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, nome FROM equipes order by id asc")
		sec.CheckInternalServerError(err, w)
		var equipes []mdl.Equipe
		var equipe mdl.Equipe
		var i = 1
		for rows.Next() {
			err = rows.Scan(&equipe.Id, &equipe.Nome)
			sec.CheckInternalServerError(err, w)
			equipe.Order = i
			i++
			equipes = append(equipes, equipe)
		}
		var page mdl.PageEquipes
		page.Equipes = equipes
		page.AppName = mdl.AppName
		page.Title = "Equipes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/equipes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Equipes", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
