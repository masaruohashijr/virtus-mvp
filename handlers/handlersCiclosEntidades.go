package handlers

import (
	"log"
	"strconv"
	mdl "virtus/models"
)

// AJAX
func ListCiclosEntidadeByEntidadeId(entidadeId string) []mdl.CicloEntidade {
	log.Println("List Ciclos Entidades By Entidade Id")
	log.Println("entidadeId: " + entidadeId)
	sql := "SELECT " +
		"a.id, " +
		"a.entidade_id, " +
		"d.nome, " +
		"a.ciclo_id, " +
		"a.tipo_media, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.inicia_em,'DD/MM/YYYY')) as inicia_em, " +
		"coalesce(to_char(a.termina_em,'DD/MM/YYYY')) as termina_em, " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM ciclos_entidades a " +
		"LEFT JOIN ciclos d ON a.ciclo_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"WHERE a.entidade_id = $1 "
	log.Println(sql)
	rows, _ := Db.Query(sql, entidadeId)
	var ciclosEntidade []mdl.CicloEntidade
	var cicloEntidade mdl.CicloEntidade
	var i = 1
	for rows.Next() {
		rows.Scan(
			&cicloEntidade.Id,
			&cicloEntidade.EntidadeId,
			&cicloEntidade.Nome,
			&cicloEntidade.CicloId,
			&cicloEntidade.TipoMediaId,
			&cicloEntidade.AuthorId,
			&cicloEntidade.AuthorName,
			&cicloEntidade.IniciaEm,
			&cicloEntidade.TerminaEm,
			&cicloEntidade.CriadoEm,
			&cicloEntidade.StatusId,
			&cicloEntidade.CStatus)
		cicloEntidade.Order = i
		i++
		switch cicloEntidade.TipoMediaId {
		case 1:
			cicloEntidade.TipoMedia = "Aritmética"
		case 2:
			cicloEntidade.TipoMedia = "Geométrica"
		case 3:
			cicloEntidade.TipoMedia = "Harmônica"
		}
		ciclosEntidade = append(ciclosEntidade, cicloEntidade)
		log.Println(cicloEntidade)
	}
	return ciclosEntidade
}

func UpdateCiclosEntidadeHandler(ciclosEntidadePage []mdl.CicloEntidade, ciclosEntidadeDB []mdl.CicloEntidade) {
	for i := range ciclosEntidadePage {
		id := ciclosEntidadePage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range ciclosEntidadeDB {
			log.Println("ciclosEntidadeDB[j].Id: " + strconv.FormatInt(ciclosEntidadeDB[j].Id, 10))
			if strconv.FormatInt(ciclosEntidadeDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChangedCicloEntidade(ciclosEntidadePage[i], ciclosEntidadeDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateCicloEntidadeHandler(ciclosEntidadePage[i], ciclosEntidadeDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedCicloEntidade(cicloEntidadePage mdl.CicloEntidade, cicloEntidadeDB mdl.CicloEntidade) bool {
	if cicloEntidadePage.TipoMediaId != cicloEntidadeDB.TipoMediaId {
		return true
	} else if cicloEntidadePage.IniciaEm != cicloEntidadeDB.IniciaEm {
		return true
	} else if cicloEntidadePage.TerminaEm != cicloEntidadeDB.TerminaEm {
		return true
	} else {
		return false
	}
}

func updateCicloEntidadeHandler(ce mdl.CicloEntidade, cicloEntidadeDB mdl.CicloEntidade) {
	sqlStatement := "UPDATE ciclos_entidades SET " +
		"tipo_media=$1, inicia_em=$2, termina_em=$3 WHERE id=$4"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(ce.TipoMediaId, ce.IniciaEm, ce.TerminaEm)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteCiclosEntidadeByEntidadeId(entidadeId string) {
	sqlStatement := "DELETE FROM ciclos_entidades WHERE entidade_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(entidadeId)
	log.Println("DELETE ciclos_entidades in Order Id: " + entidadeId)
}

func DeleteCiclosEntidadeHandler(diffDB []mdl.CicloEntidade) {
	sqlStatement := "DELETE FROM ciclos_entidades WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Ciclo Entidade Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsCicloEntidade(ciclosEntidade []mdl.CicloEntidade, cicloEntidadeCompared mdl.CicloEntidade) bool {
	for n := range ciclosEntidade {
		if ciclosEntidade[n].Id == cicloEntidadeCompared.Id {
			return true
		}
	}
	return false
}

func removeCicloEntidade(ciclosEntidade []mdl.CicloEntidade, cicloEntidadeToBeRemoved mdl.CicloEntidade) []mdl.CicloEntidade {
	var newCiclosEntidade []mdl.CicloEntidade
	for i := range ciclosEntidade {
		if ciclosEntidade[i].Id != cicloEntidadeToBeRemoved.Id {
			newCiclosEntidade = append(newCiclosEntidade, ciclosEntidade[i])
		}
	}
	return newCiclosEntidade
}
