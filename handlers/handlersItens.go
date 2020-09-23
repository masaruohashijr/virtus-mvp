package handlers

import (
	mdl "virtus/models"
	//	sec "virtus/security"
	//	"reflect"
	//	pq "github.com/lib/pq"
	//	"html/template"
	//	"fmt"
	"log"
	//	"net/http"
	//	"strconv"
	//	"strings"
	//	"time"
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
		" coalesce(to_char(a.data_criacao,'DD/MM/YYYY'))," +
		" a.status_id, " +
		" coalesce(c.name,'') as status_name " +
		" FROM itens a LEFT JOIN users b ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" ORDER BY a.titulo ASC"
	log.Println(sql)
	rows, _ := Db.Query(sql)
	var itens []mdl.Item
	var item mdl.Item
	for rows.Next() {
		rows.Scan(&item.Id)
		rows.Scan(&item.ElementoId)
		rows.Scan(&item.Titulo)
		rows.Scan(&item.Descricao)
		rows.Scan(&item.Avaliacao)
		rows.Scan(&item.AuthorId)
		rows.Scan(&item.AuthorName)
		rows.Scan(&item.CDataCriacao)
		rows.Scan(&item.StatusId)
		rows.Scan(&item.CStatus)
		itens = append(itens, item)
		log.Println(itens)
	}
	return itens
}

/*
func DeleteItemsByElementoHandler(elementoId string) {
	sqlStatement := "DELETE FROM Items WHERE elemento_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(elementoId)
	log.Println("DELETE Items in Order Id: " + elementoId)
}

func UpdateItemsHandler(itemsPage []mdl.Item, itemsDB []mdl.Item) {
	for i := range itemsPage {
		id := itemsPage[i].Id
		for j := range itemsDB {
			if itemsDB[j].Id == id {
				fieldsChanged := hasSomeFieldChanged(itemsPage[i], itemsDB[j]) //DONE
				if fieldsChanged {
					updateItemHandler(itemsPage[i], itemsDB[j]) // TODO
				}
				break
			}
		}
	}
}

func hasSomeFieldChanged(itemPage mdl.Item, itemDB mdl.Item) bool {
	if itemPage.Qtt != itemDB.Qtt {
		return true
	} else if itemPage.Price != itemDB.Price {
		return true
	} else if itemPage.ItemValue != itemDB.ItemValue {
		return true
	} else {
		return false
	}
}

func updateItemHandler(i mdl.Item, itemDB mdl.Item) {
	sqlStatement := "UPDATE items SET " +
		"beer_id=$1, price=$2, quantity=$3, " +
		"item_value=$4 WHERE id=$5"
	updtForm, _ := Db.Prepare(sqlStatement)
	log.Println(i.Price)
	log.Println(i.Qtt)
	log.Println(i.ItemValue)
	log.Println(i.Id)
	_, err := updtForm.Exec(i.Price, i.Qtt, i.ItemValue, i.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeleteItemsHandler(diffDB []mdl.Item) {
	sqlStatement := "DELETE FROM items WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Item Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func CreateNewItemHandler(w http.ResponseWriter, diffPage []mdl.Item) {
	itemid := 0
	var item mdl.Item
	for i := range diffPage {
		item = diffPage[i]
		log.Println(item)
		sqlStatement := "INSERT INTO items(meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
		log.Println(sqlStatement)
		//		err := Db.QueryRow(sqlStatement, item.MealId, item.QtdMeasure, item.Qtd, item.Cho, item.Kcal, item.FoodId).Scan(&itemid)
		//		sec.CheckInternalServerError(err, w)
		//		if err != nil {
		//			panic(err.Error())
		//		}
		//		sec.CheckInternalServerError(err, w)
		log.Println("itemid: " + strconv.Itoa(itemid))
	}
}*/
