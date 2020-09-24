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
		currentUser := GetUserInCookie(w, r)
		statusElementoId := GetStartStatus("elemento")
		titulo := r.FormValue("TituloElementoForInsert")
		descricao := r.FormValue("DescricaoElementoForInsert")
		sqlStatement := "INSERT INTO elementos(titulo, descricao, author_id, data_criacao, status_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		elementoId := 0
		authorId := strconv.FormatInt(GetUserInCookie(w, r).Id, 10)
		err := Db.QueryRow(sqlStatement, titulo, descricao, authorId, time.Now(), statusElementoId).Scan(&elementoId)
		log.Println(sqlStatement + " :: " + titulo)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(elementoId) + " | Título: " + titulo)
		statusItemId := GetStartStatus("itemAAvaliar")
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				itemId := 0
				tituloItem := strings.Split(array[3], ":")[1]
				descricaoItem := strings.Split(array[4], ":")[1]
				avaliacaoItem := strings.Split(array[5], ":")[1]
				log.Println("itemId: " + strconv.Itoa(itemId))
				sqlStatement := "INSERT INTO public.itens( " +
					" elemento_id, titulo, descricao, avaliacao, data_criacao, author_id, status_id ) " +
					" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
				log.Println(sqlStatement)
				log.Println("elementoId: " + strconv.Itoa(elementoId))
				err = Db.QueryRow(sqlStatement, elementoId, tituloItem, descricaoItem, avaliacaoItem, time.Now(), currentUser.Id, statusItemId).Scan(&itemId)
				log.Println("itemId: " + strconv.Itoa(itemId))
				if err != nil {
					panic(err.Error())
				}
			}
			http.Redirect(w, r, route.ElementosRoute, 301)
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
