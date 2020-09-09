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

func CreateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Workflow")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "INSERT INTO workflow(name) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, stereotype).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
	}
	http.Redirect(w, r, route.WorkflowsRoute, 301)
}

func UpdateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Workflow")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE workflow SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name)
	}
	http.Redirect(w, r, route.WorkflowsRoute, 301)
}

func DeleteWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Workflow")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM workflow WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.WorkflowsRoute, 301)
}

func ListWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Workflow")
	sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT id, name FROM workflow order by id asc")
	sec.CheckInternalServerError(err, w)
	var workflows []mdl.Workflow
	var workflow mdl.Workflow
	var i = 1
	for rows.Next() {
		err = rows.Scan(&workflow.Id, &workflow.Name)
		sec.CheckInternalServerError(err, w)
		workflow.Order = i
		i++
		workflows = append(workflows, workflow)
	}
	var page mdl.PageWorkflow
	page.Workflows = workflows
	page.Title = "Workflow"
	var tmpl = template.Must(template.ParseGlob("tiles/workflow/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Workflows", page)
	sec.CheckInternalServerError(err, w)
}
