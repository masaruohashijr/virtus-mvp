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

func CreateEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO entidades(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		idEntidade := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&idEntidade)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(idEntidade) + " | Nome: " + nome + " | Descrição: " + descricao)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "plano") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				planoId := 0
				nomePlano := strings.Split(array[3], ":")[1]
				descricaoPlano := strings.Split(array[4], ":")[1]
				sqlStatement := "INSERT INTO public.planos( " +
					" entidade_id, nome, descricao, author_id, criado_em ) " +
					" VALUES ($1, $2, $3, $4, $5) RETURNING id"
				log.Println(sqlStatement)
				err = Db.QueryRow(sqlStatement, idEntidade, nomePlano, descricaoPlano, currentUser.Id, time.Now()).Scan(&planoId)
				if err != nil {
					panic(err.Error())
				}
			}
		}
		for key, value := range r.Form {
			if strings.HasPrefix(key, "cicloEntidade") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				cicloEntidadeId := 0
				cicloId := strings.Split(array[3], ":")[1]
				tipoMediaId := strings.Split(array[5], ":")[1]
				iniciaEm := strings.Split(array[8], ":")[1]
				terminaEm := strings.Split(array[9], ":")[1]
				snippet1 := ""
				snippet2 := ""
				if iniciaEm != "" {
					snippet1 = ", inicia_em "
					snippet2 = ", $6"
				}
				if terminaEm != "" {
					snippet1 = snippet1 + ", termina_em "
					snippet2 = snippet2 + ", $7"
				}
				sqlStatement := "INSERT INTO public.ciclos_entidades ( " +
					" entidade_id, " +
					" ciclo_id, " +
					" tipo_media, " +
					" author_id, " +
					" criado_em " +
					snippet1 +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5" + snippet2 + ") RETURNING id"
				log.Println(sqlStatement)
				if iniciaEm != "" && terminaEm != "" {
					err = Db.QueryRow(sqlStatement, idEntidade, cicloId, tipoMediaId, currentUser.Id, time.Now(), iniciaEm, terminaEm).Scan(&cicloEntidadeId)
				} else {
					err = Db.QueryRow(sqlStatement, idEntidade, cicloId, tipoMediaId, currentUser.Id, time.Now()).Scan(&cicloEntidadeId)
				}
				if err != nil {
					panic(err.Error())
				}
			}
		}
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		entidadeId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE entidades SET nome=$1, descricao=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(nome, descricao, entidadeId)
		log.Println("UPDATE: Id: " + entidadeId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Planos
		var planosDB = ListPlanosByEntidadeId(entidadeId)
		var planosPage []mdl.Plano
		var planoPage mdl.Plano
		for key, value := range r.Form {
			if strings.HasPrefix(key, "plano") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				planoPage.Id, _ = strconv.ParseInt(id, 10, 64)
				planoPage.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
				nome := strings.Split(array[3], ":")[1]
				log.Println("nome -------- " + nome)
				planoPage.Nome = nome
				descricao := strings.Split(array[4], ":")[1]
				log.Println("descricao -------- " + descricao)
				planoPage.Descricao = descricao
				planosPage = append(planosPage, planoPage)
			}
		}
		if len(planosPage) < len(planosDB) {
			log.Println("Quantidade de Planos da Página: " + strconv.Itoa(len(planosPage)))
			if len(planosPage) == 0 {
				DeletePlanosByEntidadeId(entidadeId) //DONE
			} else {
				var diffDB []mdl.Plano = planosDB
				for n := range planosPage {
					if containsPlano(diffDB, planosPage[n]) {
						diffDB = removePlano(diffDB, planosPage[n])
					}
				}
				DeletePlanosHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Plano = planosPage
			for n := range planosDB {
				if containsPlano(diffPage, planosDB[n]) {
					diffPage = removePlano(diffPage, planosDB[n])
				}
			}
			var plano mdl.Plano
			planoId := 0
			// statusItemId := GetStartStatus("plano")
			for i := range diffPage {
				plano = diffPage[i]
				log.Println("Entidade Id: " + strconv.FormatInt(plano.EntidadeId, 10))
				sqlStatement := "INSERT INTO public.planos( " +
					" entidade_id, nome, descricao, author_id, criado_em ) " +
					" VALUES ($1, $2, $3, $4, $5) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(sqlStatement, plano.EntidadeId, plano.Nome, plano.Descricao, currentUser.Id, time.Now()).Scan(&planoId)
			}
		}
		UpdatePlanosHandler(planosPage, planosDB)

		// Ciclos Entidade
		var ciclosEntidadeDB = ListCiclosEntidadeByEntidadeId(entidadeId)
		var ciclosEntidadePage []mdl.CicloEntidade
		var cicloEntidadePage mdl.CicloEntidade
		for key, value := range r.Form {
			if strings.HasPrefix(key, "cicloEntidade") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				cicloEntidadePage.Id, _ = strconv.ParseInt(id, 10, 64)
				cicloEntidadePage.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
				cicloId := strings.Split(array[3], ":")[1]
				log.Println("cicloId -------- " + cicloId)
				cicloEntidadePage.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
				nome := strings.Split(array[4], ":")[1]
				log.Println("nome -------- " + nome)
				cicloEntidadePage.Nome = nome
				tipoMediaId := strings.Split(array[5], ":")[1]
				log.Println("tipoMediaId -------- " + tipoMediaId)
				cicloEntidadePage.TipoMediaId, _ = strconv.Atoi(tipoMediaId)
				tipoMedia := strings.Split(array[6], ":")[1]
				log.Println("tipoMedia -------- " + tipoMedia)
				cicloEntidadePage.TipoMedia = tipoMedia
				nota := strings.Split(array[7], ":")[1]
				log.Println("nota -------- " + nota)
				cicloEntidadePage.Nota = nota
				iniciaEm := strings.Split(array[8], ":")[1]
				log.Println("iniciaEm -------- " + iniciaEm)
				cicloEntidadePage.IniciaEm = iniciaEm
				terminaEm := strings.Split(array[9], ":")[1]
				log.Println("terminaEm -------- " + terminaEm)
				cicloEntidadePage.TerminaEm = terminaEm
				authorId := strings.Split(array[10], ":")[1]
				log.Println("authorId -------- " + authorId)
				cicloEntidadePage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[11], ":")[1]
				log.Println("authorName -------- " + authorName)
				cicloEntidadePage.AuthorName = authorName
				criadoEm := strings.Split(array[12], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				cicloEntidadePage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[13], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				cicloEntidadePage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[14], ":")[1]
				log.Println("idVersaoOrigem -------- " + statusId)
				cicloEntidadePage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[15], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				cicloEntidadePage.CStatus = cStatus
				ciclosEntidadePage = append(ciclosEntidadePage, cicloEntidadePage)
			}
		}
		if len(ciclosEntidadePage) < len(ciclosEntidadeDB) {
			log.Println("Quantidade de Ciclos da Entidade da Página: " + strconv.Itoa(len(ciclosEntidadePage)))
			if len(ciclosEntidadePage) == 0 {
				DeleteCiclosEntidadeByEntidadeId(entidadeId) //DONE
			} else {
				var diffDB []mdl.CicloEntidade = ciclosEntidadeDB
				for n := range ciclosEntidadePage {
					if containsCicloEntidade(diffDB, ciclosEntidadePage[n]) {
						diffDB = removeCicloEntidade(diffDB, ciclosEntidadePage[n])
					}
				}
				DeleteCiclosEntidadeHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.CicloEntidade = ciclosEntidadePage
			for n := range ciclosEntidadeDB {
				if containsCicloEntidade(diffPage, ciclosEntidadeDB[n]) {
					diffPage = removeCicloEntidade(diffPage, ciclosEntidadeDB[n])
				}
			}
			var cicloEntidade mdl.CicloEntidade
			cicloEntidadeId := 0
			// statusItemId := GetStartStatus("plano")
			for i := range diffPage {
				cicloEntidade = diffPage[i]
				log.Println("Entidade Id: " + entidadeId)
				snippet1 := ""
				snippet2 := ""
				if cicloEntidade.IniciaEm != "" {
					snippet1 = ", inicia_em "
					snippet2 = ", $6"
				}
				if cicloEntidade.TerminaEm != "" {
					snippet1 = snippet1 + ", termina_em "
					snippet2 = snippet2 + ", $7"
				}
				sqlStatement := "INSERT INTO public.ciclos_entidades ( " +
					" entidade_id, " +
					" ciclo_id, " +
					" tipo_media, " +
					" author_id, " +
					" criado_em " +
					snippet1 +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5" + snippet2 + ") RETURNING id"
				log.Println(sqlStatement)
				if cicloEntidade.IniciaEm != "" && cicloEntidade.TerminaEm != "" {
					err = Db.QueryRow(sqlStatement, entidadeId, cicloEntidade.CicloId, cicloEntidade.TipoMediaId, currentUser.Id, time.Now(), cicloEntidade.IniciaEm, cicloEntidade.TerminaEm).Scan(&cicloEntidadeId)
				} else {
					err = Db.QueryRow(sqlStatement, entidadeId, cicloEntidade.CicloId, cicloEntidade.TipoMediaId, currentUser.Id, time.Now()).Scan(&cicloEntidadeId)
				}
			}
		}
		UpdateCiclosEntidadeHandler(ciclosEntidadePage, ciclosEntidadeDB)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteEntidadeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Entidade")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM planos WHERE entidade_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM ciclos_entidades WHERE entidade_id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM entidades WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.EntidadesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListEntidadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Entidades")
	if sec.IsAuthenticated(w, r) {
		var page mdl.PageEntidades
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
			" FROM entidades a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		rows, _ := Db.Query(sql)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Nome,
				&entidade.Descricao,
				&entidade.AuthorId,
				&entidade.AuthorName,
				&entidade.C_CriadoEm,
				&entidade.CStatus,
				&entidade.StatusId,
				&entidade.IdVersaoOrigem)
			entidade.Order = i
			i++
			entidades = append(entidades, entidade)
		}
		page.Entidades = entidades
		sql = "SELECT id, nome FROM ciclos ORDER BY id asc"
		rows, _ = Db.Query(sql)
		var ciclos []mdl.Ciclo
		var ciclo mdl.Ciclo
		i = 1
		for rows.Next() {
			rows.Scan(&ciclo.Id, &ciclo.Nome)
			ciclo.Order = i
			i++
			ciclos = append(ciclos, ciclo)
		}
		page.Ciclos = ciclos
		page.AppName = mdl.AppName
		page.Title = "Entidades"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/entidades/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadPlanosByEntidadeId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Planos By Entidade Id")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	log.Println("entidadeId: " + entidadeId)
	planos := ListPlanosByEntidadeId(entidadeId)
	jsonPlanos, _ := json.Marshal(planos)
	w.Write([]byte(jsonPlanos))
	log.Println("JSON Planos de Entidades")
}

func LoadCiclosByEntidadeId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Ciclos Entidades By Entidade Id")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	log.Println("entidadeId: " + entidadeId)
	ciclosEntidade := ListCiclosEntidadeByEntidadeId(entidadeId)
	jsonCiclosEntidade, _ := json.Marshal(ciclosEntidade)
	w.Write([]byte(jsonCiclosEntidade))
	log.Println("JSON Ciclos de Entidades")
}
