package handlers

import (
	"log"
	"strconv"
	//	"time"
	mdl "virtus/models"
)

// AJAX
func ListElementosByComponenteId(componenteId string) []mdl.ElementoComponente {
	log.Println("List Elementos Componentes By Componente Id")
	log.Println("componenteId: " + componenteId)
	sql := "SELECT " +
		"a.id, " +
		"a.componente_id, " +
		"coalesce(d.nome,'') as elemento_nome, " +
		"a.elemento_id, " +
		"a.peso_padrao, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.criado_em,'YYYY-MM-DD')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM elementos_componentes a " +
		"LEFT JOIN elementos d ON a.elemento_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"WHERE a.componente_id = $1 "
	log.Println(sql)
	rows, _ := Db.Query(sql, componenteId)
	var elementosComponente []mdl.ElementoComponente
	var elementoComponente mdl.ElementoComponente
	var i = 1
	for rows.Next() {
		rows.Scan(
			&elementoComponente.Id,
			&elementoComponente.ComponenteId,
			&elementoComponente.ElementoNome,
			&elementoComponente.ElementoId,
			&elementoComponente.PesoPadrao,
			&elementoComponente.AuthorId,
			&elementoComponente.AuthorName,
			&elementoComponente.CriadoEm,
			&elementoComponente.StatusId,
			&elementoComponente.CStatus)
		elementoComponente.Order = i
		i++
		elementosComponente = append(elementosComponente, elementoComponente)
		log.Println(elementoComponente)
	}
	return elementosComponente
}

func UpdateElementosComponenteHandler(elementosComponentePage []mdl.ElementoComponente, elementosComponenteDB []mdl.ElementoComponente) {
	for i := range elementosComponentePage {
		id := elementosComponentePage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range elementosComponenteDB {
			if strconv.FormatInt(elementosComponenteDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				fieldsChanged := hasSomeFieldChangedElementoComponente(elementosComponentePage[i], elementosComponenteDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateElementoComponenteHandler(elementosComponentePage[i], elementosComponenteDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedElementoComponente(elementoComponentePage mdl.ElementoComponente, elementoComponenteDB mdl.ElementoComponente) bool {
	if elementoComponentePage.PesoPadrao != elementoComponenteDB.PesoPadrao {
		return true
	} else {
		return false
	}
}

func updateElementoComponenteHandler(elementoComponente mdl.ElementoComponente, elementoComponenteDB mdl.ElementoComponente) {
	sqlStatement := "UPDATE elementos_componentes SET " +
		"peso_padrao=$1 WHERE id=$2"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(elementoComponente.PesoPadrao, elementoComponente.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteElementosComponenteByComponenteId(componenteId string) {
	sqlStatement := "DELETE FROM elementos_componentes WHERE componente_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(componenteId)
	log.Println("DELETE elementos_componentes in Order Id: " + componenteId)
}

func DeleteElementosComponenteHandler(diffDB []mdl.ElementoComponente) {
	sqlStatement := "DELETE FROM elementos_componentes WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Elemento Componente Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsElementoComponente(elementosComponente []mdl.ElementoComponente, elementoComponenteCompared mdl.ElementoComponente) bool {
	for n := range elementosComponente {
		if elementosComponente[n].Id == elementoComponenteCompared.Id {
			return true
		}
	}
	return false
}

func removeElementoComponente(elementosComponente []mdl.ElementoComponente, elementoComponenteToBeRemoved mdl.ElementoComponente) []mdl.ElementoComponente {
	var newElementosComponente []mdl.ElementoComponente
	for i := range elementosComponente {
		if elementosComponente[i].Id != elementoComponenteToBeRemoved.Id {
			newElementosComponente = append(newElementosComponente, elementosComponente[i])
		}
	}
	return newElementosComponente
}
