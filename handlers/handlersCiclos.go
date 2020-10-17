package handlers

import (
	//	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO ciclos(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		id := 0
		Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&id)
		log.Println(sqlStatement + " - " + nome)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " - Nome: " + nome)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE ciclos SET nome = $1, " +
			" descricao = $2 " +
			" WHERE id = $3 "
		updtForm, _ := Db.Prepare(sqlStatement)
		updtForm.Exec(nome, descricao, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome + " | Descrição: " + descricao)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM ciclos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListCiclosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Ciclos")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" a.id_versao_origem " +
			" FROM ciclos a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" order by id asc"
		log.Println(query)
		rows, _ := Db.Query(query)
		var ciclos []mdl.Ciclo
		var ciclo mdl.Ciclo
		var i = 1
		for rows.Next() {
			rows.Scan(&ciclo.Id, &ciclo.Nome, &ciclo.Descricao, &ciclo.AuthorId, &ciclo.AuthorName, &ciclo.C_CriadoEm, &ciclo.IdVersaoOrigem)
			ciclo.Order = i
			i++
			ciclos = append(ciclos, ciclo)
		}
		var page mdl.PageCiclos
		page.Ciclos = ciclos
		page.AppName = mdl.AppName
		page.Title = "Ciclos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/ciclos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Ciclos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func GetNow() time.Time {
	br, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), 0, br)
	strNow := now.String()
	log.Println("------------------------------------------------------------------------")
	log.Println("- Agora são " + strNow + " em America/Sao_Paulo. ")
	log.Println("------------------------------------------------------------------------")
	txtNow := strings.Split(strings.Split(strings.Split(strNow, " ")[1], ".")[0], ":")
	hora, _ := strconv.Atoi(txtNow[0])
	minuto, _ := strconv.Atoi(txtNow[1])
	segundo, _ := strconv.Atoi(txtNow[2])
	t := time.Date(0000, time.January, 1,
		hora,
		minuto,
		segundo, 0, time.UTC)
	return t
}
