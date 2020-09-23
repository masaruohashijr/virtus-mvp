package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateElementoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Elemento")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("Titulo")
		sqlStatement := "INSERT INTO elementos(titulo) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, titulo).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		http.Redirect(w, r, route.ElementosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateElementoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Elemento")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		titulo := r.FormValue("Titulo")
		sqlStatement := "UPDATE elementos SET titulo=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(titulo, id)
		log.Println("UPDATE: Id: " + id + " | Título: " + titulo)
	}
	http.Redirect(w, r, route.ElementosRoute, 301)
}

func DeleteElementoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Elemento")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM elementos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ElementosRoute, 301)
}

func DeleteElementosByRoleHandler(roleId string) {
	sqlStatement := "DELETE FROM elementos_roles WHERE role_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(roleId)
	log.Println("DELETE elementos_roles in Role Id: " + roleId)
}
func DeleteElementosHandler(diffDB []mdl.Elemento) {
	sqlStatement := "DELETE FROM elementos_roles WHERE elemento_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Elemento Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func ListElementosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Elementos")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, titulo FROM elementos order by id asc")
		sec.CheckInternalServerError(err, w)
		var elementos []mdl.Elemento
		var elemento mdl.Elemento
		var i = 1
		for rows.Next() {
			err = rows.Scan(&elemento.Id, &elemento.Titulo)
			sec.CheckInternalServerError(err, w)
			elemento.Order = i
			i++
			elementos = append(elementos, elemento)
		}
		var page mdl.PageElementos
		page.Elementos = elementos
		page.AppName = mdl.AppName
		page.Title = "Elementos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/elementos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Elementos", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

// AJAX
func ListElementosByRoleIdHandler(roleId string) []mdl.Elemento {
	log.Println("List Elementos By Role Id")
	sql := "SELECT elemento_id" +
		" FROM elementos_roles WHERE role_id= $1"
	log.Println(sql)
	rows, _ := Db.Query(sql, roleId)
	var elementos []mdl.Elemento
	var elemento mdl.Elemento
	for rows.Next() {
		rows.Scan(&elemento.Id)
		elementos = append(elementos, elemento)
	}
	return elementos
}

func LoadAvailableElementos(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Load Available Elementos")
	r.ParseForm()
	savedUser := GetUserInCookie(w, r)
	var statusId = r.FormValue("statusId")
	var entityType = r.FormValue("entityType")
	log.Println("entityType: " + entityType)
	log.Println("statusId: " + statusId)
	sql := " SELECT a.id, a.name, a.code " +
		" FROM elementos a INNER JOIN elementos_activities b ON a.id = b.elemento_id " +
		" INNER JOIN activities c ON c.id = b.activity_id " +
		" INNER JOIN actions d ON c.action_id = d.id " +
		" INNER JOIN workflows e ON c.workflow_id = e.id " +
		" WHERE e.end_at IS null " +
		" AND e.entity_type = $1 " +
		" AND d.origin_status_id = $2 " +
		" AND a.id in ( SELECT elemento_id from elementos_roles where role_id = $3 ) "
	log.Println("Query Available Elementos: " + sql)
	rows, _ := Db.Query(sql, entityType, statusId, savedUser.Role)
	var elementos []mdl.Elemento
	var elemento mdl.Elemento
	for rows.Next() {
		rows.Scan(&elemento.Id, &elemento.Titulo)
		elementos = append(elementos, elemento)
		log.Println(elementos)
	}
	jsonElementos, _ := json.Marshal(elementos)
	w.Write([]byte(jsonElementos))
	log.Println("JSON Load Elementos")
}
