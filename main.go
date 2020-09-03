package main

import (
	hd "beerwh/handlers"
	dpk "beerwh/db"
	route "beerwh/routes"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	
)

func dbConn() *sql.DB{
	dbase, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	log.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	return dbase
}


func main() {

	hd.Db = dbConn()	
	// injeta	ndo a variável Authenticated
	dpk.Initialize()
	http.HandleFunc("/", hd.IndexHandler)
	http.HandleFunc("/login", hd.LoginHandler)
	// ----------------- BEERS
	http.HandleFunc(route.BeersRoute, hd.ListBeersHandler)
	http.HandleFunc("/createBeer", hd.CreateBeerHandler)
	http.HandleFunc("/updateBeer", hd.UpdateBeerHandler)
	http.HandleFunc("/deleteBeer", hd.DeleteBeerHandler)
	// ----------------- CLIENTS
	http.HandleFunc(route.ClientsRoute, hd.ListClientsHandler)
	http.HandleFunc("/createClient", hd.CreateClientHandler)
	http.HandleFunc("/updateClient", hd.UpdateClientHandler)
	http.HandleFunc("/deleteClient", hd.DeleteClientHandler)
	// ----------------- ORDERS
	http.HandleFunc(route.OrdersRoute, hd.ListOrdersHandler)
	http.HandleFunc("/createOrder", hd.CreateOrderHandler)
	http.HandleFunc("/updateOrder", hd.UpdateOrderHandler)
	http.HandleFunc("/deleteOrder", hd.DeleteOrderHandler)
	// ----------------- ITEMS
	http.HandleFunc("/loadItemsByOrderId", hd.LoadItemsByOrderId)
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.ListenAndServe(":5000", nil)
	defer hd.Db.Close()
}
