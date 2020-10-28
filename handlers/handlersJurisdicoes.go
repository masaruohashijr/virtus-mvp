package handlers

import (
	"log"
	"strconv"
	//	"time"
	mdl "virtus/models"
)

// AJAX
func ListJurisdicoesByEscritorioId(escritorioId string) []mdl.Jurisdicao {
	log.Println("List Jurisdições By Escritório Id")
	log.Println("escritorioId: " + escritorioId)
	sql := "SELECT " +
		"a.id, " +
		"a.escritorio_id, " +
		"a.entidade_id, " +
		"coalesce(d.nome,'') as entidade_nome, " +
		"coalesce(to_char(a.inicia_em,'DD/MM/YYYY')) as inicia_em, " +
		"coalesce(to_char(a.termina_em,'DD/MM/YYYY')) as termina_em, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM jurisdicoes a " +
		"LEFT JOIN entidades d ON a.entidade_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"WHERE a.escritorio_id = $1 ORDER BY d.nome ASC "
	log.Println(sql)
	rows, _ := Db.Query(sql, escritorioId)
	var jurisdicoes []mdl.Jurisdicao
	var jurisdicao mdl.Jurisdicao
	var i = 1
	for rows.Next() {
		rows.Scan(
			&jurisdicao.Id,
			&jurisdicao.EscritorioId,
			&jurisdicao.EntidadeId,
			&jurisdicao.EntidadeNome,
			&jurisdicao.IniciaEm,
			&jurisdicao.TerminaEm,
			&jurisdicao.AuthorId,
			&jurisdicao.AuthorName,
			&jurisdicao.CriadoEm,
			&jurisdicao.StatusId,
			&jurisdicao.CStatus)
		jurisdicao.Order = i
		i++
		jurisdicoes = append(jurisdicoes, jurisdicao)
		log.Println(jurisdicao)
	}
	return jurisdicoes
}

func UpdateJurisdicoesHandler(jurisdicoesPage []mdl.Jurisdicao, jurisdicoesDB []mdl.Jurisdicao) {
	for i := range jurisdicoesPage {
		id := jurisdicoesPage[i].Id
		log.Println("Jurisdicao Page id: " + strconv.FormatInt(id, 10))
		for j := range jurisdicoesDB {
			log.Println("jurisdicoesDB[j].Id: " + strconv.FormatInt(jurisdicoesDB[j].Id, 10))
			if strconv.FormatInt(jurisdicoesDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				fieldsChanged := hasSomeFieldChangedJurisdicao(jurisdicoesPage[i], jurisdicoesDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateJurisdicaoHandler(jurisdicoesPage[i], jurisdicoesDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedJurisdicao(jurisdicaoPage mdl.Jurisdicao, jurisdicaoDB mdl.Jurisdicao) bool {
	if jurisdicaoPage.EntidadeId != jurisdicaoDB.EntidadeId {
		return true
	} else if jurisdicaoPage.IniciaEm != jurisdicaoDB.IniciaEm {
		return true
	} else if jurisdicaoPage.TerminaEm != jurisdicaoDB.TerminaEm {
		return true
	} else {
		return false
	}
}

func updateJurisdicaoHandler(jurisdicao mdl.Jurisdicao, jurisdicaoDB mdl.Jurisdicao) {
	sqlStatement := "UPDATE jurisdicoes SET " +
		"entidade_id=$1 WHERE id=$2"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(jurisdicao.EntidadeId, jurisdicao.Id)
	if err != nil {
		log.Println(err)
	}
	log.Println("De " + jurisdicao.IniciaEm + " a " + jurisdicao.TerminaEm)
	if jurisdicao.IniciaEm != "" {
		sqlStatement = "UPDATE jurisdicoes SET inicia_em = to_date('" +
			jurisdicao.IniciaEm + "','YYYY-MM-DD') " +
			"WHERE id = " + strconv.FormatInt(jurisdicao.Id, 10)
		_, err = Db.Exec(sqlStatement)
		if err != nil {
			log.Println(err)
		}
	}
	if jurisdicao.TerminaEm != "" {
		sqlStatement = "UPDATE jurisdicoes SET termina_em = to_date('" +
			jurisdicao.TerminaEm + "','YYYY-MM-DD') " +
			"WHERE id = " + strconv.FormatInt(jurisdicao.Id, 10)
		_, err = Db.Exec(sqlStatement)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteJurisdicoesByEscritorioId(escritorioId string) {
	sqlStatement := "DELETE FROM jurisdicoes WHERE escritorio_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(escritorioId)
	log.Println("DELETE jurisdicoes do Escritorio Id: " + escritorioId)
}

func DeleteJurisdicoesHandler(diffDB []mdl.Jurisdicao) {
	sqlStatement := "DELETE FROM jurisdicoes WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Jurisdicao Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsJurisdicao(jurisdicoes []mdl.Jurisdicao, jurisdicaoCompared mdl.Jurisdicao) bool {
	for n := range jurisdicoes {
		if jurisdicoes[n].Id == jurisdicaoCompared.Id {
			return true
		}
	}
	return false
}

func removeJurisdicao(jurisdicoes []mdl.Jurisdicao, jurisdicaoToBeRemoved mdl.Jurisdicao) []mdl.Jurisdicao {
	var newJurisdicoes []mdl.Jurisdicao
	for i := range jurisdicoes {
		if jurisdicoes[i].Id != jurisdicaoToBeRemoved.Id {
			newJurisdicoes = append(newJurisdicoes, jurisdicoes[i])
		}
	}
	return newJurisdicoes
}
