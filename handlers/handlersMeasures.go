package handlers

import (
	mdl "diaria/models"
	route "diaria/routes"
	sec "diaria/security"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateMeasureHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Measure")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		sqlStatement := "INSERT INTO Measures(name) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
	}
	http.Redirect(w, r, route.MeasuresRoute, 301)
}

func UpdateMeasureHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Measure")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE measures SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name)
	}
	http.Redirect(w, r, route.MeasuresRoute, 301)
}

func DeleteMeasureHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Measure")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Measures WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.MeasuresRoute, 301)
}

func ListMeasuresHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Measures")
	sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT * FROM measures")
	sec.CheckInternalServerError(err, w)
	var measures []mdl.Measure
	var measure mdl.Measure
	var i = 1
	for rows.Next() {
		err = rows.Scan(&measure.Id, &measure.Name)
		sec.CheckInternalServerError(err, w)
		measure.Order = i
		i++
		measures = append(measures, measure)
	}
	var page mdl.PageMeasures
	page.Measures = measures
	page.Title = "Medidas Usuais"
	var tmpl = template.Must(template.ParseGlob("tiles/measures/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Measure", page)
	sec.CheckInternalServerError(err, w)
}
