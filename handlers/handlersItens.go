package handlers

import (
	"log"
	"strconv"
	mdl "virtus/models"
)

// AJAX
func ListItensHandler(elementoId string) []mdl.Item {
	log.Println("List Itens By Elemento Id")
	sql := "SELECT " +
		" a.id, " +
		" a.elemento_id, " +
		" a.titulo," +
		" coalesce(a.descricao,''), " +
		" coalesce(a.avaliacao,''), " +
		" a.author_id, " +
		" coalesce(b.name,'') as author_name, " +
		" coalesce(to_char(a.data_criacao,'DD/MM/YYYY')) as data_criacao," +
		" a.status_id, " +
		" coalesce(c.name,'') as status_name " +
		" FROM itens a LEFT JOIN users b ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" WHERE a.elemento_id = $1 " +
		" ORDER BY a.titulo ASC"
	log.Println(sql)
	rows, _ := Db.Query(sql, elementoId)
	var itens []mdl.Item
	var item mdl.Item
	var i = 1
	for rows.Next() {
		rows.Scan(&item.Id, &item.ElementoId, &item.Nome, &item.Descricao, &item.Avaliacao, &item.AuthorId, &item.AuthorName, &item.CDataCriacao, &item.StatusId, &item.CStatus)
		item.Order = i
		i++
		itens = append(itens, item)
		log.Println(item)
	}
	return itens
}

func DeleteItensByElementoHandler(elementoId string) {
	sqlStatement := "DELETE FROM Itens WHERE elemento_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(elementoId)
	log.Println("DELETE Itens in Order Id: " + elementoId)
}

func DeleteItensHandler(diffDB []mdl.Item) {
	sqlStatement := "DELETE FROM itens WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Item Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func UpdateItensHandler(itensPage []mdl.Item, itensDB []mdl.Item) {
	for i := range itensPage {
		id := itensPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range itensDB {
			log.Println("itensDB[j].Id: " + strconv.FormatInt(itensDB[j].Id, 10))
			if strconv.FormatInt(itensDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChanged(itensPage[i], itensDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updateItemHandler(itensPage[i], itensDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChanged(itemPage mdl.Item, itemDB mdl.Item) bool {
	log.Println("itemPage.Nome: " + itemPage.Nome)
	log.Println("itemDB.Nome: " + itemDB.Nome)
	if itemPage.Nome != itemDB.Nome {
		return true
	} else if itemPage.Descricao != itemDB.Descricao {
		return true
	} else if itemPage.Avaliacao != itemDB.Avaliacao {
		return true
	} else {
		return false
	}
}

func updateItemHandler(i mdl.Item, itemDB mdl.Item) {
	sqlStatement := "UPDATE itens SET " +
		"nome=$1, descricao=$2, avaliacao=$3 WHERE id=$4"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	log.Println(i.Nome)
	log.Println(i.Descricao)
	log.Println(i.Avaliacao)
	log.Println(i.Id)
	_, err := updtForm.Exec(i.Nome, i.Descricao, i.Avaliacao, i.Id)
	if err != nil {

		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}
