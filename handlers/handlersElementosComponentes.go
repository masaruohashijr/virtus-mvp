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
		"a.tipo_nota_id, " +
		"e.nome as tipo_nota_nome, " +
		"a.peso_padrao, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM elementos_componentes a " +
		"LEFT JOIN elementos d ON a.elemento_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"LEFT JOIN tipos_notas e ON a.tipo_nota_id = e.id " +
		"WHERE a.componente_id = $1 ORDER BY elemento_nome"
	log.Println(sql)
	rows, _ := Db.Query(sql, componenteId)
	defer rows.Close()
	var elementosComponente []mdl.ElementoComponente
	var elementoComponente mdl.ElementoComponente
	var i = 1
	for rows.Next() {
		rows.Scan(
			&elementoComponente.Id,
			&elementoComponente.ComponenteId,
			&elementoComponente.ElementoNome,
			&elementoComponente.ElementoId,
			&elementoComponente.TipoNotaId,
			&elementoComponente.TipoNotaNome,
			&elementoComponente.PesoPadrao,
			&elementoComponente.AuthorId,
			&elementoComponente.AuthorName,
			&elementoComponente.CriadoEm,
			&elementoComponente.StatusId,
			&elementoComponente.CStatus)
		elementoComponente.Order = i
		i++
		elementosComponente = append(elementosComponente, elementoComponente)
		//log.Println(elementoComponente)
	}
	return elementosComponente
}

func UpdateElementosComponenteHandler(elementosComponentePage []mdl.ElementoComponente, elementosComponenteDB []mdl.ElementoComponente) {
	log.Println("UpdateElementosComponenteHandler")
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
				elementosComponenteDB = removeElementoComponente(elementosComponenteDB, elementosComponentePage[i])
				break
			}
		}
	}
	DeleteElementosComponenteHandler(elementosComponenteDB)
}

func hasSomeFieldChangedElementoComponente(elementoComponentePage mdl.ElementoComponente, elementoComponenteDB mdl.ElementoComponente) bool {
	log.Println("hasSomeFieldChangedElementoComponente")
	log.Println(elementoComponentePage.PesoPadrao)
	log.Println(elementoComponenteDB.PesoPadrao)
	if elementoComponentePage.PesoPadrao != elementoComponenteDB.PesoPadrao {
		return true
	} else {
		return false
	}
}

func updateElementoComponenteHandler(elementoComponente mdl.ElementoComponente, elementoComponenteDB mdl.ElementoComponente) {
	log.Println("updateElementoComponenteHandler")
	sqlStatement := "UPDATE elementos_componentes SET " +
		"peso_padrao=$1 WHERE id=$2"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(elementoComponente.PesoPadrao, elementoComponente.Id)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
	sqlStatement = "UPDATE componentes_pilares a " +
		" SET peso_padrao = " +
		" (SELECT round(avg(b.peso_padrao),2) " +
		" FROM elementos_componentes b " +
		" WHERE b.componente_id = a.componente_id AND " +
		" b.pilar_id = a.pilar_id AND " +
		" b.ciclo_id = a.ciclo_id AND " +
		" b.entidade_id = a.entidade_id " +
		" GROUP BY b.entidade_id, b.ciclo_id, b.pilar_id, b.componente_id) " +
		" WHERE a.id=$1 "
	log.Println(sqlStatement)
	updtForm, _ = Db.Prepare(sqlStatement)
	_, err = updtForm.Exec(elementoComponente.ComponenteId)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteElementosComponenteByComponenteId(componenteId string) {
	sqlStatement := "DELETE FROM elementos_componentes WHERE componente_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	deleteForm.Exec(componenteId)
	log.Println("DELETE elementos_componentes in Order Id: " + componenteId)
}

func DeleteElementosComponenteHandler(diffDB []mdl.ElementoComponente) {
	sqlStatement := "DELETE FROM elementos_componentes WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
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
