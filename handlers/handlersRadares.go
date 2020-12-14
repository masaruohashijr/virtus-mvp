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

func CreateRadarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Radar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		referencia := r.FormValue("Referencia")
		dataRadar := r.FormValue("DataRadar")
		sqlStatement := "INSERT INTO radares(" +
			" nome, descricao, referencia, data_radar, author_id, criado_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
		idRadar := 0
		Db.QueryRow(sqlStatement, nome, descricao, referencia, dataRadar, currentUser.Id, time.Now()).Scan(&idRadar)
		log.Println(sqlStatement + " - " + nome)
		log.Println("INSERT: Id: " + strconv.Itoa(idRadar) + " - Nome: " + nome)

		for key, value := range r.Form {
			if strings.HasPrefix(key, "anotacaoRadar") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				anotacaoRadarId := 0
				anotacaoId := strings.Split(array[3], ":")[1]
				observacoes := strings.Split(array[4], ":")[1]
				registroAta := strings.Split(array[5], ":")[1]
				sqlStatement := " INSERT INTO " +
					" anotacoes_radares( " +
					" radar_id, " +
					" anotacao_id, " +
					" observacoes, " +
					" registro_ata, " +
					" author_id, " +
					" criado_em ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(
					sqlStatement,
					idRadar,
					anotacaoId,
					observacoes,
					registroAta,
					currentUser.Id,
					time.Now()).Scan(&anotacaoRadarId)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
		http.Redirect(w, r, route.RadaresRoute+"?msg=Radar criado com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateRadarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Radar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		//currentUser := GetUserInCookie(w, r)
		radarId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		dataRadar := r.FormValue("DataRadar")
		descricao := r.FormValue("Descricao")
		referencia := r.FormValue("Referencia")
		sqlStatement := "UPDATE radares SET nome = $1, " +
			" dataRadar = $2, " +
			" descricao = $3, " +
			" referencia = $4 " +
			" WHERE id = $5 "
		updtForm, _ := Db.Prepare(sqlStatement)
		updtForm.Exec(nome, dataRadar, descricao, referencia, radarId)
		log.Println("UPDATE: Id: " + radarId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Questoes Radares
		/*var questoesRadarDB = ListQuestoesByRadarId(radarId)
		var questoesRadarPage []mdl.PilarRadar
		var pilarRadarPage mdl.PilarRadar
		for key, value := range r.Form {
			if strings.HasPrefix(key, "pilarRadar") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				pilarRadarPage.Id, _ = strconv.ParseInt(id, 10, 64)
				pilarRadarPage.RadarId, _ = strconv.ParseInt(radarId, 10, 64)
				pilarId := strings.Split(array[3], ":")[1]
				log.Println("pilarId -------- " + pilarId)
				pilarRadarPage.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
				pilarNome := strings.Split(array[4], ":")[1]
				log.Println("pilarNome -------- " + pilarNome)
				pilarRadarPage.PilarNome = pilarNome
				tipoMediaId := strings.Split(array[5], ":")[1]
				log.Println("tipoMediaId -------- " + tipoMediaId)
				pilarRadarPage.TipoMediaId, _ = strconv.Atoi(tipoMediaId)
				tipoMedia := strings.Split(array[6], ":")[1]
				log.Println("tipoMedia -------- " + tipoMedia)
				pilarRadarPage.TipoMedia = tipoMedia
				pesoPadrao := strings.Split(array[7], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				pilarRadarPage.PesoPadrao = pesoPadrao
				authorId := strings.Split(array[8], ":")[1]
				log.Println("authorId -------- " + authorId)
				pilarRadarPage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[9], ":")[1]
				log.Println("authorName -------- " + authorName)
				pilarRadarPage.AuthorName = authorName
				criadoEm := strings.Split(array[10], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				pilarRadarPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[11], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				pilarRadarPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[12], ":")[1]
				log.Println("StatusId -------- " + statusId)
				pilarRadarPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[13], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				pilarRadarPage.CStatus = cStatus
				questoesRadarPage = append(questoesRadarPage, pilarRadarPage)
			}
		}
		if len(questoesRadarPage) < len(questoesRadarDB) {
			log.Println("Quantidade de Questoes do Radar da Página: " + strconv.Itoa(len(questoesRadarPage)))
			if len(questoesRadarPage) == 0 {
				DeleteQuestoesRadarByRadarId(radarId) //DONE
			} else {
				var diffDB []mdl.PilarRadar = questoesRadarDB
				for n := range questoesRadarPage {
					if containsPilarRadar(diffDB, questoesRadarPage[n]) {
						diffDB = removePilarRadar(diffDB, questoesRadarPage[n])
					}
				}
				DeleteQuestoesRadarHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.PilarRadar = questoesRadarPage
			for n := range questoesRadarDB {
				if containsPilarRadar(diffPage, questoesRadarDB[n]) {
					diffPage = removePilarRadar(diffPage, questoesRadarDB[n])
				}
			}
			var pilarRadar mdl.PilarRadar
			pilarRadarId := 0
			// statusItemId := GetStartStatus("plano")
			for i := range diffPage {
				pilarRadar = diffPage[i]
				log.Println("Radar Id: " + radarId)
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
					pilarRadar.PilarId,
					pilarRadar.TipoMediaId,
					pilarRadar.PesoPadrao,
					currentUser.Id,
					time.Now()).Scan(&pilarRadarId)
			}
		}
		UpdateQuestoesRadarHandler(questoesRadarPage, questoesRadarDB)*/

		http.Redirect(w, r, route.RadaresRoute+"?msg=Radar atualizado com sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteRadarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Radar")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		errMsg := "O Radar está associado a um registro e não pôde ser removido."
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM radares WHERE id=$1"
		log.Println(sqlStatement)
		deleteForm, _ := Db.Prepare(sqlStatement)
		_, err := deleteForm.Exec(id)
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			log.Println("ENTROU NO ERRO " + errMsg)
			http.Redirect(w, r, route.RadaresRoute+"?errMsg="+errMsg, 301)
		} else {
			http.Redirect(w, r, route.RadaresRoute+"?msg=Radar removido com sucesso.", 301)
		}
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListRadaresHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Radares")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listRadares") {
		msg := r.FormValue("msg")
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, " +
			" a.nome, " +
			" a.descricao, " +
			" a.referencia, " +
			" to_char(a.data_radar,'DD/MM/YYYY'), " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM radares a " +
			" LEFT JOIN users b ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var radares []mdl.Radar
		var radar mdl.Radar
		var i = 1
		for rows.Next() {
			rows.Scan(
				&radar.Id,
				&radar.Nome,
				&radar.Descricao,
				&radar.Referencia,
				&radar.DataRadar,
				&radar.AuthorId,
				&radar.AuthorName,
				&radar.C_CriadoEm,
				&radar.CStatus,
				&radar.StatusId,
				&radar.IdVersaoOrigem)
			radar.Order = i
			i++
			radares = append(radares, radar)
		}
		sql = "SELECT " +
			" a.id, " +
			" a.assunto, " +
			" a.entidade_id, " +
			" d.sigla as entidade_sigla, " +
			" case when a.risco = 'A' then 'Baixo' when a.risco = 'M' then 'Médio' else 'Baixo' end, " +
			" case when a.tendencia = 'M' then 'Melhora' when a.tendencia = 'E' then 'Estabilidade' else 'Piora' end, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM anotacoes a " +
			" LEFT JOIN users b ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" INNER JOIN entidades d ON a.entidade_id = d.id " +
			" ORDER BY a.id asc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var anotacoes []mdl.Anotacao
		var anotacao mdl.Anotacao
		for rows.Next() {
			rows.Scan(&anotacao.Id,
				&anotacao.Assunto,
				&anotacao.EntidadeId,
				&anotacao.EntidadeSigla,
				&anotacao.Risco,
				&anotacao.Tendencia,
				&anotacao.AuthorId,
				&anotacao.AuthorName,
				&anotacao.CriadoEm,
				&anotacao.CStatus,
				&anotacao.StatusId,
				&anotacao.IdVersaoOrigem)
			anotacoes = append(anotacoes, anotacao)
		}
		var page mdl.PageRadares
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		if msg != "" {
			page.Msg = msg
		}
		page.Radares = radares
		page.Anotacoes = anotacoes
		page.AppName = mdl.AppName
		page.Title = "Radares"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/radares/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Radares", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadQuestoesByRadarId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Questoes Radares By Radar Id")
	r.ParseForm()
	var radarId = r.FormValue("radarId")
	log.Println("radarId: " + radarId)
	/*questoesRadar := ListQuestoesByRadarId(radarId)
	jsonQuestoesRadar, _ := json.Marshal(questoesRadar)
	w.Write([]byte(jsonQuestoesRadar))*/
	log.Println("JSON Questoes de Radares")
}
