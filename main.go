package main

import (
	hd "beer-warehouse/handlers"
	"database/sql"
	//	sec "diaria/security"
	route "beer-warehouse/routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	db  *sql.DB
	err error
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/aria?sslmode=disable")
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
	// injetando a variável Authenticated
	hd.Db = database
	http.HandleFunc("/", hd.IndexHandler)
	http.HandleFunc("/login", hd.LoginHandler)
	// ----------------- FOODS
	http.HandleFunc(route.FoodsRoute, hd.ListFoodsHandler)
	http.HandleFunc("/createFood", hd.CreateFoodHandler)
	http.HandleFunc("/updateFood", hd.UpdateFoodHandler)
	http.HandleFunc("/deleteFood", hd.DeleteFoodHandler)
	// ----------------- MEAL TYPES
	http.HandleFunc(route.MealTypesRoute, hd.ListMealTypesHandler)
	http.HandleFunc("/createMealType", hd.CreateMealTypeHandler)
	http.HandleFunc("/updateMealType", hd.UpdateMealTypeHandler)
	http.HandleFunc("/deleteMealType", hd.DeleteMealTypeHandler)
	// ----------------- MEASURES
	http.HandleFunc(route.MeasuresRoute, hd.ListMeasuresHandler)
	http.HandleFunc("/createMeasure", hd.CreateMeasureHandler)
	http.HandleFunc("/updateMeasure", hd.UpdateMeasureHandler)
	http.HandleFunc("/deleteMeasure", hd.DeleteMeasureHandler)
	// ----------------- MEALS
	http.HandleFunc(route.MealsRoute, hd.ListarMealsHandler)
	http.HandleFunc("/createMeal", hd.CreateMealHandler)
	http.HandleFunc("/updateMeal", hd.UpdateMealHandler)
	http.HandleFunc("/deleteMeal", hd.DeleteMealHandler)
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.ListenAndServe(":5000", nil)
	defer database.Close()
}
