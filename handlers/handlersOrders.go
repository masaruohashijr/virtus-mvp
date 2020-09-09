package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Order")

	if r.Method == "POST" {
		orderedDate := r.FormValue("OrderDateForInsert")
		orderedAt := r.FormValue("OrderedAtForInsert")
		takeOutDate := r.FormValue("TakeOutDateForInsert")
		takeOutAt := r.FormValue("TakeOutAtForInsert")
		orderedAt = orderedDate + " " + orderedAt
		takeOutAt = takeOutDate + " " + takeOutAt
		log.Println("orderedAt: " + orderedAt)
		log.Println("takeOutAt: " + takeOutAt)
		sqlStatement := "INSERT INTO public.orders ( " +
			" user_id, ordered_at, take_out_at ) " +
			" VALUES ($1, TO_TIMESTAMP($2, 'YYYY-MM-DD HH24:MI:SS'), TO_TIMESTAMP($3, 'YYYY-MM-DD HH24:MI:SS')) RETURNING id"
		orderId := 0
		var user mdl.User
		session, _ := store.Get(r, "beerwh")
		sessionUser := session.Values["user"]
		if sessionUser != nil {
			strUser := sessionUser.(string)
			json.Unmarshal([]byte(strUser), &user)
			err := Db.QueryRow(sqlStatement, user.Id, orderedAt, takeOutAt).Scan(&orderId)
			sec.CheckInternalServerError(err, w)
			if err != nil {
				panic(err.Error())
			}
		}
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				beerId := strings.Split(array[1], ":")[1]
				qtd := extraiValor(strings.Split(array[3], ":"))
				price := extraiValor(strings.Split(array[4], ":"))
				itemValue := extraiValor(strings.Split(array[5], ":"))
				itemId := 0
				log.Println("beerId: " + beerId)
				sqlStatement := "INSERT INTO items(order_id, beer_id, quantity, price, item_value) VALUES ($1,$2,$3,$4,$5) RETURNING id"
				err := Db.QueryRow(sqlStatement, orderId, beerId, qtd, price, itemValue).Scan(&itemId)
				sec.CheckInternalServerError(err, w)
				if err != nil {
					panic(err.Error())
				}
				sec.CheckInternalServerError(err, w)
				l := "INSERT: Id: " + strconv.Itoa(orderId)
				l += " | OrderedDate: " + orderedDate
				l += " | OrderedAt: " + orderedAt
				l += " | TakeOutDate: " + takeOutDate
				l += " | TakeOutAt: " + takeOutAt
				log.Println(l)
			}
		}
	}
	http.Redirect(w, r, route.OrdersRoute, 301)
}

func extraiValor(arr []string) string {
	r := "0.0"
	if len(arr) > 1 && arr[1] != "" {
		r = arr[1]
	}
	return r
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Order")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Items WHERE order_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM Orders WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.OrdersRoute, 301)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Order")
	if r.Method == "POST" {
		sec.IsAuthenticated(w, r)
		orderId := r.FormValue("Id")
		log.Println("OrderId: " + orderId)
		userId := r.FormValue("UserForUpdate")
		orderedDate := r.FormValue("OrderDateForUpdate")
		orderedAt := r.FormValue("OrderedAtForUpdate")
		orderedDT := orderedDate + " " + orderedAt
		takeoutDate := r.FormValue("TakeOutDateForUpdate")
		takeoutAt := r.FormValue("TakeOutAtForUpdate")
		takeoutDT := takeoutDate + " " + takeoutAt
		sqlStatement := "UPDATE orders SET " +
			" user_id = $1, " +
			" ordered_at = TO_TIMESTAMP($2, 'YYYY-MM-DD HH24:MI:SS'), " +
			" take_out_at = TO_TIMESTAMP($3, 'YYYY-MM-DD HH24:MI:SS') " +
			" WHERE id = $4"
		log.Println(sqlStatement)
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(userId, orderedDT, takeoutDT, orderId)
		log.Println("UPDATE: Id: " + orderId +
			" | User Id: " + userId +
			" | Ordered At: " + orderedDT +
			" | Take Out At: " + takeoutDT)

		var itemsDB = ListItemsHandler(orderId)
		var itemsPage []mdl.Item
		var itemPage mdl.Item
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				itemid := strings.Split(array[0], ":")[1]
				beerid := strings.Split(array[1], ":")[1]
				qtd := extraiValor(strings.Split(array[3], ":"))
				price := extraiValor(strings.Split(array[4], ":"))
				itemValue := extraiValor(strings.Split(array[5], ":"))
				n, _ := strconv.ParseInt(itemid, 10, 64)
				itemPage.Id = n
				m, _ := strconv.ParseInt(beerid, 10, 64)
				itemPage.BeerId = m
				q, _ := strconv.ParseFloat(qtd, 64)
				itemPage.Qtt = q
				r, _ := strconv.ParseFloat(price, 64)
				itemPage.Price = r
				s, _ := strconv.ParseFloat(itemValue, 64)
				itemPage.ItemValue = s
				itemsPage = append(itemsPage, itemPage)
			}
		}
		if len(itemsPage) < len(itemsDB) {
			log.Println("Quantidade de Itens da Página: " + strconv.Itoa(len(itemsPage)))
			if len(itemsPage) == 0 {
				DeleteItemsByOrderHandler(orderId) //DONE
			} else {
				var diffDB []mdl.Item = itemsDB
				for n := range itemsPage {
					if containsItem(diffDB, itemsPage[n]) {
						diffDB = removeItem(diffDB, itemsPage[n])
					}
				}
				DeleteItemsHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Item = itemsPage
			for n := range itemsDB {
				if containsItem(diffPage, itemsDB[n]) {
					diffPage = removeItem(diffPage, itemsDB[n])
				}
			}
			itemId := 0
			var item mdl.Item
			for i := range diffPage {
				item = diffPage[i]
				log.Println("Order Id: " + strconv.FormatInt(item.IdOrder, 10))
				sqlStatement := "INSERT INTO items(order_id, beer_id, quantity, price, item_value) VALUES ($1,$2,$3,$4,$5) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(sqlStatement, orderId, item.BeerId, item.Qtt, item.Price, item.ItemValue).Scan(&itemId)
				sec.CheckInternalServerError(err, w)
				if err != nil {
					panic(err.Error())
				}
				sec.CheckInternalServerError(err, w)
			}
		}
		UpdateItemsHandler(itemsPage, itemsDB) // TODO
		http.Redirect(w, r, route.OrdersRoute, 301)
	}
}

func containsItem(items []mdl.Item, itemCompared mdl.Item) bool {
	for n := range items {
		if items[n].Id == itemCompared.Id {
			return true
		}
	}
	return false
}

func removeItem(items []mdl.Item, itemToBeRemoved mdl.Item) []mdl.Item {
	var newItems []mdl.Item
	for i := range items {
		if items[i].Id != itemToBeRemoved.Id {
			newItems = append(newItems, items[i])
		}
	}
	return newItems
}

func LoadItemsByOrderId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Items By Order Id")
	r.ParseForm()
	var idOrder = r.FormValue("idOrder")
	log.Println("idOrder: " + idOrder)
	items := ListItemsHandler(idOrder)
	jsonItems, _ := json.Marshal(items)
	w.Write([]byte(jsonItems))
	log.Println("JSON")
}

func ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Orders")
	sec.IsAuthenticated(w, r)
	query := "SELECT a.id, a.user_id, b.name, a.ordered_at, a.take_out_at, " +
		" coalesce(to_char(a.ordered_at,'DD/MM/YYYY'),'') as c_ordered_date," +
		" coalesce(to_char(a.take_out_at,'DD/MM/YYYY'),'') as c_takeout_date," +
		" coalesce(to_char(a.ordered_at,'DD/MM/YYYY HH24:MI:SS'),'') as c_ordered_date_time," +
		" coalesce(to_char(a.take_out_at,'DD/MM/YYYY HH24:MI:SS'),'') as c_takeout_date_time" +
		" FROM orders a, users b where a.user_id = b.id order by a.take_out_at desc "
	rows, err := Db.Query(query)
	log.Println("Query: " + query)
	sec.CheckInternalServerError(err, w)
	var orders []mdl.Order
	var order mdl.Order
	var i = 1
	for rows.Next() {
		err = rows.Scan(
			&order.Id,
			&order.UserId,
			&order.UserName,
			&order.OrderedAt,
			&order.TakeOutAt,
			&order.COrderedAt,
			&order.CTakeOutAt,
			&order.COrderedDateTime,
			&order.CTakeOutDateTime,
		)
		sec.CheckInternalServerError(err, w)
		order.Order = i
		i++
		orders = append(orders, order)
	}
	rows, err = Db.Query("SELECT id, name FROM users order by name")
	sec.CheckInternalServerError(err, w)
	var users []mdl.User
	var user mdl.User
	var savedUser mdl.User
	session, _ := store.Get(r, "beerwh")
	sessionUser := session.Values["user"]
	if sessionUser != nil {
		strUser := sessionUser.(string)
		json.Unmarshal([]byte(strUser), &savedUser)
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if user.Id == savedUser.Id {
			user.Selected = true
		} else {
			user.Selected = false
		}
		sec.CheckInternalServerError(err, w)
		users = append(users, user)
	}
	var page mdl.PageOrders
	page.Users = users
	rows, err = Db.Query("SELECT id, name, qtd, price FROM beers order by name")
	sec.CheckInternalServerError(err, w)
	var beers []mdl.Beer
	var beer mdl.Beer
	for rows.Next() {
		err = rows.Scan(&beer.Id, &beer.Name, &beer.Qtd, &beer.Price)
		sec.CheckInternalServerError(err, w)
		beers = append(beers, beer)
	}
	page.Beers = beers
	page.Orders = orders
	page.Title = "Pedidos"
	var tmpl = template.Must(template.ParseGlob("tiles/orders/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Orders", page)
	sec.CheckInternalServerError(err, w)
}
