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

func CreateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO componentes(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome + " | Descrição: " + descricao)
		http.Redirect(w, r, route.ComponentesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE componentes SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome + " | Descrição: " + descricao)
	}
	http.Redirect(w, r, route.ComponentesRoute, 301)
}

func DeleteComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM componentes WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ComponentesRoute, 301)
}

func ListComponentesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Componentes")
	if sec.IsAuthenticated(w, r) {
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM componentes a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
		var componentes []mdl.Componente
		var componente mdl.Componente
		var i = 1
		for rows.Next() {
			rows.Scan(
				&componente.Id,
				&componente.Nome,
				&componente.Descricao,
				&componente.AuthorId,
				&componente.AuthorName,
				&componente.C_CriadoEm,
				&componente.CStatus,
				&componente.StatusId,
				&componente.IdVersaoOrigem)
			componente.Order = i
			i++
			componentes = append(componentes, componente)
		}
		sql = "SELECT id, nome FROM elementos ORDER BY id asc"
		rows, _ = Db.Query(sql)
		var elementos []mdl.Elemento
		var elemento mdl.Elemento
		i = 1
		for rows.Next() {
			rows.Scan(&elemento.Id, &elemento.Nome)
			elemento.Order = i
			i++
			elementos = append(elementos, elemento)
		}
		var page mdl.PageComponentes
		page.Componentes = componentes
		page.Elementos = elementos
		page.AppName = mdl.AppName
		page.Title = "Componentes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/componentes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Componentes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
