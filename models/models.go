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
	Descricao  string `json:"descricao"`
	Stereotype string `json:"stereotype"`
}

type Item struct {
	Order        int    `json:"order"`
	Id           int64  `json:"id"`
	ElementoId   int64  `json:"elementoId"`
	Nome         string `json:"nome"`
	Descricao    string `json:"descricao"`
	Avaliacao    string `json:"avaliacao"`
	AuthorId     int64  `json:"autorId"`
	AuthorName   string `json:"autorNome"`
	CDataCriacao string `json:"dataCriacao"`
	StatusId     int64  `json:"status"`
	CStatus      string `json:"cStatus"`
}

type Elemento struct {
	Order        int
	Id           int64  `json:"id"`
	Nome         string `json:"nome"`
	Descricao    string `json:"descricao"`
	AuthorId     int64  `json:"authorId"`
	AuthorName   string `json:"authorName"`
	Peso         int    `json:"weight"`
	CDataCriacao string `json:"cDataCriacao`
	StatusId     int64  `json:"statusId"`
	CStatus      string `json:"cStatus"`
}

type Componente struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
}

type Matriz struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
}

type Ciclo struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
}

type CicloEntidade struct {
	Order          int
	Id             int64  `json:"id"`
	EntidadeId     int64  `json:"entidadeId"`
	CicloId        int64  `json:"cicloId"`
	Nome           string `json:"nome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	Nota           string `json:"nota"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Entidade struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
}

type Plano struct {
	Order        int
	Id           int64  `json:"id"`
	Nome         string `json:"nome"`
	Descricao    string `json:"descricao"`
	EntidadeId   int64  `json:"entidadeId"`
	EntidadeNome string `json:"entidadeNome"`
	AuthorId     int64  `json:"authorId"`
	AuthorName   string `json:"authorName"`
	CriadoEm     string `json:"criadoEm"`
	StatusId     int64  `json:"statusId"`
	CStatus      string `json:"cStatus"`
}

type Carteira struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
}

type Equipe struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
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

type PageComponentes struct {
	AppName     string
	Title       string
	UserId      int
	Componentes []Componente
	Users       []User
	LoggedUser  LoggedUser
}

type PageMatrizes struct {
	AppName    string
	Title      string
	UserId     int
	Matrizes   []Matriz
	Users      []User
	LoggedUser LoggedUser
}

type PageCiclos struct {
	AppName    string
	Title      string
	UserId     int
	Ciclos     []Ciclo
	Users      []User
	LoggedUser LoggedUser
}

type PageEntidades struct {
	AppName    string
	Title      string
	UserId     int
	Entidades  []Entidade
	Ciclos     []Ciclo
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PagePlanos struct {
	AppName    string
	Title      string
	UserId     int
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageEquipes struct {
	AppName    string
	Title      string
	UserId     int
	Equipes    []Equipe
	Users      []User
	LoggedUser LoggedUser
}
type PageCarteiras struct {
	AppName    string
	Title      string
	UserId     int
	Carteiras  []Carteira
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

type PageActions struct {
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageWorkflows struct {
	AppName    string
	Title      string
	Features   []Feature
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
