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
		nome := r.FormValue("NomeElementoForInsert")
		descricao := r.FormValue("DescricaoElementoForInsert")
		sqlStatement := "INSERT INTO elementos(nome, descricao, author_id, criado_em, status_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		elementoId := 0
		authorId := strconv.FormatInt(GetUserInCookie(w, r).Id, 10)
		err := Db.QueryRow(sqlStatement, nome, descricao, authorId, time.Now(), statusElementoId).Scan(&elementoId)
		log.Println(sqlStatement + " :: " + nome)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(elementoId) + " | Nome: " + nome)
		statusItemId := GetStartStatus("itemAAvaliar")
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				itemId := 0
				nomeItem := strings.Split(array[3], ":")[1]
				descricaoItem := strings.Split(array[4], ":")[1]
				//				log.Println("itemId: " + strconv.Itoa(itemId))
				sqlStatement := "INSERT INTO public.itens( " +
					" elemento_id, nome, descricao, criado_em, author_id, status_id ) " +
					" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
				log.Println(sqlStatement)
				//				log.Println("elementoId: " + strconv.Itoa(elementoId))
				err = Db.QueryRow(sqlStatement, elementoId, nomeItem, descricaoItem, time.Now(), currentUser.Id, statusItemId).Scan(&itemId)
				//				log.Println("itemId: " + strconv.Itoa(itemId))
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
		currentUser := GetUserInCookie(w, r)
		elementoId := r.FormValue("ElementoIdForUpdate")
		nome := r.FormValue("ElementoNomeForUpdate")
		descricao := r.FormValue("ElementoDescricaoForUpdate")
		sqlStatement := "UPDATE elementos SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, elementoId)
		log.Println("UPDATE: Id: " + elementoId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Itens
		var itensDB = ListItensHandler(elementoId)
		var itensPage []mdl.Item
		var itemPage mdl.Item
		qtdItensPage := 0
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				log.Println("key: " + key)
				qtdItensPage++
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				itemPage.Id, _ = strconv.ParseInt(id, 10, 64)
				elementoId := strings.Split(array[2], ":")[1]
				log.Println("elementoId -------- " + elementoId)
				itemPage.ElementoId, _ = strconv.ParseInt(id, 10, 64)
				nome := strings.Split(array[3], ":")[1]
				log.Println("nome -------- " + nome)
				itemPage.Nome = nome
				descricao := strings.Split(array[4], ":")[1]
				log.Println("descricao -------- " + descricao)
				itemPage.Descricao = descricao
				itensPage = append(itensPage, itemPage)
			}
		}
		log.Println("-----------------------")
		log.Println("Qtd Itens Page: " + strconv.Itoa(qtdItensPage))
		log.Println("Quantidade de itens da página é: " + strconv.Itoa(len(itensPage)))
		log.Println("Quantidade de itens do banco de dados é: " + strconv.Itoa(len(itensDB)))
		log.Println("-----------------------")
		if len(itensPage) < len(itensDB) {
			log.Println("Quantidade de Itens da Página: " + strconv.Itoa(len(itensPage)))
			if len(itensPage) == 0 {
				DeleteItensByElementoHandler(elementoId) //DONE
			} else {
				var diffDB []mdl.Item = itensDB
				for n := range itensPage {
					if containsItem(diffDB, itensPage[n]) {
						diffDB = removeItem(diffDB, itensPage[n])
					}
				}
				DeleteItensHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Item = itensPage
			for n := range itensDB {
				if containsItem(diffPage, itensDB[n]) {
					diffPage = removeItem(diffPage, itensDB[n])
				}
			}
			log.Println("Quantidade de itens a mais: " + strconv.Itoa(len(diffPage)))
			var item mdl.Item
			itemId := 0
			statusItemId := GetStartStatus("item")
			for i := range diffPage {
				item = diffPage[i]
				log.Println("Elemento Id: " + strconv.FormatInt(item.ElementoId, 10))
				sqlStatement := "INSERT INTO " +
					"itens(elemento_id, nome, descricao, criado_em, author_id, status_id) " +
					"VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5, 'YYYY-MM-DD HH24:MI:SS'),$6,$7) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(sqlStatement, elementoId, item.Nome, item.Descricao, time.Now(), currentUser.Id, statusItemId).Scan(&itemId)
			}
		}
		UpdateItensHandler(itensPage, itensDB) // TODO
		http.Redirect(w, r, route.ElementosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func containsItem(itens []mdl.Item, itemCompared mdl.Item) bool {
	for n := range itens {
		if itens[n].Id == itemCompared.Id {
			return true
		}
	}
	return false
}

func removeItem(itens []mdl.Item, itemToBeRemoved mdl.Item) []mdl.Item {
	var newItens []mdl.Item
	for i := range itens {
		if itens[i].Id != itemToBeRemoved.Id {
			newItens = append(newItens, itens[i])
		}
	}
	return newItens
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
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listElementos") {
		query := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" coalesce(a.descricao,''), " +
			" coalesce(b.name,'') as author_name, " +
			" coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'),'') as criado_em, " +
			" a.peso, " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id " +
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
				&elemento.Nome,
				&elemento.Descricao,
				&elemento.AuthorName,
				&elemento.C_CriadoEm,
				&elemento.Peso,
				&elemento.CStatus,
				&elemento.StatusId)
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
