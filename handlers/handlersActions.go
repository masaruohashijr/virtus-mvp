package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Action")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		sqlStatement := "INSERT INTO actions(name) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
	}
	http.Redirect(w, r, route.ActionsRoute, 301)
}

func UpdateActionHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Action")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE status SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name)
	}
	http.Redirect(w, r, route.ActionsRoute, 301)
}

func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Action")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM actions WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ActionsRoute, 301)
}

func ListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Actions")
	sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT id, name FROM actions order by id asc")
	sec.CheckInternalServerError(err, w)
	var actions []mdl.Action
	var action mdl.Action
	var i = 1
	for rows.Next() {
		err = rows.Scan(&action.Id, &action.Name)
		sec.CheckInternalServerError(err, w)
		action.Order = i
		i++
		actions = append(actions, action)
	}
	var page mdl.PageAction
	page.Actions = actions
	page.Title = "Action"
	var tmpl = template.Must(template.ParseGlob("tiles/status/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Action", page)
	sec.CheckInternalServerError(err, w)
}
