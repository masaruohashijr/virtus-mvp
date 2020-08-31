package handlers

import (
	mdl "beerwh/models"
	//	sec "beerwh/security"
	//	"reflect"
	//	pq "github.com/lib/pq"
	//	"html/template"
	//	"fmt"
	"log"
	"net/http"
	"strconv"
	//	"strings"
	//	"time"
)

// AJAX
func ListItemsHandler(idOrder string) []mdl.Item {
	log.Println("List Items By Order Id")
	sql := "SELECT A.id, A.order_id, A.beer_id, B.name, " +
		" A.quantity, A.price, A.item_value " +
		" FROM items A, beers B WHERE A.order_id= $1 AND A.beer_id = B.id"
	log.Println(sql)
	rows, _ := Db.Query(sql, idOrder)
	var items []mdl.Item
	var item mdl.Item
	for rows.Next() {
		rows.Scan(&item.Id, &item.IdOrder, &item.BeerId, &item.BeerName, &item.Qtt, &item.Price, &item.ItemValue)
		items = append(items, item)
		log.Println(items)
	}
	return items
}

func DeleteItemsByOrderHandler(orderId string) {
	sqlStatement := "DELETE FROM Items WHERE order_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(orderId)
	log.Println("DELETE Items in Order Id: " + orderId)
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
	} else if itemPage.BeerId != itemDB.BeerId {
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
	updtForm.Exec(i.BeerId, i.Price, i.Qtt, i.ItemValue, i.Id)
	log.Println("UPDATE: " + sqlStatement)
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
}
