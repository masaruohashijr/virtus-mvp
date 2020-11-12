package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	e "virtus/errors"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Status")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		name := r.FormValue("Name")
		description := r.FormValue("Descricao")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "INSERT INTO status(name, description, stereotype) VALUES ($1, $2, $3) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, description, stereotype).Scan(&id)
		if err != nil && strings.Contains(err.Error(), "duplicate key") {
			page := listStatus(e.ErroChaveDuplicada)
			page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
			var tmpl = template.Must(template.ParseGlob("tiles/status/*"))
			tmpl.ParseGlob("tiles/*")
			tmpl.ExecuteTemplate(w, "Main-Status", page)
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Descricao: " + description + " | Stereotype: " + stereotype)
		http.Redirect(w, r, route.StatusRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Status")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		description := r.FormValue("Description")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "UPDATE status SET name=$1, description=$2, stereotype=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, description, stereotype, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Descricao: " + description + " | Stereotype: " + stereotype)
		http.Redirect(w, r, route.StatusRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Status")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM status WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		_, err = deleteForm.Exec(id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.StatusRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Status")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listStatus") {
		page := listStatus("")
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/status/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Status", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func listStatus(errorMsg string) mdl.PageStatus {
	sql := "SELECT " +
		" a.id, " +
		" a.name, " +
		" coalesce(a.description,'') as desc, " +
		" coalesce(a.stereotype,'') as stereo_type, " +
		" a.author_id, " +
		" b.name, " +
		" to_char(a.created_at,'DD/MM/YYYY HH24:MI:SS'), " +
		" coalesce(c.name,'') as cstatus, " +
		" a.status_id, " +
		" a.id_versao_origem " +
		" FROM status a LEFT JOIN users b " +
		" ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" order by id asc"
	log.Println("sql: " + sql)
	rows, _ := Db.Query(sql)
	var statuss []mdl.Status
	var status mdl.Status
	var i = 1
	for rows.Next() {
		rows.Scan(
			&status.Id,
			&status.Name,
			&status.Description,
			&status.Stereotype,
			&status.AuthorId,
			&status.AuthorName,
			&status.C_CreatedAt,
			&status.CStatus,
			&status.StatusId,
			&status.IdVersaoOrigem)
		status.Order = i
		i++
		statuss = append(statuss, status)
	}
	var page mdl.PageStatus
	page.Statuss = statuss
	page.AppName = mdl.AppName
	page.Title = "Status"
	if errorMsg != "" {
		page.ErrMsg = errorMsg
	}
	return page
}
