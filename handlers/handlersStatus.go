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

func CreateStatusHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Status")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "INSERT INTO status(name, stereotype) VALUES ($1, $2) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, stereotype).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Stereotype: " + stereotype)
	}
	http.Redirect(w, r, route.StatusRoute, 301)
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Status")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "UPDATE status SET name=$1, stereotype=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, stereotype, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Stereotype: " + stereotype)
	}
	http.Redirect(w, r, route.StatusRoute, 301)
}

func DeleteStatusHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Status")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM status WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.StatusRoute, 301)
}

func ListStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Status")
	sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT id, name, stereotype FROM status order by id asc")
	sec.CheckInternalServerError(err, w)
	var status_array []mdl.Status
	var status mdl.Status
	var i = 1
	for rows.Next() {
		err = rows.Scan(&status.Id, &status.Name, &status.Stereotype)
		sec.CheckInternalServerError(err, w)
		status.Order = i
		i++
		status_array = append(status_array, status)
	}
	var page mdl.PageStatus
	page.Status = status_array
	page.Title = "Status"
	var tmpl = template.Must(template.ParseGlob("tiles/status/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Status", page)
	sec.CheckInternalServerError(err, w)
}
