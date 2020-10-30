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

func CreateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		statusComponenteId := GetStartStatus("componente")
		sqlStatement := "INSERT INTO componentes(nome, descricao, author_id, criado_em, status_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		idComponente := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now(), statusComponenteId).Scan(&idComponente)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(idComponente) + " | Nome: " + nome + " | Descrição: " + descricao)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "elementoComponente") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				elementoComponenteId := 0
				statusElementoId := GetStartStatus("elemento")
				elementoId := strings.Split(array[3], ":")[1]
				tipoNotaId := strings.Split(array[3], ":")[1]
				pesoPadrao := strings.Split(array[5], ":")[1]
				sqlStatement := " INSERT INTO " +
					" public.elementos_componentes( " +
					" componente_id, " +
					" elemento_id, " +
					" tipo_nota_id, " +
					" peso_padrao, " +
					" author_id, " +
					" criado_em, " +
					" status_id) " +
					" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(
					sqlStatement,
					idComponente,
					elementoId,
					tipoNotaId,
					pesoPadrao,
					currentUser.Id,
					time.Now(),
					statusElementoId).Scan(&elementoComponenteId)
				if err != nil {
					panic(err.Error())
				}
			}
		}

		http.Redirect(w, r, route.ComponentesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		componenteId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE componentes SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, componenteId)
		log.Println("UPDATE: Id: " + componenteId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Elementos Componentes
		var elementosComponenteDB = ListElementosByComponenteId(componenteId)
		var elementosComponentePage []mdl.ElementoComponente
		var elementoComponentePage mdl.ElementoComponente
		for key, value := range r.Form {
			if strings.HasPrefix(key, "elementoComponente") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				elementoComponentePage.Id, _ = strconv.ParseInt(id, 10, 64)
				elementoComponentePage.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
				elementoId := strings.Split(array[3], ":")[1]
				log.Println("elementoId -------- " + elementoId)
				elementoComponentePage.ElementoId, _ = strconv.ParseInt(elementoId, 10, 64)
				elementoNome := strings.Split(array[4], ":")[1]
				log.Println("elementoNome -------- " + elementoNome)
				elementoComponentePage.ElementoNome = elementoNome
				pesoPadrao := strings.Split(array[5], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				elementoComponentePage.PesoPadrao, _ = strconv.Atoi(pesoPadrao)
				authorId := strings.Split(array[6], ":")[1]
				log.Println("authorId -------- " + authorId)
				elementoComponentePage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[7], ":")[1]
				log.Println("authorName -------- " + authorName)
				elementoComponentePage.AuthorName = authorName
				criadoEm := strings.Split(array[8], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				elementoComponentePage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[9], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				elementoComponentePage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[10], ":")[1]
				log.Println("StatusId -------- " + statusId)
				elementoComponentePage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[11], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				elementoComponentePage.CStatus = cStatus
				elementosComponentePage = append(elementosComponentePage, elementoComponentePage)
			}
		}
		if len(elementosComponentePage) < len(elementosComponenteDB) {
			log.Println("Quantidade de Elementos do Componente da Página: " + strconv.Itoa(len(elementosComponentePage)))
			if len(elementosComponentePage) == 0 {
				DeleteElementosComponenteByComponenteId(componenteId) //DONE
			} else {
				var diffDB []mdl.ElementoComponente = elementosComponenteDB
				for n := range elementosComponentePage {
					if containsElementoComponente(diffDB, elementosComponentePage[n]) {
						diffDB = removeElementoComponente(diffDB, elementosComponentePage[n])
					}
				}
				DeleteElementosComponenteHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.ElementoComponente = elementosComponentePage
			for n := range elementosComponenteDB {
				if containsElementoComponente(diffPage, elementosComponenteDB[n]) {
					diffPage = removeElementoComponente(diffPage, elementosComponenteDB[n])
				}
			}
			var elementoComponente mdl.ElementoComponente
			elementoComponenteId := 0
			statusElementoId := GetStartStatus("elemento")
			for i := range diffPage {
				elementoComponente = diffPage[i]
				log.Println("Componente Id: " + componenteId)
				sqlStatement := "INSERT INTO public.elementos_componentes ( " +
					" componente_id, " +
					" elemento_id, " +
					" peso_padrao, " +
					" author_id, " +
					" criado_em, " +
					" status_id " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(
					sqlStatement,
					componenteId,
					elementoComponente.ElementoId,
					elementoComponente.PesoPadrao,
					currentUser.Id,
					time.Now(),
					statusElementoId).Scan(&elementoComponenteId)
			}
		}
		UpdateElementosComponenteHandler(elementosComponentePage, elementosComponenteDB)

	}
	http.Redirect(w, r, route.ComponentesRoute, 301)
}

func DeleteComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM elementos_componentes WHERE componente_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)

		sqlStatement = "DELETE FROM componentes WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
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

		sql = "SELECT id, nome FROM tipos_notas ORDER BY nome desc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var tiposNota []mdl.TipoNota
		var tipoNota mdl.TipoNota
		i = 1
		for rows.Next() {
			rows.Scan(&tipoNota.Id, &tipoNota.Nome)
			tipoNota.Order = i
			i++
			tiposNota = append(tiposNota, tipoNota)
		}

		var page mdl.PageComponentes
		page.TiposNota = tiposNota
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

func LoadElementosByComponenteId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Elementos By Componente Id")
	r.ParseForm()
	var componenteId = r.FormValue("componenteId")
	log.Println("componenteId: " + componenteId)
	elementosComponente := ListElementosByComponenteId(componenteId)
	jsonElementosComponente, _ := json.Marshal(elementosComponente)
	w.Write([]byte(jsonElementosComponente))
	log.Println("JSON Elementos de Componente")
}
