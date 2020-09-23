package handlers

import (
	"encoding/json"
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

func CreateElementoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Elemento")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		titulo := r.FormValue("TituloElementoForInsert")
		descricao := r.FormValue("DescricaoElementoForInsert")
		sqlStatement := "INSERT INTO elementos(titulo, descricao, author_id, data_criacao) VALUES ($1, $2, $3, $4) RETURNING id"
		id := 0
		authorId := strconv.FormatInt(GetUserInCookie(w, r).Id, 10)
		err := Db.QueryRow(sqlStatement, titulo, descricao, authorId, time.Now()).Scan(&id)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Título: " + titulo)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				activityId := 0
				actionId := strings.Split(array[3], ":")[1]
				startAt, _ := time.Parse("yyyy-mm-dd", strings.Split(array[8], ":")[1])
				endAt, _ := time.Parse("yyyy-mm-dd", strings.Split(array[9], ":")[1])
				expTime := strings.Split(array[7], ":")[1]
				if expTime == "" {
					expTime = "0"
				}
				expActionId := strings.Split(array[5], ":")[1]
				log.Println("actionId: " + actionId)
				sqlStatement := "INSERT INTO " +
					"activities(workflow_id, action_id, start_at, end_at, expiration_time_days, expiration_action_id) " +
					"VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
				log.Println(sqlStatement)
				log.Println("wId: " + strconv.Itoa(id) + " | Action: " + actionId + " | ExpDays: " + expTime + " | ExpAction: " + expActionId)
				if expActionId == "" {
					err := Db.QueryRow(sqlStatement, id, actionId, startAt, endAt, expTime, nil).Scan(&activityId)
					sec.CheckInternalServerError(err, w)
				} else {
					err := Db.QueryRow(sqlStatement, id, actionId, startAt, endAt, expTime, expActionId).Scan(&activityId)
					sec.CheckInternalServerError(err, w)
				}
			}
		}
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
		http.Redirect(w, r, route.ElementosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
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
		http.Redirect(w, r, route.ElementosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListElementosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Elementos")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			" a.id, " +
			" a.titulo, " +
			" coalesce(a.descricao,''), " +
			" coalesce(b.name,'') as author_name, " +
			" coalesce(to_char(a.data_criacao,'DD/MM/YYYY'),'') as data_criacao, " +
			" a.peso, " +
			" coalesce(c.name,'') as cstatus " +
			" FROM elementos a " +
			" LEFT JOIN users b ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc "
		rows, err := Db.Query(query)
		log.Println(query)
		sec.CheckInternalServerError(err, w)
		var elementos []mdl.Elemento
		var elemento mdl.Elemento
		var i = 1
		for rows.Next() {
			err = rows.Scan(
				&elemento.Id,
				&elemento.Titulo,
				&elemento.Descricao,
				&elemento.AuthorName,
				&elemento.CDataCriacao,
				&elemento.Peso,
				&elemento.CStatus)
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
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadItensByElementoId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Itens By Elemento Id")
	r.ParseForm()
	var elementoId = r.FormValue("elementoId")
	log.Println("elementoId: " + elementoId)
	itens := ListItensHandler(elementoId)
	jsonItens, _ := json.Marshal(itens)
	w.Write([]byte(jsonItens))
	log.Println("JSON Itens de Elementos")
}
