package models

import (
	"github.com/lib/pq"
	"time"
)

type Role struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Selected bool
	Features []Feature
}

type Feature struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
}

type Workflow struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type Action struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type Status struct {
	Order      int
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Stereotype string `json:"stereotype"`
}

type Beer struct {
	Order int
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Qtd   int     `json:"qtd"`
	Price float64 `json:"price"`
}

type Item struct {
	Order     int
	Id        int64   `json:"id"`
	IdOrder   int64   `json:"orderId"`
	BeerId    int64   `json:"beerId"`
	BeerName  string  `json:"beerName"`
	Qtt       float64 `json:"qtd"`
	Price     float64 `json:"price"`
	ItemValue float64 `json:"value"`
}

type Order struct {
	Order            int
	Id               int64     `json:"id"`
	UserId           int64     `json:"clientId"`
	UserName         string    `json:"clientName`
	OrderedAt        time.Time `json:"orderedAt`
	TakeOutAt        time.Time `json:"endAt"`
	COrderedDateTime string    `json:"corderedDateTime`
	CTakeOutDateTime string    `json:"ctakeOutDateTime`
	COrderedAt       string    `json:"corderedAt`
	CTakeOutAt       string    `json:"ctakeOutAt`
}

type NullTime struct {
	pq.NullTime
}

type User struct {
	Order         int
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Role          int64  `json:"role"`
	RoleName      string `json:"rolename"`
	HasPermission func(string) bool
	Selected      bool
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

type PageOrders struct {
	Title  string
	UserId int
	Orders []Order
	Beers  []Beer
	Users  []User
}

type PageUsers struct {
	Title string
	Users []User
	Roles []Role
}

type PageBeers struct {
	Title string
	Beers []Beer
}

type PageRoles struct {
	Title    string
	Roles    []Role
	Features []Feature
}

type PageFeatures struct {
	Title    string
	Features []Feature
}

type PageStatus struct {
	Title  string
	Status []Status
}

type PageAction struct {
	Title   string
	Actions []Action
}

type PageWorkflow struct {
	Title     string
	Workflows []Workflow
}
