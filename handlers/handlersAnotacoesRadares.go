package handlers

import (
	"log"
	"strconv"
	"strings"
	"time"
	mdl "virtus/models"
)

// AJAX
func ListAnotacoesRadarByRadarId(radarId string) []mdl.AnotacaoRadar {
	log.Println("List Anotacoes Radares By Radar Id")
	log.Println("radarId: " + radarId)
	sql := " SELECT a.id, " +
		" a.radar_id, " +
		" d.entidade_id, " +
		" a.anotacao_id, " +
		" coalesce(a.observacoes,''), " +
		" coalesce(a.registro_ata,''), " +
		" a.author_id, " +
		" coalesce(b.name,'') as author_name, " +
		" coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em, " +
		" a.status_id, " +
		" coalesce(c.name,'') as status_name, " +
		" a.ultimo_atualizador_id, " +
		" coalesce(e.name,'') as ultimo_atualizador_name, " +
		" coalesce(to_char(a.ultima_atualizacao,'DD/MM/YYYY')) " +
		" FROM anotacoes_radares a " +
		" LEFT JOIN anotacoes d ON a.anotacao_id = d.id " +
		" LEFT JOIN users b ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" LEFT JOIN users e ON a.ultimo_atualizador_id = e.id " +
		" WHERE a.radar_id = $1 "
	log.Println(sql)
	rows, _ := Db.Query(sql, radarId)
	defer rows.Close()
	var anotacoesRadar []mdl.AnotacaoRadar
	var anotacaoRadar mdl.AnotacaoRadar
	for rows.Next() {
		rows.Scan(
			&anotacaoRadar.Id,
			&anotacaoRadar.RadarId,
			&anotacaoRadar.EntidadeId,
			&anotacaoRadar.AnotacaoId,
			&anotacaoRadar.Observacoes,
			&anotacaoRadar.RegistroAta,
			&anotacaoRadar.AuthorId,
			&anotacaoRadar.AuthorName,
			&anotacaoRadar.CriadoEm,
			&anotacaoRadar.StatusId,
			&anotacaoRadar.CStatus,
			&anotacaoRadar.UltimoAtualizadorId,
			&anotacaoRadar.UltimoAtualizadorNome,
			&anotacaoRadar.UltimaAtualizacao)
		anotacoesRadar = append(anotacoesRadar, anotacaoRadar)
		log.Println(anotacaoRadar)
	}
	return anotacoesRadar
}

func UpdateAnotacoesRadarHandler(anotacoesRadarPage []mdl.AnotacaoRadar, anotacoesRadarDB []mdl.AnotacaoRadar, currentUserId int64) {
	for i := range anotacoesRadarPage {
		id := anotacoesRadarPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range anotacoesRadarDB {
			log.Println("anotacoesRadarDB[j].Id: " + strconv.FormatInt(anotacoesRadarDB[j].Id, 10))
			if strconv.FormatInt(anotacoesRadarDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChangedAnotacaoRadar(anotacoesRadarPage[i], anotacoesRadarDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateAnotacaoRadarHandler(anotacoesRadarPage[i], anotacoesRadarDB[j], currentUserId)
				}
				anotacoesRadarDB = removeAnotacaoRadar(anotacoesRadarDB, anotacoesRadarPage[i])
				break
			}
		}
	}
	DeleteAnotacoesRadarHandler(anotacoesRadarDB) // CORREÇÃO
}

func hasSomeFieldChangedAnotacaoRadar(anotacaoRadarPage mdl.AnotacaoRadar, anotacaoRadarDB mdl.AnotacaoRadar) bool {
	if anotacaoRadarPage.Observacoes != anotacaoRadarDB.Observacoes {
		return true
	} else if anotacaoRadarPage.RegistroAta != anotacaoRadarDB.RegistroAta {
		return true
	} else {
		return false
	}
}

func updateAnotacaoRadarHandler(anotacaoRadar mdl.AnotacaoRadar, anotacaoRadarDB mdl.AnotacaoRadar, currentUserId int64) {
	sqlStatement := "UPDATE anotacoes_radares " +
		" SET radar_id=$1, anotacao_id=$2, observacoes=$3, registro_ata=$4, " +
		" ultimo_atualizador_id=$5, ultima_atualizacao=$6 " +
		" WHERE id = $7 "
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(anotacaoRadar.RadarId,
		anotacaoRadar.AnotacaoId,
		anotacaoRadar.Observacoes,
		anotacaoRadar.RegistroAta,
		currentUserId,
		time.Now(),
		anotacaoRadar.Id)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteAnotacoesRadarByRadarId(radarId string) {
	sqlStatement := "DELETE FROM anotacoes_radares WHERE radar_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	deleteForm.Exec(radarId)
	log.Println("DELETE anotacoes_radares in Order Id: " + radarId)
}

func DeleteAnotacoesRadarHandler(diffDB []mdl.AnotacaoRadar) string {
	sqlStatement := "DELETE FROM anotacoes_radares WHERE id=$1"
	deleteForm, _ := Db.Prepare(sqlStatement)
	for n := range diffDB {
		errMsg := ""
		_, err := deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Anotacao Radar Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			errMsg = "Anotacao está associada a um registro e não pôde ser removida."
			return errMsg
		}
	}
	return ""
}

func containsAnotacaoRadar(anotacoesRadar []mdl.AnotacaoRadar, anotacaoRadarCompared mdl.AnotacaoRadar) bool {
	for n := range anotacoesRadar {
		if anotacoesRadar[n].Id == anotacaoRadarCompared.Id {
			return true
		}
	}
	return false
}

func removeAnotacaoRadar(anotacoesRadar []mdl.AnotacaoRadar, anotacaoRadarToBeRemoved mdl.AnotacaoRadar) []mdl.AnotacaoRadar {
	var newAnotacoesRadar []mdl.AnotacaoRadar
	for i := range anotacoesRadar {
		if anotacoesRadar[i].Id != anotacaoRadarToBeRemoved.Id {
			newAnotacoesRadar = append(newAnotacoesRadar, anotacoesRadar[i])
		}
	}
	return newAnotacoesRadar
}
