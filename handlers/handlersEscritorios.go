package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateEscritorioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Escritorio")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		abreviatura := r.FormValue("Abreviatura")
		chefe := r.FormValue("Chefe")
		sqlStatement := "INSERT INTO escritorios(nome, descricao, abreviatura, chefe_id, author_id, criado_em) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, abreviatura, chefe, currentUser.Id, time.Now()).Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome + " | Descrição: " + descricao)
		if chefe != "" {
			sqlStatement = "INSERT INTO membros ( " +
				" escritorio_id, " +
				" usuario_id, " +
				" author_id, " +
				" criado_em " +
				" ) " +
				" VALUES ($1, $2, $3, $4) "
			log.Println(sqlStatement)
			Db.QueryRow(
				sqlStatement,
				id,
				chefe,
				currentUser.Id,
				time.Now())
		}
		http.Redirect(w, r, route.EscritoriosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateEscritorioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Escritorio")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		abreviatura := r.FormValue("Abreviatura")
		chefe := r.FormValue("Chefe")
		if chefe != "" {
			sqlStatement := "UPDATE escritorios SET nome=$1, descricao=$2, abreviatura=$3, chefe_id=$4 WHERE id=$5"
			updtForm, _ := Db.Prepare(sqlStatement)
			updtForm.Exec(nome, descricao, abreviatura, chefe, id)
			sqlStatement = "INSERT INTO membros ( " +
				" escritorio_id, " +
				" usuario_id, " +
				" author_id, " +
				" criado_em " +
				" ) " +
				" SELECT $1, $2, $3, $4 WHERE NOT EXISTS " +
				" (SELECT 1 FROM membros WHERE escritorio_id = $5 AND usuario_id =$6)"
			log.Println(sqlStatement)
			Db.QueryRow(
				sqlStatement,
				id,
				chefe,
				GetUserInCookie(w, r).Id,
				time.Now(),
				id,
				chefe)
		} else {
			sqlStatement := "UPDATE escritorios SET nome=$1, descricao=$2, abreviatura=$3 WHERE id=$4"
			updtForm, _ := Db.Prepare(sqlStatement)
			updtForm.Exec(nome, descricao, abreviatura, id)
		}
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome + " | Abreviatura: " + abreviatura + " | Descrição: " + descricao + " | Chefe: " + chefe)
		http.Redirect(w, r, route.EscritoriosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteEscritorioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Escritorio")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM escritorios WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.EscritoriosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListEscritoriosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Escritorios")
	if sec.IsAuthenticated(w, r) {
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.abreviatura, " +
			" coalesce(a.chefe_id,0), " +
			" coalesce(d.name,'') as chefe_name, " +
			" a.author_id, " +
			" coalesce(b.name,'') as author_name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM escritorios a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" LEFT JOIN users d ON a.chefe_id = d.id " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
		var escritorios []mdl.Escritorio
		var escritorio mdl.Escritorio
		var i = 1
		for rows.Next() {
			rows.Scan(
				&escritorio.Id,
				&escritorio.Nome,
				&escritorio.Descricao,
				&escritorio.Abreviatura,
				&escritorio.ChefeId,
				&escritorio.ChefeNome,
				&escritorio.AuthorId,
				&escritorio.AuthorName,
				&escritorio.C_CriadoEm,
				&escritorio.CStatus,
				&escritorio.StatusId,
				&escritorio.IdVersaoOrigem)
			escritorio.Order = i
			i++
			escritorios = append(escritorios, escritorio)
		}
		var page mdl.PageEscritorios
		page.Escritorios = escritorios

		sql = "SELECT id, name FROM users ORDER BY name asc"
		rows, _ = Db.Query(sql)
		var users []mdl.User
		var user mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&user.Id, &user.Name)
			user.Order = i
			i++
			users = append(users, user)
		}
		page.Users = users

		sql = "SELECT id, nome FROM entidades ORDER BY nome asc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		i = 1
		for rows.Next() {
			rows.Scan(&entidade.Id, &entidade.Nome)
			entidade.Order = i
			i++
			entidades = append(entidades, entidade)
		}
		page.Entidades = entidades

		page.AppName = mdl.AppName
		page.Title = "Escritórios"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/escritorios/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Escritorios", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadJurisdicoesByEscritorioId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Jurisdições By Escritório Id")
	r.ParseForm()
	var escritorioId = r.FormValue("escritorioId")
	log.Println("escritorioId: " + escritorioId)
	jurisdicoes := ListJurisdicoesByEscritorioId(escritorioId)
	jsonJurisdicoes, _ := json.Marshal(jurisdicoes)
	w.Write([]byte(jsonJurisdicoes))
	log.Println("JSON Jurisdições")
}

func LoadMembrosByEscritorioId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Membros By Escritório Id")
	r.ParseForm()
	var escritorioId = r.FormValue("escritorioId")
	log.Println("escritorioId: " + escritorioId)
	membros := ListMembrosByEscritorioId(escritorioId)
	jsonMembros, _ := json.Marshal(membros)
	w.Write([]byte(jsonMembros))
	log.Println("JSON Membros")
}
