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

func CreateEscritorioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Escritorio")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		chefe := r.FormValue("Chefe")
		sqlStatement := "INSERT INTO escritorios(nome, descricao, chefe_id, author_id, criado_em) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome, descricao, chefe, currentUser.Id, time.Now()).Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome + " | Descrição: " + descricao)
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
		chefe := r.FormValue("Chefe")
		sqlStatement := "UPDATE escritorios SET nome=$1, descricao=$2, chefe_id=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, chefe, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome + " | Descrição: " + descricao + " | Chefe: " + chefe)
		http.Redirect(w, r, route.EscritoriosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateJurisdicaoEscritorioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Jurisdicao")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		escritorioId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		chefe := r.FormValue("Chefe")
		sqlStatement := "UPDATE escritorios SET nome=$1, descricao=$2, chefe_id=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(nome, descricao, chefe, escritorioId)
		log.Println("UPDATE: Id: " + escritorioId + " | Nome: " + nome + " | Descrição: " + descricao + " | Chefe: " + chefe)

		// Jurisdições
		var jurisdicoesDB = ListJurisdicoesByEscritorioId(escritorioId)
		var jurisdicoesPage []mdl.Jurisdicao
		var jurisdicaoPage mdl.Jurisdicao
		for key, value := range r.Form {
			if strings.HasPrefix(key, "jurisdicao") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				jurisdicaoPage.Id, _ = strconv.ParseInt(id, 10, 64)
				jurisdicaoPage.EscritorioId, _ = strconv.ParseInt(escritorioId, 10, 64)
				entidadeId := strings.Split(array[3], ":")[1]
				log.Println("entidadeId -------- " + entidadeId)
				jurisdicaoPage.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
				entidadeNome := strings.Split(array[4], ":")[1]
				log.Println("entidadeNome -------- " + entidadeNome)
				jurisdicaoPage.EntidadeNome = entidadeNome
				iniciaEm := strings.Split(array[5], ":")[1]
				log.Println("iniciaEm -------- " + iniciaEm)
				jurisdicaoPage.IniciaEm = iniciaEm
				terminaEm := strings.Split(array[6], ":")[1]
				log.Println("terminaEm -------- " + terminaEm)
				jurisdicaoPage.TerminaEm = terminaEm
				autorId := strings.Split(array[7], ":")[1]
				log.Println("autorId -------- " + autorId)
				jurisdicaoPage.AuthorId, _ = strconv.ParseInt(autorId, 10, 64)
				autorNome := strings.Split(array[8], ":")[1]
				log.Println("autorNome -------- " + autorNome)
				jurisdicaoPage.AuthorName = autorNome
				criadoEm := strings.Split(array[9], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				jurisdicaoPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[10], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				jurisdicaoPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[11], ":")[1]
				log.Println("statusId -------- " + statusId)
				jurisdicaoPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[12], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				jurisdicaoPage.CStatus = cStatus
				jurisdicoesPage = append(jurisdicoesPage, jurisdicaoPage)
			}
		}
		if len(jurisdicoesPage) < len(jurisdicoesDB) {
			log.Println("Quantidade de Entidades do Escritório na Página: " + strconv.Itoa(len(jurisdicoesPage)))
			if len(jurisdicoesPage) == 0 {
				DeleteJurisdicoesByEscritorioId(escritorioId) //DONE
			} else {
				var diffDB []mdl.Jurisdicao = jurisdicoesDB
				for n := range jurisdicoesPage {
					if containsJurisdicao(diffDB, jurisdicoesPage[n]) {
						diffDB = removeJurisdicao(diffDB, jurisdicoesPage[n])
					}
				}
				DeleteJurisdicoesHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Jurisdicao = jurisdicoesPage
			for n := range jurisdicoesDB {
				if containsJurisdicao(diffPage, jurisdicoesDB[n]) {
					diffPage = removeJurisdicao(diffPage, jurisdicoesDB[n])
				}
			}
			var jurisdicao mdl.Jurisdicao
			jurisdicaoId := 0
			statusJurisdicaoId := GetStartStatus("jurisdicao")
			for i := range diffPage {
				jurisdicao = diffPage[i]
				log.Println("Escritório Id: " + escritorioId)
				sqlStatement := "INSERT INTO jurisdicoes ( " +
					" escritorio_id, " +
					" entidade_id, " +
					" author_id, " +
					" criado_em, " +
					" status_id " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(
					sqlStatement,
					escritorioId,
					jurisdicao.EntidadeId,
					currentUser.Id,
					time.Now(),
					statusJurisdicaoId).Scan(&jurisdicaoId)
				if err != nil {
					log.Println("ERRO AO INSERIR JURISDICAO")
					log.Println(err)
				}
				log.Println("De " + jurisdicao.IniciaEm + " a " + jurisdicao.TerminaEm)
				if jurisdicao.IniciaEm != "" {
					sqlStatement = "UPDATE jurisdicoes SET inicia_em = to_date('" +
						jurisdicao.IniciaEm + "','YYYY-MM-DD') " +
						"WHERE id = " + strconv.Itoa(jurisdicaoId)
					_, err = Db.Exec(sqlStatement)
					if err != nil {
						log.Println(err)
					}
				}
				if jurisdicao.TerminaEm != "" {
					sqlStatement = "UPDATE jurisdicoes SET termina_em = to_date('" +
						jurisdicao.TerminaEm + "','YYYY-MM-DD') " +
						"WHERE id = " + strconv.Itoa(jurisdicaoId)
					_, err = Db.Exec(sqlStatement)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
		UpdateJurisdicoesHandler(jurisdicoesPage, jurisdicoesDB)

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
