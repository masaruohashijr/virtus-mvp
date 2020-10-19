package handlers

import (
	//	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreatePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO pilares(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome + " | Descrição: " + descricao)
		http.Redirect(w, r, route.PilaresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdatePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE pilares SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome + " | Descrição: " + descricao)
		http.Redirect(w, r, route.PilaresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeletePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM pilares WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.PilaresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListPilaresHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Pilares")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM pilares a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		log.Println(query)
		rows, _ := Db.Query(query)
		var pilares []mdl.Pilar
		var pilar mdl.Pilar
		var i = 1
		for rows.Next() {
			rows.Scan(
				&pilar.Id,
				&pilar.Nome,
				&pilar.Descricao,
				&pilar.AuthorId,
				&pilar.AuthorName,
				&pilar.C_CriadoEm,
				&pilar.CStatus,
				&pilar.StatusId,
				&pilar.IdVersaoOrigem)
			pilar.Order = i
			i++
			pilares = append(pilares, pilar)
		}
		var page mdl.PagePilares
		page.Pilares = pilares
		page.AppName = mdl.AppName
		page.Title = "Pilares"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/pilares/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Pilares", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
