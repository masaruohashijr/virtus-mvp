package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateTipoNotaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Tipo Nota")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		r.ParseForm()
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		referencia := r.FormValue("Referencia")
		letra := r.FormValue("Letra")
		corLetra := r.FormValue("CorLetra")
		sqlStatement := "INSERT INTO tipos_notas(nome, descricao, referencia, letra, cor_letra, author_id, criado_em) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
		tipoNotaId := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, referencia, letra, corLetra, currentUser.Id, time.Now()).Scan(&tipoNotaId)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(tipoNotaId) + " | Nome: " + nome)
		http.Redirect(w, r, route.RolesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateTipoNotaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Tipo Nota")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		tipoNotaId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		referencia := r.FormValue("Referencia")
		letra := r.FormValue("Letra")
		corLetra := r.FormValue("CorLetra")
		sqlStatement := "UPDATE tipos_notas SET nome=$1, descricao=$2, descricao=$3, letra=$4, cor_letra=$5 WHERE id=$6"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = updtForm.Exec(nome, descricao, referencia, letra, corLetra, tipoNotaId)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("UPDATE: Id: " + tipoNotaId + " | Nome: " + nome + " | Descrição: " + descricao + " | Cor Letra: " + corLetra)
		http.Redirect(w, r, route.TiposNotasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteTipoNotaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Perfil")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM tipos_notas WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.TiposNotasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListTiposNotasHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Tipos de Notas")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listTiposNotas") {
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.referencia, " +
			" a.letra, " +
			" a.cor_letra, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM tipos_notas a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by id asc"
		log.Println("sql: " + sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var tiposNotas []mdl.TipoNota
		var tipoNota mdl.TipoNota
		var i = 1
		for rows.Next() {
			rows.Scan(
				&tipoNota.Id,
				&tipoNota.Nome,
				&tipoNota.Descricao,
				&tipoNota.Referencia,
				&tipoNota.Letra,
				&tipoNota.CorLetra,
				&tipoNota.AuthorId,
				&tipoNota.AuthorName,
				&tipoNota.C_CreatedAt,
				&tipoNota.CStatus,
				&tipoNota.StatusId,
				&tipoNota.IdVersaoOrigem)
			tipoNota.Order = i
			i++
			tiposNotas = append(tiposNotas, tipoNota)
		}
		var page mdl.PageTiposNotas
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		page.TiposNotas = tiposNotas
		page.AppName = mdl.AppName
		page.Title = "Tipos de Notas"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/tiposnotas/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Tipos-Notas", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
