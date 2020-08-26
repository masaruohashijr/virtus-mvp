package handlers

import (
	mdl "diaria/models"
	sec "diaria/security"
	"encoding/json"
	pq "github.com/lib/pq"
	//	htemplate "html/template"
	route "diaria/routes"
	"log"
	"net/http"
	"strconv"
	"strings"
	ttemplate "text/template"
	"time"
)

func CreateMealHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Meal")

	if r.Method == "POST" {
		mealType := r.FormValue("MealTypeForInsert")
		date := r.FormValue("DateForInsert")
		startAt := r.FormValue("StartAtForInsert")
		endAt := r.FormValue("EndAtForInsert")
		bolus := r.FormValue("BolusForInsert")
		if bolus == "" {
			bolus = "0"
		}
		sqlStatement := "INSERT INTO meals(meal_type_id, date, start_at, end_at, bolus) VALUES ($1,$2,$3,$4,$5) RETURNING id"
		mealId := 0
		var err error
		if endAt == "" {
			err = Db.QueryRow(sqlStatement, mealType, date, startAt, pq.NullTime{}, bolus).Scan(&mealId)
		} else {
			err = Db.QueryRow(sqlStatement, mealType, date, startAt, endAt, bolus).Scan(&mealId)
		}
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}

		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				//				log.Println(value)
				array := strings.Split(value[0], "#")
				foodid := strings.Split(array[2], ":")[1]
				//				alimento := array[2]
				qtdMedida := extraiValor(strings.Split(array[4], ":"))
				qtd := extraiValor(strings.Split(array[5], ":"))
				cho := extraiValor(strings.Split(array[6], ":"))
				kcal := extraiValor(strings.Split(array[7], ":"))
				itemid := 0
				log.Println("foodid: " + foodid)
				sqlStatement := "INSERT INTO items(meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
				err := Db.QueryRow(sqlStatement, mealId, qtdMedida, qtd, cho, kcal, foodid).Scan(&itemid)
				sec.CheckInternalServerError(err, w)
				if err != nil {
					panic(err.Error())
				}
			}
		}

		sec.CheckInternalServerError(err, w)
		l := "INSERT: Id: " + strconv.Itoa(mealId)
		l += " | Date: " + date
		l += " | MealType: " + mealType
		l += " | StartAt: " + startAt
		l += " | endAt: " + endAt
		l += " | Bolus: " + bolus
		log.Println(l)
	}
	http.Redirect(w, r, route.MealsRoute, 301)
}

func extraiValor(arr []string) string {
	r := "0.0"
	if len(arr) > 1 && arr[1] != "" {
		r = arr[1]
	}
	return r
}

func DeleteMealHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Meal")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM Items WHERE meal_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM Meals WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.MealsRoute, 301)
}

func UpdateMealHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Meal")
	if r.Method == "POST" {
		sec.IsAuthenticated(w, r)
		mealId := r.FormValue("Id")
		mealType := r.FormValue("MealTypeForUpdate")
		date := r.FormValue("DateForUpdate")
		startAt := r.FormValue("StartAtForUpdate")
		endAt := r.FormValue("EndAtForUpdate")
		bolus := r.FormValue("BolusForUpdate")
		sqlStatement := "UPDATE meals SET " +
			" meal_type_id = $1, " +
			" date = $2, " +
			" start_at = $3, " +
			" end_at = $4, " +
			" bolus = $5 " +
			" WHERE id = $6"
		log.Println(sqlStatement)
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		if endAt == "" {
			updtForm.Exec(mealType, date, startAt, pq.NullTime{}, bolus, mealId)
		} else {
			updtForm.Exec(mealType, date, startAt, endAt, bolus, mealId)
		}
		log.Println("UPDATE: Id: " + mealId +
			" | MealTypeId: " + mealType +
			" | Date: " + date +
			" | Start At: " + startAt +
			" | End At: " + endAt +
			" | Bolus: " + bolus)

		var itemsDB = ListItemsHandler(mealId)
		var itemsPage []mdl.Item
		var itemPage mdl.Item
		for key, value := range r.Form {
			if strings.HasPrefix(key, "item") {
				array := strings.Split(value[0], "#")
				//				log.Println(array)
				itemid := strings.Split(array[0], ":")[1]
				foodid := strings.Split(array[2], ":")[1]
				qtdMeasure := extraiValor(strings.Split(array[3], ":"))
				qtd := extraiValor(strings.Split(array[5], ":"))
				cho := extraiValor(strings.Split(array[6], ":"))
				kcal := extraiValor(strings.Split(array[7], ":"))
				n, _ := strconv.ParseInt(itemid, 10, 64)
				itemPage.Id = n
				m, _ := strconv.ParseInt(foodid, 10, 64)
				itemPage.FoodId = m
				o, _ := strconv.ParseInt(mealId, 10, 64)
				itemPage.MealId = o
				p, _ := strconv.ParseFloat(qtdMeasure, 64)
				itemPage.QtdMeasure = p
				q, _ := strconv.ParseFloat(qtd, 64)
				itemPage.Qtd = q
				r, _ := strconv.ParseFloat(cho, 64)
				itemPage.Cho = r
				s, _ := strconv.ParseFloat(kcal, 64)
				itemPage.Kcal = s
				itemsPage = append(itemsPage, itemPage)
			}
		}
		if len(itemsPage) < len(itemsDB) {
			log.Println("Quantidade de Itens da Página: " + strconv.Itoa(len(itemsPage)))
			if len(itemsPage) == 0 {
				DeleteItemsByMealHandler(mealId)
			} else {
				var diffDB []mdl.Item = itemsDB
				for n := range itemsPage {
					if contains(diffDB, itemsPage[n]) {
						diffDB = remove(diffDB, itemsPage[n])
					}
				}
				DeleteItemsHandler(diffDB)
			}
		} else {
			var diffPage []mdl.Item = itemsPage
			for n := range itemsDB {
				if contains(diffPage, itemsDB[n]) {
					diffPage = remove(diffPage, itemsDB[n])
				}
			}
			//			log.Println("CreateNewItemHandler")
			itemId := 0
			var item mdl.Item
			for i := range diffPage {
				item = diffPage[i]
				//				log.Println(item)
				sqlStatement := "INSERT INTO items(meal_id, food_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
				log.Println(sqlStatement)
				err := Db.QueryRow(sqlStatement, item.MealId, item.FoodId, item.QtdMeasure, item.Qtd, item.Cho, item.Kcal).Scan(&itemId)
				sec.CheckInternalServerError(err, w)
				if err != nil {
					panic(err.Error())
				}
				sec.CheckInternalServerError(err, w)
				// log.Println("itemid: " + strconv.Itoa(itemId))
			}
		}
		UpdateItemsHandler(itemsPage, itemsDB) // TODO Comparando campo a campo os elementos de intersecção.
		http.Redirect(w, r, route.MealsRoute, 301)
	} else {
		r.ParseForm()
		var idMeal = r.FormValue("idMeal")
		log.Println(idMeal)
		items := ListItemsHandler(idMeal)
		jsonItems, _ := json.Marshal(items)
		w.Write([]byte(jsonItems))
		log.Println("JSON")
	}
}

func contains(items []mdl.Item, itemCompared mdl.Item) bool {
	for n := range items {
		if items[n].Id == itemCompared.Id {
			return true
		}
	}
	return false
}

func remove(items []mdl.Item, itemToBeRemoved mdl.Item) []mdl.Item {
	var newItems []mdl.Item
	for i := range items {
		if items[i].Id != itemToBeRemoved.Id {
			newItems = append(newItems, items[i])
		}
	}
	return newItems
}

func ListarMealsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Listar Meals")
	sec.IsAuthenticated(w, r)
	query := "SELECT " +
		"R1.meal_id, " +
		"R1.meal_type_id, " +
		"R1.meal_type_name, " +
		"R1.meal_date, " +
		"R1.start_at, " +
		"R1.end_at, " +
		"R1.c_meal_date, " +
		"R1.c_start_at, " +
		"R1.c_end_at, " +
		"R2.cho_total," +
		"R2.kcal_total, " +
		"R1.bolus " +
		"FROM " +
		"( " +
		"SELECT  " +
		"a.id as meal_id, b.id as meal_type_id, b.name as meal_type_name, " +
		"a.date as meal_date, a.start_at as start_at, a.end_at as end_at, " +
		"coalesce(to_char(a.date,'DD/MM/YYYY'),'') as c_meal_date, " +
		"coalesce(to_char(a.start_at,'HH24:MI:SS'),'') as c_start_at, " +
		"coalesce(to_char(a.end_at,'HH24:MI:SS'),'') as c_end_at, " +
		"coalesce(a.bolus,0.00) as bolus FROM " +
		"meals a, meal_types b " +
		"WHERE a.meal_type_id = b.id " +
		"order by a.id desc " +
		") R1, " +
		"(SELECT a.id as meal_id, " +
		"sum(b.cho) as cho_total, " +
		"sum(b.kcal) as kcal_total " +
		"from meals a, items b " +
		"where a.id = b.meal_id " +
		"group by a.id " +
		"order by a.id desc " +
		" ) R2 " +
		"WHERE R1.meal_id = R2.meal_id "

	log.Println("QUERY: " + query)
	rows, err := Db.Query(query)
	sec.CheckInternalServerError(err, w)
	var funcMap = ttemplate.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	var meals []mdl.Meal
	var meal mdl.Meal
	var i = 1
	for rows.Next() {
		err = rows.Scan(
			&meal.Id,
			&meal.MealTypeId,
			&meal.MealTypeName,
			&meal.Date,
			&meal.StartAt,
			&meal.EndAt,
			&meal.CDate,
			&meal.CStartAt,
			&meal.CEndAt,
			&meal.CCho,
			&meal.CKcal,
			&meal.Bolus)
		sec.CheckInternalServerError(err, w)
		meal.Order = i
		i++
		meals = append(meals, meal)
	}
	rows, err = Db.Query("SELECT id, name, start_at, end_at FROM meal_types")
	sec.CheckInternalServerError(err, w)
	var mealTypes []mdl.MealType
	var mealType mdl.MealType
	now := GetNow()
	for rows.Next() {
		err = rows.Scan(&mealType.Id, &mealType.Name, &mealType.StartAt, &mealType.EndAt)
		sec.CheckInternalServerError(err, w)
		if mealType.EndAt.Before(mealType.StartAt) {
			if mealType.StartAt.Before(now) && GetMidnight().After(now) {
				mealType.Selected = true
			} else {
				if GetMidnight().Before(now) && mealType.EndAt.After(now) {
					mealType.Selected = true
				} else {
					mealType.Selected = false
				}
			}
		} else {
			if mealType.StartAt.Before(now) && mealType.EndAt.After(now) {
				mealType.Selected = true
			} else {
				mealType.Selected = false
			}
		}
		mealTypes = append(mealTypes, mealType)
	}
	rows, err = Db.Query("SELECT id, name, measure, qtd, cho, kcal FROM foods order by id asc")
	sec.CheckInternalServerError(err, w)
	var foods []mdl.Food
	var food mdl.Food
	for rows.Next() {
		err = rows.Scan(&food.Id, &food.Name, &food.Measure, &food.Qtd, &food.Cho, &food.Kcal)
		sec.CheckInternalServerError(err, w)
		foods = append(foods, food)
	}
	var page mdl.PageMeals
	page.Meals = meals
	page.MealTypes = mealTypes
	page.Foods = foods
	page.Title = "Refeições"
	var tmpl = ttemplate.Must(ttemplate.ParseGlob("tiles/meals/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.Funcs(funcMap)
	tmpl.ExecuteTemplate(w, "Main-Meal", page)
	sec.CheckInternalServerError(err, w)
}

func GetNow() time.Time {
	now := time.Now().String()
	txtNow := strings.Split(strings.Split(strings.Split(now, " ")[1], ".")[0], ":")
	hora, _ := strconv.Atoi(txtNow[0])
	minuto, _ := strconv.Atoi(txtNow[1])
	segundo, _ := strconv.Atoi(txtNow[2])
	t := time.Date(0000, time.January, 1,
		hora,
		minuto,
		segundo, 0, time.UTC)
	return t
}

func GetMidnight() time.Time {
	t := time.Date(0000, time.January, 1, 24, 0, 0, 0, time.UTC)
	return t
}
