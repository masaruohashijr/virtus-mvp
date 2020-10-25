package handlers

import (
	"log"
	"strconv"
	//	"time"
	mdl "virtus/models"
)

// AJAX
func ListPilaresByCicloId(cicloId string) []mdl.PilarCiclo {
	log.Println("List Pilares Ciclos By Ciclo Id")
	log.Println("cicloId: " + cicloId)
	sql := "SELECT " +
		"a.id, " +
		"a.ciclo_id, " +
		"coalesce(d.nome,'') as pilar_nome, " +
		"a.pilar_id, " +
		"a.peso_padrao, " +
		"a.tipo_media, " +
		"a.author_id, " +
		"coalesce(b.name,'') as author_name, " +
		"coalesce(to_char(a.criado_em,'YYYY-MM-DD')) as criado_em, " +
		"a.status_id, " +
		"coalesce(c.name,'') as status_name " +
		"FROM pilares_ciclos a " +
		"LEFT JOIN pilares d ON a.pilar_id = d.id " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"LEFT JOIN status c ON a.status_id = c.id " +
		"WHERE a.ciclo_id = $1 "
	log.Println(sql)
	rows, _ := Db.Query(sql, cicloId)
	var pilaresCiclo []mdl.PilarCiclo
	var pilarCiclo mdl.PilarCiclo
	var i = 1
	for rows.Next() {
		rows.Scan(
			&pilarCiclo.Id,
			&pilarCiclo.CicloId,
			&pilarCiclo.PilarNome,
			&pilarCiclo.PilarId,
			&pilarCiclo.PesoPadrao,
			&pilarCiclo.TipoMediaId,
			&pilarCiclo.AuthorId,
			&pilarCiclo.AuthorName,
			&pilarCiclo.CriadoEm,
			&pilarCiclo.StatusId,
			&pilarCiclo.CStatus)
		pilarCiclo.Order = i
		i++
		switch pilarCiclo.TipoMediaId {
		case 1:
			pilarCiclo.TipoMedia = "Aritmética"
		case 2:
			pilarCiclo.TipoMedia = "Geométrica"
		case 3:
			pilarCiclo.TipoMedia = "Harmônica"
		}
		pilaresCiclo = append(pilaresCiclo, pilarCiclo)
		log.Print("log linha 60: ")
		log.Println(pilarCiclo)
	}
	return pilaresCiclo
}

func UpdatePilaresCicloHandler(pilaresCicloPage []mdl.PilarCiclo, pilaresCicloDB []mdl.PilarCiclo) {
	for i := range pilaresCicloPage {
		id := pilaresCicloPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range pilaresCicloDB {
			log.Println("pilaresCicloDB[j].Id: " + strconv.FormatInt(pilaresCicloDB[j].Id, 10))
			if strconv.FormatInt(pilaresCicloDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChangedPilarCiclo(pilaresCicloPage[i], pilaresCicloDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updatePilarCicloHandler(pilaresCicloPage[i], pilaresCicloDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChangedPilarCiclo(pilarCicloPage mdl.PilarCiclo, pilarCicloDB mdl.PilarCiclo) bool {
	if pilarCicloPage.TipoMediaId != pilarCicloDB.TipoMediaId {
		return true
	} else if pilarCicloPage.PesoPadrao != pilarCicloDB.PesoPadrao {
		return true
	} else {
		return false
	}
}

func updatePilarCicloHandler(ce mdl.PilarCiclo, pilarCicloDB mdl.PilarCiclo) {
	sqlStatement := "UPDATE pilares_ciclos SET " +
		"tipo_media=$1, peso_padrao=$2 WHERE id=$3"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(ce.TipoMediaId, ce.PesoPadrao, ce.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeletePilaresCicloByCicloId(cicloId string) {
	sqlStatement := "DELETE FROM pilares_ciclos WHERE ciclo_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(cicloId)
	log.Println("DELETE pilares_ciclos in Order Id: " + cicloId)
}

func DeletePilaresCicloHandler(diffDB []mdl.PilarCiclo) {
	sqlStatement := "DELETE FROM pilares_ciclos WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Pilar Ciclo Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsPilarCiclo(pilaresCiclo []mdl.PilarCiclo, pilarCicloCompared mdl.PilarCiclo) bool {
	for n := range pilaresCiclo {
		if pilaresCiclo[n].Id == pilarCicloCompared.Id {
			return true
		}
	}
	return false
}

func removePilarCiclo(pilaresCiclo []mdl.PilarCiclo, pilarCicloToBeRemoved mdl.PilarCiclo) []mdl.PilarCiclo {
	var newPilaresCiclo []mdl.PilarCiclo
	for i := range pilaresCiclo {
		if pilaresCiclo[i].Id != pilarCicloToBeRemoved.Id {
			newPilaresCiclo = append(newPilaresCiclo, pilaresCiclo[i])
		}
	}
	return newPilaresCiclo
}
