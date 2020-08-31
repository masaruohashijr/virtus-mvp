package main

import (
	hd "beerwh/handlers"
	route "beerwh/routes"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	db  *sql.DB
	err error
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/beerwh?sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		panic(err)
	}
	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	database := dbConn()
	log.Println("O database está disponível.")
	// injeta	ndo a variável Authenticated
	hd.Db = database
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
	defer database.Close()
}
