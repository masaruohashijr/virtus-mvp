package handlers

import (
	"log"
	"strconv"
	//	"time"
	mdl "virtus/models"
)

// AJAX
func ListComponentesByPilarId(pilarId string) []mdl.ComponentePilar {
	log.Println("List Componentes By Pilar Id")
	sql := "SELECT " +
		" a.id, " +
		" a.pilar_id, " +
		" a.componente_id," +
		" coalesce(c.nome,'') as componente_nome," +
		" a.tipo_media," +
		" b.peso_padrao," +
		" coalesce(a.sonda,'') as sonda," +
		" a.author_id, " +
		" coalesce(u.name,'') as author_name, " +
		" coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em," +
		" a.status_id, " +
		" coalesce(s.name,'') as status_name " +
		" FROM componentes_pilares a " +
		" LEFT JOIN componentes c ON a.componente_id = c.id " +
		" LEFT JOIN users u ON a.author_id = u.id " +
		" LEFT JOIN status s ON a.status_id = s.id " +
		" LEFT JOIN (select componente_id, round(avg(peso_padrao),2) as peso_padrao from elementos_componentes group by componente_id) b ON a.componente_id = b.componente_id " +
		" WHERE a.pilar_id = $1 " +
		" ORDER BY c.nome ASC"
	log.Println(sql)
	rows, _ := Db.Query(sql, pilarId)
	var componentesPilar []mdl.ComponentePilar
	var componentePilar mdl.ComponentePilar
	var i = 1
	for rows.Next() {
		rows.Scan(
			&componentePilar.Id,
			&componentePilar.PilarId,
			&componentePilar.ComponenteId,
			&componentePilar.ComponenteNome,
			&componentePilar.TipoMediaId,
			&componentePilar.PesoPadrao,
			&componentePilar.Sonda,
			&componentePilar.AuthorId,
			&componentePilar.AuthorName,
			&componentePilar.CriadoEm,
			&componentePilar.StatusId,
			&componentePilar.CStatus)
		componentePilar.Order = i
		i++
		switch componentePilar.TipoMediaId {
		case 1:
			componentePilar.TipoMedia = "Aritmética"
		case 2:
			componentePilar.TipoMedia = "Geométrica"
		case 3:
			componentePilar.TipoMedia = "Harmônica"
		}
		componentesPilar = append(componentesPilar, componentePilar)
		log.Println(componentePilar)
	}
	return componentesPilar
}

func UpdateComponentesPilarHandler(componentesPilarPage []mdl.ComponentePilar, componentesPilarDB []mdl.ComponentePilar) {
	for i := range componentesPilarPage {
		id := componentesPilarPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range componentesPilarDB {
			log.Println("componentesPilarDB[j].Id: " + strconv.FormatInt(componentesPilarDB[j].Id, 10))
			if strconv.FormatInt(componentesPilarDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChangedPilar(componentesPilarPage[i], componentesPilarDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateComponentePilarHandler(componentesPilarPage[i], componentesPilarDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedPilar(componentePilarPage mdl.ComponentePilar, componentePilarDB mdl.ComponentePilar) bool {
	log.Println("Id componente pilar: " + strconv.FormatInt(componentePilarPage.Id, 10))
	log.Println(componentePilarPage)
	if componentePilarPage.TipoMediaId != componentePilarDB.TipoMediaId {
		return true
	} else if componentePilarPage.PesoPadrao != componentePilarDB.PesoPadrao {
		return true
	} else if componentePilarPage.Sonda != componentePilarDB.Sonda {
		return true
	} else {
		return false
	}
}

func updateComponentePilarHandler(componentePilarPage mdl.ComponentePilar, componentePilarDB mdl.ComponentePilar) {
	sqlStatement := "UPDATE componentes_pilares SET " +
		"tipo_media=$1, peso_padrao=$2, sonda=$3 WHERE id=$4"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(componentePilarPage.TipoMediaId, componentePilarPage.PesoPadrao, componentePilarPage.Sonda, componentePilarPage.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteComponentesPilarByPilarId(pilarId string) {
	sqlStatement := "DELETE FROM componentes_pilares WHERE pilar_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(pilarId)
	log.Println("DELETE componentes_pilares in Order Id: " + pilarId)
}

func DeleteComponentesPilarHandler(diffDB []mdl.ComponentePilar) {
	sqlStatement := "DELETE FROM componentes_pilares WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE:  Pilar Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsComponentePilar(componentesPilar []mdl.ComponentePilar, componentePilarCompared mdl.ComponentePilar) bool {
	for n := range componentesPilar {
		if componentesPilar[n].Id == componentePilarCompared.Id {
			return true
		}
	}
	return false
}

func removeComponentePilar(componentesPilar []mdl.ComponentePilar, componentePilarToBeRemoved mdl.ComponentePilar) []mdl.ComponentePilar {
	var newComponentesPilar []mdl.ComponentePilar
	for i := range componentesPilar {
		if componentesPilar[i].Id != componentePilarToBeRemoved.Id {
			newComponentesPilar = append(newComponentesPilar, componentesPilar[i])
		}
	}
	return newComponentesPilar
}
