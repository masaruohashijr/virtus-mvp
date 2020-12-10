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
		referencia := r.FormValue("Referencia")
		statusComponenteId := GetStartStatus("componente")
		sqlStatement := "INSERT INTO componentes(nome, descricao, referencia, author_id, criado_em, status_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		idComponente := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, referencia, currentUser.Id, time.Now(), statusComponenteId).Scan(&idComponente)
		if err != nil {
			log.Println(err.Error())
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
					" elementos_componentes( " +
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
					log.Println(err.Error())
				}
			}
			if strings.HasPrefix(key, "TipoNota_") {
				log.Println(value[0])
				tipoNotaComponenteId := 0
				tipoNotaId := strings.Split(key, "_")[1]
				pesoPadrao := value[0]
				statusTipoNotaId := GetStartStatus("tipo_nota")
				sqlStatement := " INSERT INTO " +
					" tipos_notas_componentes( " +
					" componente_id," +
					" tipo_nota_id," +
					" peso_padrao, " +
					" author_id, " +
					" criado_em, " +
					" status_id) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(
					sqlStatement,
					idComponente,
					tipoNotaId,
					pesoPadrao,
					currentUser.Id,
					time.Now(),
					statusTipoNotaId).Scan(&tipoNotaComponenteId)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
		http.Redirect(w, r, route.ComponentesRoute+"?msg=Componente criado com sucesso.", 301)
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
		referencia := r.FormValue("Referencia")
		sqlStatement := "UPDATE componentes SET nome=$1, descricao=$2, referencia=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		updtForm.Exec(nome, descricao, referencia, componenteId)
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
				tipoNotaId := strings.Split(array[5], ":")[1]
				log.Println("tipoNotaId -------- " + tipoNotaId)
				elementoComponentePage.TipoNotaId, _ = strconv.ParseInt(tipoNotaId, 10, 64)
				tipoNotaNome := strings.Split(array[6], ":")[1]
				log.Println("tipoNotaNome -------- " + tipoNotaNome)
				elementoComponentePage.TipoNotaNome = tipoNotaNome
				pesoPadrao := strings.Split(array[7], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				elementoComponentePage.PesoPadrao, _ = strconv.Atoi(pesoPadrao)
				authorId := strings.Split(array[8], ":")[1]
				log.Println("authorId -------- " + authorId)
				elementoComponentePage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[9], ":")[1]
				log.Println("authorName -------- " + authorName)
				elementoComponentePage.AuthorName = authorName
				criadoEm := strings.Split(array[10], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				elementoComponentePage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[11], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				elementoComponentePage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[12], ":")[1]
				log.Println("StatusId -------- " + statusId)
				elementoComponentePage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[13], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				elementoComponentePage.CStatus = cStatus
				elementosComponentePage = append(elementosComponentePage, elementoComponentePage)
			}
			if strings.HasPrefix(key, "TipoNota_") {
				tipoNotaId := strings.Split(key, "_")[1]
				tipoNotaPeso := value[0]
				sqlStatement = "INSERT INTO tipos_notas_componentes (tipo_nota_id,componente_id) " +
					" SELECT " + tipoNotaId + ", " + componenteId +
					" WHERE NOT EXISTS (select 1 from tipos_notas_componentes " +
					" WHERE tipo_nota_id = " + tipoNotaId + " AND componente_id = " + componenteId + ")"
				log.Println(sqlStatement)
				Db.QueryRow(sqlStatement)
				sqlStatement = "UPDATE tipos_notas_componentes SET peso_padrao=$1 WHERE tipo_nota_id=$2 AND componente_id = $3"
				updtForm, err = Db.Prepare(sqlStatement)
				if err != nil {
					log.Println(err.Error())
				}
				updtForm.Exec(value[0], tipoNotaId, componenteId)
				log.Println("UPDATE: Tipo Nota PESO: " + tipoNotaPeso + " - Id: " + tipoNotaId)
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
	http.Redirect(w, r, route.ComponentesRoute+"?msg=Componente atualizado com sucesso.", 301)
}

func DeleteComponenteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Componente")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM tipos_notas_componentes WHERE componente_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)

		sqlStatement = "DELETE FROM elementos_componentes WHERE componente_id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)

		sqlStatement = "DELETE FROM componentes WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ComponentesRoute+"?msg=Componente removido com sucesso.", 301)
}

func ListComponentesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Componentes")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listCiclos") {
		errMsg := r.FormValue("errMsg")
		msg := r.FormValue("msg")
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" coalesce(a.descricao,''), " +
			" coalesce(a.referencia,''), " +
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
		defer rows.Close()
		var componentes []mdl.Componente
		var componente mdl.Componente
		var i = 1
		for rows.Next() {
			rows.Scan(
				&componente.Id,
				&componente.Nome,
				&componente.Descricao,
				&componente.Referencia,
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
		defer rows.Close()
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
		defer rows.Close()
		var tiposNota []mdl.TipoNota
		var tipoNota mdl.TipoNota
		i = 1
		for rows.Next() {
			rows.Scan(&tipoNota.Id, &tipoNota.Nome)
			tiposNota = append(tiposNota, tipoNota)
		}

		var page mdl.PageComponentes
		if msg != "" {
			page.Msg = msg
		}
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
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

func LoadTiposNotaByComponenteId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Tipos Nota By Componente Id")
	r.ParseForm()
	var componenteId = r.FormValue("componenteId")
	tiposNotasComponentes := ListTiposNotaByComponenteId(componenteId)
	jsonTiposNotasComponentes, _ := json.Marshal(tiposNotasComponentes)
	w.Write([]byte(jsonTiposNotasComponentes))
	log.Println("JSON Tipos Notas Componentes")
}
