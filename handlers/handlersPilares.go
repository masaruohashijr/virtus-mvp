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

func CreatePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO pilares(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		idPilar := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&idPilar)
		if err != nil {
			panic(err.Error())
		}

		log.Println("INSERT: Id: " + strconv.Itoa(idPilar) + " | Nome: " + nome + " | Descrição: " + descricao)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "componentePilar") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				componentePilarId := 0
				componenteId := strings.Split(array[3], ":")[1]
				tipoMediaId := strings.Split(array[5], ":")[1]
				sonda := strings.Split(array[7], ":")[1]
				pesoPadrao := strings.Split(array[8], ":")[1]
				sqlStatement := " INSERT INTO " +
					" public.componentes_pilares( " +
					" pilar_id, " +
					" componente_id, " +
					" tipo_media, " +
					" peso_padrao, " +
					" sonda, " +
					" author_id, " +
					" criado_em ) " +
					" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
				log.Println(sqlStatement)
				err = Db.QueryRow(
					sqlStatement,
					idPilar,
					componenteId,
					tipoMediaId,
					pesoPadrao,
					sonda,
					currentUser.Id,
					time.Now()).Scan(&componentePilarId)
				if err != nil {
					panic(err.Error())
				}
			}
		}
		http.Redirect(w, r, route.PilaresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdatePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		pilarId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE pilares SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, pilarId)
		log.Println("UPDATE: Id: " + pilarId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Componentes Pilares
		var componentesPilarDB = ListComponentesByPilarId(pilarId)
		var componentesPilarPage []mdl.ComponentePilar
		var componentePilarPage mdl.ComponentePilar
		for key, value := range r.Form {
			if strings.HasPrefix(key, "componentePilar") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				componentePilarPage.Id, _ = strconv.ParseInt(id, 10, 64)
				componentePilarPage.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
				componenteId := strings.Split(array[3], ":")[1]
				log.Println("componenteId -------- " + componenteId)
				componentePilarPage.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
				componenteNome := strings.Split(array[4], ":")[1]
				log.Println("componenteNome -------- " + componenteNome)
				componentePilarPage.ComponenteNome = componenteNome
				tipoMediaId := strings.Split(array[5], ":")[1]
				log.Println("tipoMediaId -------- " + tipoMediaId)
				componentePilarPage.TipoMediaId, _ = strconv.Atoi(tipoMediaId)
				tipoMedia := strings.Split(array[6], ":")[1]
				log.Println("tipoMedia -------- " + tipoMedia)
				componentePilarPage.TipoMedia = tipoMedia
				sonda := strings.Split(array[7], ":")[1]
				log.Println("sonda -------- " + sonda)
				componentePilarPage.Sonda = sonda
				pesoPadrao := strings.Split(array[8], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				componentePilarPage.PesoPadrao = pesoPadrao
				authorId := strings.Split(array[9], ":")[1]
				log.Println("authorId -------- " + authorId)
				componentePilarPage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[10], ":")[1]
				log.Println("authorName -------- " + authorName)
				componentePilarPage.AuthorName = authorName
				criadoEm := strings.Split(array[11], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				componentePilarPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[12], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				componentePilarPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[13], ":")[1]
				log.Println("StatusId -------- " + statusId)
				componentePilarPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[14], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				componentePilarPage.CStatus = cStatus
				componentesPilarPage = append(componentesPilarPage, componentePilarPage)
			}
		}
		if len(componentesPilarPage) < len(componentesPilarDB) {
			log.Println("Quantidade de Componentes do Pilar da Página: " + strconv.Itoa(len(componentesPilarPage)))
			if len(componentesPilarPage) == 0 {
				DeleteComponentesPilarByPilarId(pilarId) //DONE
			} else {
				var diffDB []mdl.ComponentePilar = componentesPilarDB
				for n := range componentesPilarPage {
					if containsComponentePilar(diffDB, componentesPilarPage[n]) {
						diffDB = removeComponentePilar(diffDB, componentesPilarPage[n])
					}
				}
				DeleteComponentesPilarHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.ComponentePilar = componentesPilarPage
			for n := range componentesPilarDB {
				if containsComponentePilar(diffPage, componentesPilarDB[n]) {
					diffPage = removeComponentePilar(diffPage, componentesPilarDB[n])
				}
			}
			var componentePilar mdl.ComponentePilar
			componentePilarId := 0
			statusComponenteId := GetStartStatus("plano")
			for i := range diffPage {
				componentePilar = diffPage[i]
				log.Println("Pilar Id: " + pilarId)
				sqlStatement := "INSERT INTO public.componentes_pilares ( " +
					" pilar_id, " +
					" componente_id, " +
					" peso_padrao, " +
					" author_id, " +
					" criado_em, " +
					" status_id " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(
					sqlStatement,
					pilarId,
					componentePilar.ComponenteId,
					componentePilar.PesoPadrao,
					currentUser.Id,
					time.Now(),
					statusComponenteId).Scan(&componentePilarId)
			}
		}
		UpdateComponentesPilarHandler(componentesPilarPage, componentesPilarDB)

		http.Redirect(w, r, route.PilaresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeletePilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Pilar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM componentes_pilares WHERE pilar_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)

		sqlStatement = "DELETE FROM pilares WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
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
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listPilares") {
		errMsg := r.FormValue("errMsg")
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
			" FROM pilares a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
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
		sql = "SELECT id, nome FROM componentes ORDER BY id asc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var componentes []mdl.Componente
		var componente mdl.Componente
		i = 1
		for rows.Next() {
			rows.Scan(&componente.Id, &componente.Nome)
			componente.Order = i
			i++
			componentes = append(componentes, componente)
		}
		var page mdl.PagePilares
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		page.Pilares = pilares
		page.Componentes = componentes
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

func LoadComponentesByPilarId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Componentes By Pilar Id")
	r.ParseForm()
	var pilarId = r.FormValue("pilarId")
	log.Println("pilarId: " + pilarId)
	componentesPilar := ListComponentesByPilarId(pilarId)
	jsonComponenesPilar, _ := json.Marshal(componentesPilar)
	w.Write([]byte(jsonComponenesPilar))
	log.Println("JSON Componentes de Pilar")
}
