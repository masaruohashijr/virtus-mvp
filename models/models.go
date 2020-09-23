package models

import (
	"github.com/lib/pq"
)

var AppName = "Virtus"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

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
	Order      int
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	EntityType string `json:"entity"`
	StartAt    string `json:"startAt`
	EndAt      string `json:"endAt"`
}

type Activity struct {
	Order                int    `json:"order"`
	Id                   int64  `json:"id"`
	WorkflowId           int64  `json:"wid"`
	ActionId             int64  `json:"actionId"`
	ActionName           string `json:"actionName"`
	ExpirationActionId   int64  `json:"expActionId"`
	ExpirationActionName string `json:"expActionName"`
	ExpirationTimeDays   int    `json:"expTime"`
	CStartAt             string `json:"startAt"`
	CEndAt               string `json:"endAt"`
	CRoles               string `json:"roles"`
	CRoleNames           string `json:"roleNames"`
	Roles                []Role `json:"roles_array"`
	CFeatures            string `json:"features"`
	CFeatureNames        string `json:"featureNames"`
}

type Action struct {
	Order         int
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	OriginId      int64  `json:"originid"`
	Origin        string `json:"originName"`
	DestinationId int64  `json:"destinationid"`
	Destination   string `json:"destinationName"`
	OtherThan     bool   `json:"otherthan"`
	Roles         []Role
}

type Status struct {
	Order      int
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Stereotype string `json:"stereotype"`
}

type Item struct {
	Ordem      int     `json:"ordem"`
	Id         int64   `json:"id"`
	IdElemento int64   `json:"elementoId"`
	Qtt        float64 `json:"qtd"`
	Price      float64 `json:"price"`
	ItemValue  float64 `json:"value"`
}

type Elemento struct {
	Order    int
	Id       int64  `json:"id"`
	Titulo   string `json:"titulo`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId`
	CStatus  string `json:"cStatus`
}

type NullTime struct {
	pq.NullTime
}

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type User struct {
	Order    int       `json:"order"`
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Mobile   string    `json:"mobile"`
	Role     int64     `json:"role"`
	RoleName string    `json:"rolename"`
	Features []Feature `json:"features"`
	Selected bool      `json:"selected"`
}

type PageElementos struct {
	AppName    string
	Title      string
	UserId     int
	Elementos  []Elemento
	Users      []User
	LoggedUser LoggedUser
}

type PageUsers struct {
	AppName    string
	Title      string
	Users      []User
	Roles      []Role
	LoggedUser LoggedUser
}

type PageRoles struct {
	AppName    string
	Title      string
	Roles      []Role
	Features   []Feature
	LoggedUser LoggedUser
}

type PageFeatures struct {
	AppName    string
	Title      string
	Features   []Feature
	LoggedUser LoggedUser
}

type PageStatus struct {
	AppName    string
	Title      string
	Status     []Status
	LoggedUser LoggedUser
}

type PageAction struct {
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageWorkflow struct {
	AppName    string
	Title      string
	Features   []Feature
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
