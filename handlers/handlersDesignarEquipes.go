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
	sec "virtus/security"
)

func ListDesignarEquipesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Designar Equipes Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "designarEquipes") {
		log.Println("--------------")
		var page mdl.PageEntidadesCiclos
		sql := " SELECT b.entidade_id, c.nome " +
			" FROM escritorios a " +
			" LEFT JOIN jurisdicoes b ON a.id = b.escritorio_id " +
			" LEFT JOIN entidades c ON c.id = b.entidade_id " +
			" INNER JOIN ciclos_entidades d ON d.entidade_id = b.entidade_id " +
			" WHERE a.chefe_id = $1"
		log.Println(sql)
		rows, _ := Db.Query(sql, currentUser.Id)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Nome)
			entidade.Order = i
			i++
			sql = "SELECT b.id, b.nome " +
				" FROM ciclos_entidades a " +
				" LEFT JOIN ciclos b ON a.ciclo_id = b.id " +
				" WHERE a.entidade_id = $1 " +
				" ORDER BY id asc"
			rows, _ = Db.Query(sql, entidade.Id)
			var ciclosEntidade []mdl.CicloEntidade
			var cicloEntidade mdl.CicloEntidade
			i = 1
			for rows.Next() {
				rows.Scan(&cicloEntidade.Id, &cicloEntidade.Nome)
				cicloEntidade.Order = i
				i++
				ciclosEntidade = append(ciclosEntidade, cicloEntidade)
			}
			entidade.CiclosEntidade = ciclosEntidade
			entidades = append(entidades, entidade)
		}

		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,''), " +
			" coalesce(d.name,'') as role_name " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" LEFT JOIN roles d ON c.role_id = d.id " +
			" WHERE a.chefe_id = $1 AND c.role_id in (2,3,4) "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var membros []mdl.Membro
		var membro mdl.Membro
		i = 1
		for rows.Next() {
			rows.Scan(&membro.UsuarioId, &membro.UsuarioNome, &membro.UsuarioPerfil)
			membros = append(membros, membro)
		}
		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,''), " +
			" coalesce(d.name,'') as role_name " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" LEFT JOIN roles d ON c.role_id = d.id " +
			" WHERE a.chefe_id = $1 AND c.role_id in (3) "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var supervisores []mdl.User
		var supervisor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name, &supervisor.RoleName)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores
		page.Membros = membros
		page.Entidades = entidades
		page.AppName = mdl.AppName
		page.Title = "Designar Equipes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/designarequipes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Designar-Equipes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateDesignarEquipeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("************ Update Designar Equipe ************")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		supervisorId := r.FormValue("SupervisorId")

		if supervisorId != "" {
			sqlStatement := "UPDATE ciclos_entidades SET supervisor_id=$1 WHERE entidade_id=$2 AND ciclo_id=$3 " +
				" AND NOT EXISTS (SELECT 1 FROM ciclos_entidades WHERE supervisor_id=$4 AND entidade_id=$5 AND ciclo_id=$6)"
			log.Println(sqlStatement)
			updtForm, _ := Db.Prepare(sqlStatement)
			updtForm.Exec(supervisorId, entidadeId, cicloId, supervisorId, entidadeId, cicloId)
			sqlStatement = "INSERT INTO integrantes ( " +
				" entidade_id, " +
				" ciclo_id, " +
				" usuario_id, " +
				" author_id, " +
				" criado_em " +
				" ) " +
				" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS " +
				" (SELECT 1 FROM integrantes WHERE entidade_id = $6 AND ciclo_id = $7 AND usuario_id =$8)"
			log.Println(sqlStatement)
			Db.QueryRow(
				sqlStatement,
				entidadeId,
				cicloId,
				supervisorId,
				GetUserInCookie(w, r).Id,
				time.Now(),
				entidadeId,
				cicloId,
				supervisorId)
		}
		log.Println("UPDATE: EntidadeId: " + entidadeId + " | CicloId: " + cicloId + " | SupervisorId: " + supervisorId)

		// Integrantes
		var integrantesDB = ListIntegrantesByEntidadeIdByCicloId(entidadeId, cicloId)
		var integrantesPage []mdl.Integrante
		var integrantePage mdl.Integrante
		for key, value := range r.Form {
			if strings.HasPrefix(key, "integrante") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				integrantePage.Id, _ = strconv.ParseInt(id, 10, 64)
				integrantePage.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
				integrantePage.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
				usuarioId := strings.Split(array[4], ":")[1]
				log.Println("usuarioId -------- " + usuarioId)
				integrantePage.UsuarioId, _ = strconv.ParseInt(usuarioId, 10, 64)
				usuarioNome := strings.Split(array[5], ":")[1]
				log.Println("usuarioNome -------- " + usuarioNome)
				integrantePage.UsuarioNome = usuarioNome
				usuarioPerfil := strings.Split(array[6], ":")[1]
				log.Println("usuarioPerfil -------- " + usuarioPerfil)
				integrantePage.UsuarioPerfil = usuarioPerfil
				iniciaEm := strings.Split(array[7], ":")[1]
				log.Println("iniciaEm -------- " + iniciaEm)
				integrantePage.IniciaEm = iniciaEm
				terminaEm := strings.Split(array[8], ":")[1]
				log.Println("terminaEm -------- " + terminaEm)
				integrantePage.TerminaEm = terminaEm
				autorId := strings.Split(array[9], ":")[1]
				log.Println("autorId -------- " + autorId)
				integrantePage.AuthorId, _ = strconv.ParseInt(autorId, 10, 64)
				autorNome := strings.Split(array[10], ":")[1]
				log.Println("autorNome -------- " + autorNome)
				integrantePage.AuthorName = autorNome
				criadoEm := strings.Split(array[11], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				integrantePage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[12], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				integrantePage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[13], ":")[1]
				log.Println("statusId -------- " + statusId)
				integrantePage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[14], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				integrantePage.CStatus = cStatus
				integrantesPage = append(integrantesPage, integrantePage)
			}
		}
		if len(integrantesPage) < len(integrantesDB) {
			log.Println("Quantidade de Integrante da EntidadeCiclo da Página: " + strconv.Itoa(len(integrantesPage)))
			if len(integrantesPage) == 0 {
				DeleteIntegrantesByEntidadeCicloId(entidadeId, cicloId) //DONE
			} else {
				var diffDB []mdl.Integrante = integrantesDB
				for n := range integrantesPage {
					if containsIntegrante(diffDB, integrantesPage[n]) {
						diffDB = removeIntegrante(diffDB, integrantesPage[n])
					}
				}
				DeleteIntegrantesHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Integrante = integrantesPage
			for n := range integrantesDB {
				if containsIntegrante(diffPage, integrantesDB[n]) {
					diffPage = removeIntegrante(diffPage, integrantesDB[n])
				}
			}
			var integrante mdl.Integrante
			integranteId := 0
			statusComponenteId := GetStartStatus("integrante")
			for i := range diffPage {
				integrante = diffPage[i]
				sqlStatement := "INSERT INTO integrantes ( " +
					" entidade_id, " +
					" ciclo_id, " +
					" usuario_id, " +
					" author_id, " +
					" criado_em, " +
					" status_id " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(
					sqlStatement,
					entidadeId,
					cicloId,
					integrante.UsuarioId,
					currentUser.Id,
					time.Now(),
					statusComponenteId).Scan(&integranteId)
			}
		}
		UpdateIntegrantesHandler(integrantesPage, integrantesDB)

		http.Redirect(w, r, "/listDesignarEquipes", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func containsIntegrante(membros []mdl.Integrante, membroCompared mdl.Integrante) bool {
	for n := range membros {
		if membros[n].Id == membroCompared.Id {
			return true
		}
	}
	return false
}

func removeIntegrante(membros []mdl.Integrante, membroToBeRemoved mdl.Integrante) []mdl.Integrante {
	var newIntegrantes []mdl.Integrante
	for i := range membros {
		if membros[i].Id != membroToBeRemoved.Id {
			newIntegrantes = append(newIntegrantes, membros[i])
		}
	}
	return newIntegrantes
}

func DeleteIntegrantesByEntidadeCicloId(entidadeId string, cicloId string) {
	sqlStatement := "DELETE FROM integrantes WHERE entidade_id=$1 AND ciclo_id = $2"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(entidadeId, cicloId)
	log.Println("DELETE integrantes in Entidade Id: " + entidadeId + " - " + cicloId)
}

func DeleteIntegrantesHandler(diffDB []mdl.Integrante) {
	sqlStatement := "DELETE FROM integrantes WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Integrante Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func UpdateIntegrantesHandler(integrantesPage []mdl.Integrante, integrantesDB []mdl.Integrante) {
	for i := range integrantesPage {
		id := integrantesPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range integrantesDB {
			log.Println("integrantesDB[j].Id: " + strconv.FormatInt(integrantesDB[j].Id, 10))
			if strconv.FormatInt(integrantesDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				fieldsChanged := hasSomeFieldChangedIntegrante(integrantesPage[i], integrantesDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateIntegranteHandler(integrantesPage[i], integrantesDB[j]) // TODO
				}
				integrantesDB = removeIntegrante(integrantesDB, integrantesPage[i]) // CORREÇÃO
				break
			}
		}
	}
	DeleteIntegrantesHandler(integrantesDB) // CORREÇÃO
}

func hasSomeFieldChangedIntegrante(integrantePage mdl.Integrante, integranteDB mdl.Integrante) bool {
	log.Println("integrantePage.Nome: " + integrantePage.UsuarioNome)
	log.Println("integranteDB.Nome: " + integranteDB.UsuarioNome)
	if integrantePage.IniciaEm != integranteDB.IniciaEm {
		return true
	} else if integrantePage.TerminaEm != integranteDB.TerminaEm {
		return true
	} else {
		return false
	}
}

func updateIntegranteHandler(integrante mdl.Integrante, jurisdicaoDB mdl.Integrante) {
	sqlStatement := "UPDATE integrantes SET " +
		"incia_em=$1, termina_em=$2 WHERE id=$3"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	updtForm.Exec(integrante.IniciaEm, integrante.TerminaEm, integrante.Id)
	log.Println("Statement: " + sqlStatement)
}

func LoadIntegrantesByEntidadeIdCicloId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Integrantes")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	log.Println("entidadeId: " + entidadeId)
	log.Println("cicloId: " + cicloId)
	integrantes := ListIntegrantesByEntidadeIdByCicloId(entidadeId, cicloId)
	jsonIntegrantes, _ := json.Marshal(integrantes)
	w.Write([]byte(jsonIntegrantes))
	log.Println("JSON Integrantes")
}

func LoadSupervisorByEntidadeIdCicloId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Supervisor")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	log.Println("entidadeId: " + entidadeId)
	log.Println("cicloId: " + cicloId)
	sql := "SELECT a.supervisor_id, b.name " +
		"FROM ciclos_entidades a " +
		"LEFT JOIN users b ON a.supervisor_id = b.id " +
		"WHERE a.entidade_id = $1 AND a.ciclo_id = $2 "
	log.Println(sql)
	rows, _ := Db.Query(sql, entidadeId, cicloId)
	var supervisor mdl.User
	if rows.Next() {
		rows.Scan(
			&supervisor.Id,
			&supervisor.Name)
		log.Println(supervisor)
	}
	jsonSupervisor, _ := json.Marshal(supervisor)
	w.Write([]byte(jsonSupervisor))
	log.Println("JSON Supervisor")
}
