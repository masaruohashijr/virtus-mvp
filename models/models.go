package models

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	Order    int
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}

type Food struct {
	Order   int
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Measure string  `json:"measure"`
	Qtd     float64 `json:"qtd"`
	Cho     float64 `json:"cho"`
	Kcal    float64 `json:"kcal"`
}

type Item struct {
	Order      int
	Id         int64   `json:"id"`
	MealId     int64   `json:"mealId"`
	FoodId     int64   `json:"foodId"`
	FoodName   string  `json:"foodName"`
	QtdMeasure float64 `json:"qtdMeasure"`
	Qtd        float64 `json:"qtd"`
	Cho        float64 `json:"cho"`
	Kcal       float64 `json:"kcal"`
}

type Meal struct {
	Order        int
	Id           int64    `json:"id"`
	MealTypeName string   `json:"mealTypeName"`
	MealTypeId   int64    `json:"mealTypeId"`
	Date         NullTime `json:"date"`
	StartAt      NullTime `json:"startAt"`
	EndAt        NullTime `json:"endAt"`
	CDate        string   `json:"cdate"`
	CStartAt     string   `json:"cstartAt"`
	CEndAt       string   `json:"cendAt"`
	CCho         string   `json:"ccho"`
	CKcal        string   `json:"ckcal"`
	Bolus        float64  `json:"bolus"`
}

type NullTime struct {
	pq.NullTime
}

type MealType struct {
	Order    int
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	StartAt  time.Time `json:"startAt"`
	EndAt    time.Time `json:"endAt"`
	CStartAt string    `json:"cstartAt"`
	CEndAt   string    `json:"cendAt"`
	Selected bool
}

type Measure struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type PageMeasures struct {
	Title    string
	Measures []Measure
}

type PageMeals struct {
	Title     string
	Meals     []Meal
	MealTypes []MealType
	Foods     []Food
}

type PageMealTypes struct {
	Title     string
	MealTypes []MealType
}

type PageFoods struct {
	Title string
	Foods []Food
}
