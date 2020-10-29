package handlers

import (
	//	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func UpdateEquipeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Equipe")
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

		// Membros
		var membrosDB = ListMembrosByEscritorioId(escritorioId)
		var membrosPage []mdl.Membro
		var membroPage mdl.Membro
		for key, value := range r.Form {
			if strings.HasPrefix(key, "membro") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				membroPage.Id, _ = strconv.ParseInt(id, 10, 64)
				membroPage.EscritorioId, _ = strconv.ParseInt(escritorioId, 10, 64)
				usuarioId := strings.Split(array[3], ":")[1]
				log.Println("usuarioId -------- " + usuarioId)
				membroPage.UsuarioId, _ = strconv.ParseInt(usuarioId, 10, 64)
				usuarioNome := strings.Split(array[4], ":")[1]
				log.Println("usuarioNome -------- " + usuarioNome)
				membroPage.UsuarioNome = usuarioNome
				iniciaEm := strings.Split(array[5], ":")[1]
				log.Println("iniciaEm -------- " + iniciaEm)
				membroPage.IniciaEm = iniciaEm
				terminaEm := strings.Split(array[6], ":")[1]
				log.Println("terminaEm -------- " + terminaEm)
				membroPage.TerminaEm = terminaEm
				autorId := strings.Split(array[7], ":")[1]
				log.Println("autorId -------- " + autorId)
				membroPage.AuthorId, _ = strconv.ParseInt(autorId, 10, 64)
				autorNome := strings.Split(array[8], ":")[1]
				log.Println("autorNome -------- " + autorNome)
				membroPage.AuthorName = autorNome
				criadoEm := strings.Split(array[9], ":")[1]
				log.Println("criadoEm -------- " + criadoEm)
				membroPage.CriadoEm = criadoEm
				idVersaoOrigem := strings.Split(array[10], ":")[1]
				log.Println("idVersaoOrigem -------- " + idVersaoOrigem)
				membroPage.IdVersaoOrigem, _ = strconv.ParseInt(idVersaoOrigem, 10, 64)
				statusId := strings.Split(array[11], ":")[1]
				log.Println("statusId -------- " + statusId)
				membroPage.StatusId, _ = strconv.ParseInt(statusId, 10, 64)
				cStatus := strings.Split(array[12], ":")[1]
				log.Println("cStatus -------- " + cStatus)
				membroPage.CStatus = cStatus
				membrosPage = append(membrosPage, membroPage)
			}
		}
		if len(membrosPage) < len(membrosDB) {
			log.Println("Quantidade de Entidades do Escritório da Página: " + strconv.Itoa(len(membrosPage)))
			if len(membrosPage) == 0 {
				DeleteMembrosByEscritorioId(escritorioId) //DONE
			} else {
				var diffDB []mdl.Membro = membrosDB
				for n := range membrosPage {
					if containsMembro(diffDB, membrosPage[n]) {
						diffDB = removeMembro(diffDB, membrosPage[n])
					}
				}
				DeleteMembrosHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Membro = membrosPage
			for n := range membrosDB {
				if containsMembro(diffPage, membrosDB[n]) {
					diffPage = removeMembro(diffPage, membrosDB[n])
				}
			}
			var membro mdl.Membro
			membroId := 0
			statusComponenteId := GetStartStatus("plano")
			for i := range diffPage {
				membro = diffPage[i]
				log.Println("Escritorio Id: " + escritorioId)
				sqlStatement := "INSERT INTO membros ( " +
					" escritorio_id, " +
					" usuario_id, " +
					" inicia_em, " +
					" termina_em, " +
					" author_id, " +
					" criado_em, " +
					" status_id " +
					" ) " +
					" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
				log.Println(sqlStatement)
				Db.QueryRow(
					sqlStatement,
					escritorioId,
					membro.UsuarioId,
					membro.IniciaEm,
					membro.TerminaEm,
					currentUser.Id,
					time.Now(),
					statusComponenteId).Scan(&membroId)
			}
		}
		UpdateMembrosHandler(membrosPage, membrosDB)

		http.Redirect(w, r, route.EscritoriosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListMembrosByEscritorioId(escritorioId string) []mdl.Membro {
	log.Println("List Membros By Escritório Id")
	log.Println("escritorioId: " + escritorioId)
	sql := "SELECT " +
		"a.id, " +
		"a.escritorio_id, " +
		"coalesce(d.nome,'') as escritorio_nome, " +
		"a.pilar_id, " +
		"a.peso_padrao || ' %', " +
		"a.tipo_media, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM jurisdicoes a " +
		"LEFT JOIN pilares d ON a.pilar_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"WHERE a.escritorio_id = $1 ORDER BY d.nome ASC "
	log.Println(sql)
	rows, _ := Db.Query(sql, escritorioId)
	var membros []mdl.Membro
	var membro mdl.Membro
	var i = 1
	for rows.Next() {
		rows.Scan(
			&membro.Id,
			&membro.EscritorioId,
			&membro.UsuarioId,
			&membro.UsuarioNome,
			&membro.IniciaEm,
			&membro.TerminaEm,
			&membro.AuthorId,
			&membro.AuthorName,
			&membro.CriadoEm,
			&membro.StatusId,
			&membro.CStatus)
		membro.Order = i
		i++
		membros = append(membros, membro)
		log.Println(membro)
	}
	return membros
}

func DeleteMembrosByEscritorioId(escritorioId string) {
	sqlStatement := "DELETE FROM membros WHERE escritorio_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(escritorioId)
	log.Println("DELETE membros in Escritorio Id: " + escritorioId)
}

func DeleteMembrosHandler(diffDB []mdl.Membro) {
	sqlStatement := "DELETE FROM membros WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Membro Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsMembro(membros []mdl.Membro, membroCompared mdl.Membro) bool {
	for n := range membros {
		if membros[n].Id == membroCompared.Id {
			return true
		}
	}
	return false
}

func removeMembro(membros []mdl.Membro, membroToBeRemoved mdl.Membro) []mdl.Membro {
	var newMembros []mdl.Membro
	for i := range membros {
		if membros[i].Id != membroToBeRemoved.Id {
			newMembros = append(newMembros, membros[i])
		}
	}
	return newMembros
}

func UpdateMembrosHandler(membrosPage []mdl.Membro, membrosDB []mdl.Membro) {
	for i := range membrosPage {
		id := membrosPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range membrosDB {
			log.Println("membrosDB[j].Id: " + strconv.FormatInt(membrosDB[j].Id, 10))
			if strconv.FormatInt(membrosDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				fieldsChanged := hasSomeFieldChangedMembro(membrosPage[i], membrosDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateMembroHandler(membrosPage[i], membrosDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedMembro(jurisdicaoPage mdl.Membro, jurisdicaoDB mdl.Membro) bool {
	return false
}

func updateMembroHandler(ce mdl.Membro, jurisdicaoDB mdl.Membro) {
	sqlStatement := "UPDATE jurisdicoes SET " +
		"tipo_media=$1, peso_padrao=$2 WHERE id=$3"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	updtForm.Exec(ce.Id)
	log.Println("Statement: " + sqlStatement)
}
