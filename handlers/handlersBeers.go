package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateBeerHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Beer")
	if r.Method == "POST" {
		name := r.FormValue("Name")
		qtd := r.FormValue("Qtd")
		price := r.FormValue("Price")
		sqlStatement := "INSERT INTO beers(name, qtd, price) VALUES ($1, $2, $3) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, qtd, price).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
	}
	http.Redirect(w, r, route.BeersRoute, 301)
}

func UpdateBeerHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Beer")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		qtd := r.FormValue("Qtd")
		price := r.FormValue("Price")
		sqlStatement := "UPDATE beers SET name=$1, qtd=$2, price=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, qtd, price, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
	}
	http.Redirect(w, r, route.BeersRoute, 301)
}

func DeleteBeerHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Beer")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM beers WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.BeersRoute, 301)
}

func ListBeersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Beers")
	//sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT id, name, qtd, price FROM beers order by id asc")
	sec.CheckInternalServerError(err, w)
	var beers []mdl.Beer
	var beer mdl.Beer
	var i = 1
	for rows.Next() {
		err = rows.Scan(&beer.Id, &beer.Name, &beer.Qtd, &beer.Price)
		sec.CheckInternalServerError(err, w)
		beer.Order = i
		i++
		beers = append(beers, beer)
	}
	var page mdl.PageBeers
	page.Beers = beers
	page.Title = "Cervejas"
	var tmpl = template.Must(template.ParseGlob("tiles/beers/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Beers", page)
	sec.CheckInternalServerError(err, w)
}
