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

func CreateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "INSERT INTO ciclos(nome, descricao, author_id, criado_em) VALUES ($1, $2, $3, $4) RETURNING id"
		idCiclo := 0
		Db.QueryRow(sqlStatement, nome, descricao, currentUser.Id, time.Now()).Scan(&idCiclo)
		log.Println(sqlStatement + " - " + nome)
		log.Println("INSERT: Id: " + strconv.Itoa(idCiclo) + " - Nome: " + nome)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "pilarCiclo") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				pilarCicloId := 0
				pilarId := strings.Split(array[3], ":")[1]
				tipoMediaId := strings.Split(array[5], ":")[1]
				pesoPadrao := strings.Split(array[7], ":")[1]
				sqlStatement := " INSERT INTO " +
					" public.pilares_ciclos( " +
					" ciclo_id, " +
					" pilar_id, " +
					" tipo_media, " +
					" peso_padrao, " +
					" author_id, " +
					" criado_em ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(
					sqlStatement,
					idCiclo,
					pilarId,
					tipoMediaId,
					pesoPadrao,
					currentUser.Id,
					time.Now()).Scan(&pilarCicloId)
				if err != nil {
					panic(err.Error())
				}
			}
		}
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func IniciarCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Iniciar Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		r.ParseForm()
		entidades := r.Form["Entidades"]
		cicloId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		iniciaEm := r.FormValue("IniciaEm")
		terminaEm := r.FormValue("TerminaEm")
		sqlStatement := "UPDATE ciclos SET nome = $1, " +
			" descricao = $2 " +
			" WHERE id = $3 "
		updtForm, _ := Db.Prepare(sqlStatement)
		updtForm.Exec(nome, descricao, cicloId)
		log.Println("UPDATE: Id: " + cicloId + " | Nome: " + nome + " | Descrição: " + descricao)
		log.Println(len(entidades))
		for _, idEntidade := range entidades {
			cicloEntidadeId := 0
			snippet1 := ""
			snippet2 := ""
			if iniciaEm != "" {
				snippet1 = ", inicia_em "
				snippet2 = ", $5"
			}
			if terminaEm != "" {
				snippet1 = snippet1 + ", termina_em "
				snippet2 = snippet2 + ", $6"
			}
			sqlStatement := "INSERT INTO public.ciclos_entidades ( " +
				" entidade_id, " +
				" ciclo_id, " +
				" tipo_media, " +
				" author_id, " +
				" criado_em " +
				snippet1 +
				" ) " +
				" VALUES ($1, $2, 1, $3, $4 " + snippet2 + ") RETURNING id"
			log.Println(sqlStatement)
			log.Println("idEntidade: " + idEntidade)
			log.Println("cicloId: " + cicloId)
			log.Println("currentUser.Id: " + strconv.FormatInt(currentUser.Id, 10))
			log.Println("iniciaEm: " + iniciaEm)
			log.Println("terminaEm: " + terminaEm)
			var err error
			if iniciaEm != "" && terminaEm != "" {
				err = Db.QueryRow(sqlStatement, idEntidade, cicloId, currentUser.Id, time.Now(), iniciaEm, terminaEm).Scan(&cicloEntidadeId)
			} else {
				err = Db.QueryRow(sqlStatement, idEntidade, cicloId, currentUser.Id, time.Now()).Scan(&cicloEntidadeId)
			}
			if err != nil {
				panic(err.Error())
			}
		}
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		cicloId := r.FormValue("Id")
		nome := r.FormValue("Nome")
		descricao := r.FormValue("Descricao")
		sqlStatement := "UPDATE ciclos SET nome = $1, " +
			" descricao = $2 " +
			" WHERE id = $3 "
		updtForm, _ := Db.Prepare(sqlStatement)
		updtForm.Exec(nome, descricao, cicloId)
		log.Println("UPDATE: Id: " + cicloId + " | Nome: " + nome + " | Descrição: " + descricao)

		// Pilares Ciclos
		var pilaresCicloDB = ListPilaresByCicloId(cicloId)
		var pilaresCicloPage []mdl.PilarCiclo
		var pilarCicloPage mdl.PilarCiclo
		for key, value := range r.Form {
			if strings.HasPrefix(key, "pilarCiclo") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				pilarCicloPage.Id, _ = strconv.ParseInt(id, 10, 64)
				pilarCicloPage.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
				pilarId := strings.Split(array[3], ":")[1]
				log.Println("pilarId -------- " + pilarId)
				pilarCicloPage.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
				pilarNome := strings.Split(array[4], ":")[1]
				log.Println("pilarNome -------- " + pilarNome)
				pilarCicloPage.PilarNome = pilarNome
				tipoMediaId := strings.Split(array[5], ":")[1]
				log.Println("tipoMediaId -------- " + tipoMediaId)
				pilarCicloPage.TipoMediaId, _ = strconv.Atoi(tipoMediaId)

				tipoMedia := strings.Split(array[6], ":")[1]
				log.Println("tipoMedia -------- " + tipoMedia)
				pilarCicloPage.TipoMedia = tipoMedia

				pesoPadrao := strings.Split(array[7], ":")[1]
				log.Println("pesoPadrao -------- " + pesoPadrao)
				pilarCicloPage.PesoPadrao = pesoPadrao

				authorId := strings.Split(array[8], ":")[1]
				log.Println("authorId -------- " + authorId)
				pilarCicloPage.AuthorId, _ = strconv.ParseInt(authorId, 10, 64)
				authorName := strings.Split(array[9], ":")[1]
				log.Println("authorName -------- " + authorName)
				pilarCicloPage.AuthorName = authorName
				criadoEm := strings.Split(array[10], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				pilarCicloPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[11], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				pilarCicloPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[12], ":")[1]
				log.Println("StatusId -------- " + statusId)
				pilarCicloPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[13], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				pilarCicloPage.CStatus = cStatus
				pilaresCicloPage = append(pilaresCicloPage, pilarCicloPage)
			}
		}
		if len(pilaresCicloPage) < len(pilaresCicloDB) {
			log.Println("Quantidade de Pilares do Ciclo da Página: " + strconv.Itoa(len(pilaresCicloPage)))
			if len(pilaresCicloPage) == 0 {
				DeletePilaresCicloByCicloId(cicloId) //DONE
			} else {
				var diffDB []mdl.PilarCiclo = pilaresCicloDB
				for n := range pilaresCicloPage {
					if containsPilarCiclo(diffDB, pilaresCicloPage[n]) {
						diffDB = removePilarCiclo(diffDB, pilaresCicloPage[n])
					}
				}
				DeletePilaresCicloHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.PilarCiclo = pilaresCicloPage
			for n := range pilaresCicloDB {
				if containsPilarCiclo(diffPage, pilaresCicloDB[n]) {
					diffPage = removePilarCiclo(diffPage, pilaresCicloDB[n])
				}
			}
			var pilarCiclo mdl.PilarCiclo
			pilarCicloId := 0
			// statusItemId := GetStartStatus("plano")
			for i := range diffPage {
				pilarCiclo = diffPage[i]
				log.Println("Ciclo Id: " + cicloId)
				sqlStatement := "INSERT INTO public.pilares_ciclos ( " +
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
					cicloId,
					pilarCiclo.PilarId,
					pilarCiclo.TipoMediaId,
					pilarCiclo.PesoPadrao,
					currentUser.Id,
					time.Now()).Scan(&pilarCicloId)
			}
		}
		UpdatePilaresCicloHandler(pilaresCicloPage, pilaresCicloDB)

		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteCicloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Ciclo")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM pilares_ciclos WHERE ciclo_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)

		sqlStatement = "DELETE FROM ciclos WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.CiclosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListCiclosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Ciclos")
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
			" FROM ciclos a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ := Db.Query(sql)
		var ciclos []mdl.Ciclo
		var ciclo mdl.Ciclo
		var i = 1
		for rows.Next() {
			rows.Scan(
				&ciclo.Id,
				&ciclo.Nome,
				&ciclo.Descricao,
				&ciclo.AuthorId,
				&ciclo.AuthorName,
				&ciclo.C_CriadoEm,
				&ciclo.CStatus,
				&ciclo.StatusId,
				&ciclo.IdVersaoOrigem)
			ciclo.Order = i
			i++
			ciclos = append(ciclos, ciclo)
		}
		sql = "SELECT " +
			" a.id, " +
			" a.nome " +
			" FROM pilares a " +
			" order by a.id asc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var pilares []mdl.Pilar
		var pilar mdl.Pilar
		i = 1
		for rows.Next() {
			rows.Scan(
				&pilar.Id,
				&pilar.Nome)
			pilar.Order = i
			i++
			pilares = append(pilares, pilar)
		}
		sql = "SELECT a.id, a.sigla, a.codigo, a.nome " +
			"FROM entidades a " +
			"WHERE NOT EXISTS " +
			"(SELECT 1 FROM ciclos_entidades b " +
			" WHERE b.entidade_id = a.id) " +
			"ORDER BY a.sigla"
		log.Println(sql)
		rows, _ = Db.Query(sql)
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
		}
		var page mdl.PageCiclos
		page.Pilares = pilares
		page.Entidades = entidades
		page.Ciclos = ciclos
		page.AppName = mdl.AppName
		page.Title = "Ciclos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/ciclos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Ciclos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func GetNow() time.Time {
	br, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), 0, br)
	strNow := now.String()
	log.Println("------------------------------------------------------------------------")
	log.Println("- Agora são " + strNow + " em America/Sao_Paulo. ")
	log.Println("------------------------------------------------------------------------")
	txtNow := strings.Split(strings.Split(strings.Split(strNow, " ")[1], ".")[0], ":")
	hora, _ := strconv.Atoi(txtNow[0])
	minuto, _ := strconv.Atoi(txtNow[1])
	segundo, _ := strconv.Atoi(txtNow[2])
	t := time.Date(0000, time.January, 1,
		hora,
		minuto,
		segundo, 0, time.UTC)
	return t
}

func LoadPilaresByCicloId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Pilares Ciclos By Ciclo Id")
	r.ParseForm()
	var cicloId = r.FormValue("cicloId")
	log.Println("cicloId: " + cicloId)
	pilaresCiclo := ListPilaresByCicloId(cicloId)
	jsonPilaresCiclo, _ := json.Marshal(pilaresCiclo)
	w.Write([]byte(jsonPilaresCiclo))
	log.Println("JSON Pilares de Ciclos")
}
