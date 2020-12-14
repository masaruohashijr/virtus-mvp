package handlers

import (
	//"encoding/json"
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

func CreateVersaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Versão")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		objetivo := r.FormValue("Objetivo")
		definicaoPronto := r.FormValue("DefinicaoPronto")
		iniciaEm := r.FormValue("IniciaEm")
		terminaEm := r.FormValue("TerminaEm")
		sqlStatement := "INSERT INTO versoes(" +
			" nome, " +
			" objetivo, " +
			" definicao_pronto, " +
			" inicia_em, " +
			" termina_em, " +
			" author_id, " +
			" criado_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
		idVersao := 0
		row := Db.QueryRow(sqlStatement,
			nome,
			objetivo,
			definicaoPronto,
			iniciaEm,
			terminaEm,
			currentUser.Id,
			time.Now())
		err := row.Scan(&idVersao)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(sqlStatement + " - " + nome)
		log.Println("INSERT: Id: " + strconv.Itoa(idVersao) + " - Nome: " + nome)
		/*
			for key, value := range r.Form {
				if strings.HasPrefix(key, "questaoVersão") {
					array := strings.Split(value[0], "#")
					log.Println(value[0])
					questaoVersaoId := 0
					questaoId := strings.Split(array[3], ":")[1]
					sqlStatement := " INSERT INTO " +
						" public.questoes_radares( " +
						" radar_id, " +
						" questao_id, " +
						" registro_ata, " +
						" author_id, " +
						" criado_em ) " +
						" VALUES ($1, $2, $3, $4, $5) RETURNING id"
					log.Println(sqlStatement)
					err := Db.QueryRow(
						sqlStatement,
						idVersao,
						questaoId,
						tipoMediaId,
						pesoPadrao,
						currentUser.Id,
						time.Now()).Scan(&questaoVersaoId)
					if err != nil {
						log.Println(err.Error())
					}
				}
			}*/
		http.Redirect(w, r, route.VersoesRoute+"?msg=Versão criado com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateVersaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Versão")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		//currentUser := GetUserInCookie(w, r)
		versaoId := r.FormValue("Id")
		log.Println(versaoId)
		nome := r.FormValue("Nome")
		log.Println(nome)
		objetivo := r.FormValue("Objetivo")
		log.Println(objetivo)
		definicaoPronto := r.FormValue("DefinicaoPronto")
		log.Println(definicaoPronto)
		iniciaEm := r.FormValue("IniciaEm")
		log.Println(iniciaEm)
		terminaEm := r.FormValue("TerminaEm")
		log.Println(terminaEm)
		sqlStatement := "UPDATE versoes SET " +
			" nome = $1, " +
			" objetivo = $2, " +
			" definicao_pronto = $3, " +
			" inicia_em = $4, " +
			" termina_em = $5 " +
			" WHERE id = $6 "
		log.Println(sqlStatement)
		updtForm, _ := Db.Prepare(sqlStatement)
		_, err := updtForm.Exec(nome, objetivo, definicaoPronto, iniciaEm, terminaEm, versaoId)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("UPDATE: Id: " + versaoId + " | Nome: " + nome + " | Objetivo: " + objetivo)

		// Chamados Versoes
		/*var chamadosVersaoDB = ListChamadosByVersaoId(radarId)
		var chamadosVersaoPage []mdl.ChamadoVersao
		var chamadoVersaoPage mdl.ChamadoVersao
		for key, value := range r.Form {
			if strings.HasPrefix(key, "chamadoVersao") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				pilarVersaoPage.Id, _ = strconv.ParseInt(id, 10, 64)
				pilarVersaoPage.VersaoId, _ = strconv.ParseInt(radarId, 10, 64)
				pilarId := strings.Split(array[3], ":")[1]
				log.Println("pilarId -------- " + pilarId)
				pilarVersaoPage.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
				pilarNome := strings.Split(array[4], ":")[1]
				log.Println("pilarNome -------- " + pilarNome)
				pilarVersaoPage.PilarNome = pilarNome
				tipoMediaId := strings.Split(array[5], ":")[1]
				log.Println("tipoMediaId -------- " + tipoMediaId)
				pilarVersaoPage.TipoMediaId, _ = strconv.Atoi(tipoMediaId)
				tipoMedia := strings.Split(array[6], ":")[1]
				log.Println("tipoMedia -------- " + tipoMedia)
				pilarVersaoPage.TipoMedia = tipoMedia
				pesoPadrao := strings.Split(array[7], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				pilarVersaoPage.PesoPadrao = pesoPadrao
				authorId := strings.Split(array[8], ":")[1]
				log.Println("authorId -------- " + authorId)
				pilarVersaoPage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[9], ":")[1]
				log.Println("authorName -------- " + authorName)
				pilarVersaoPage.AuthorName = authorName
				criadoEm := strings.Split(array[10], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				pilarVersaoPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[11], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				pilarVersaoPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[12], ":")[1]
				log.Println("StatusId -------- " + statusId)
				pilarVersaoPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[13], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				pilarVersaoPage.CStatus = cStatus
				questoesVersaoPage = append(questoesVersaoPage, pilarVersaoPage)
			}
		}
		if len(questoesVersaoPage) < len(questoesVersaoDB) {
			log.Println("Quantidade de Questoes do Versao da Página: " + strconv.Itoa(len(questoesVersaoPage)))
			if len(questoesVersaoPage) == 0 {
				DeleteQuestoesVersaoByVersaoId(radarId) //DONE
			} else {
				var diffDB []mdl.PilarVersao = questoesVersaoDB
				for n := range questoesVersaoPage {
					if containsPilarVersao(diffDB, questoesVersaoPage[n]) {
						diffDB = removePilarVersao(diffDB, questoesVersaoPage[n])
					}
				}
				DeleteQuestoesVersaoHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.PilarVersao = questoesVersaoPage
			for n := range questoesVersaoDB {
				if containsPilarVersao(diffPage, questoesVersaoDB[n]) {
					diffPage = removePilarVersao(diffPage, questoesVersaoDB[n])
				}
			}
			var pilarVersao mdl.PilarVersao
			pilarVersaoId := 0
			// statusItemId := GetStartStatus("plano")
			for i := range diffPage {
				pilarVersao = diffPage[i]
				log.Println("Versao Id: " + radarId)
				sqlStatement := "INSERT INTO questoes_radares ( " +
					" ciclo_id, " +
					" pilar_id, " +
					" tipo_media, " +
					" peso_padrao, " +
					" author_id, " +
					" criado_em " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(
					sqlStatement,
					radarId,
					pilarVersao.PilarId,
					pilarVersao.TipoMediaId,
					pilarVersao.PesoPadrao,
					currentUser.Id,
					time.Now()).Scan(&pilarVersaoId)
			}
		}
		UpdateQuestoesVersaoHandler(questoesVersaoPage, questoesVersaoDB)*/

		http.Redirect(w, r, route.VersoesRoute+"?msg=Versão atualizado com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteVersaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Versão")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		errMsg := "O Versão está associado a um registro e não pode ser removido."
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM radares WHERE id=$1"
		log.Println(sqlStatement)
		deleteForm, _ := Db.Prepare(sqlStatement)
		_, err := deleteForm.Exec(id)
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			log.Println("ENTROU NO ERRO " + errMsg)
			http.Redirect(w, r, route.VersoesRoute+"?errMsg="+errMsg, 301)
		} else {
			http.Redirect(w, r, route.VersoesRoute+"?msg=Versão removido com sucesso.", 301)
		}
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListVersoesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Versoes")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listVersoes") {
		msg := r.FormValue("msg")
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" coalesce(a.objetivo,''), " +
			" coalesce(a.definicao_pronto,''), " +
			" to_char(a.inicia_em,'DD/MM/YYYY'), " +
			" to_char(a.termina_em,'DD/MM/YYYY'), " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM versoes a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var versoes []mdl.Versao
		var versao mdl.Versao
		var i = 1
		for rows.Next() {
			rows.Scan(
				&versao.Id,
				&versao.Nome,
				&versao.Objetivo,
				&versao.DefinicaoPronto,
				&versao.IniciaEm,
				&versao.TerminaEm,
				&versao.AuthorId,
				&versao.AuthorName,
				&versao.C_CriadoEm,
				&versao.CStatus,
				&versao.StatusId,
				&versao.IdVersaoOrigem)
			versao.Order = i
			i++
			versoes = append(versoes, versao)
		}
		/*sql = "SELECT " +
			" a.id, " +
			" a.nome " +
			" FROM questoes a " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var questoes []mdl.Pilar
		var pilar mdl.Pilar
		i = 1
		for rows.Next() {
			rows.Scan(
				&pilar.Id,
				&pilar.Nome)
			pilar.Order = i
			i++
			questoes = append(questoes, pilar)
		}
		sql = "SELECT a.id, a.sigla, a.codigo, a.nome " +
			"FROM entidades a " +
			"WHERE NOT EXISTS " +
			"(SELECT 1 FROM radares_entidades b " +
			" WHERE b.entidade_id = a.id) " +
			"ORDER BY a.sigla"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Sigla,
				&entidade.Codigo,
				&entidade.Nome)
			entidade.Order = i
			i++
			entidades = append(entidades, entidade)
		}*/
		var page mdl.PageVersoes
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		if msg != "" {
			page.Msg = msg
		}
		page.Versoes = versoes
		page.AppName = mdl.AppName
		page.Title = "Versões"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/versoes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Versoes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadQuestoesByVersaoId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Questoes Versoes By Versao Id")
	r.ParseForm()
	var radarId = r.FormValue("radarId")
	log.Println("radarId: " + radarId)
	/*questoesVersao := ListQuestoesByVersaoId(radarId)
	jsonQuestoesVersao, _ := json.Marshal(questoesVersao)
	w.Write([]byte(jsonQuestoesVersao))*/
	log.Println("JSON Questoes de Versoes")
}
