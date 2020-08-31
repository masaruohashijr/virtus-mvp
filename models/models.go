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
	ClientId         int64     `json:"clientId"`
	ClientName       string    `json:"clientName`
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

type Client struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
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

type PageOrders struct {
	Title   string
	UserId  int
	Orders  []Order
	Beers   []Beer
	Clients []Client
}

type PageClients struct {
	Title   string
	Clients []Client
}

type PageBeers struct {
	Title string
	Beers []Beer
}
