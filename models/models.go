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
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
	Features       []Feature
}

type Feature struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type Workflow struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	EntityType     string `json:"entity"`
	StartAt        string `json:"startAt`
	EndAt          string `json:"endAt"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
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
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OriginId       int64  `json:"originid"`
	Origin         string `json:"originName"`
	DestinationId  int64  `json:"destinationid"`
	Destination    string `json:"destinationName"`
	OtherThan      bool   `json:"otherthan"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
	Roles          []Role
}

type Status struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Stereotype     string `json:"stereotype"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type Item struct {
	Order          int    `json:"order"`
	Id             int64  `json:"id"`
	ElementoId     int64  `json:"elementoId"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Avaliacao      string `json:"avaliacao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"status"`
	CStatus        string `json:"cStatus"`
}

type Elemento struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	Peso           int    `json:"weight"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ElementoComponente struct {
	Order          int
	Id             int64  `json:"id"`
	ComponenteId   int64  `json:"componenteId"`
	ElementoId     int64  `json:"elementoId"`
	ElementoNome   string `json:"elementoNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	PesoPadrao     int    `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Componente struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ComponentePilar struct {
	Order          int
	Id             int64  `json:"id"`
	PilarId        int64  `json:"pilarId"`
	ComponenteId   int64  `json:"componenteId"`
	ComponenteNome string `json:"componenteNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	PesoPadrao     int    `json:"pesoPadrao"`
	Sonda          string `json:"sonda"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Pilar struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
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
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type CicloEntidade struct {
	Order          int
	Id             int64  `json:"id"`
	EntidadeId     int64  `json:"entidadeId"`
	CicloId        int64  `json:"cicloId"`
	Nome           string `json:"nome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type PilarCiclo struct {
	Order          int
	Id             int64  `json:"id"`
	CicloId        int64  `json:"cicloId"`
	PilarId        int64  `json:"pilarId"`
	PilarNome      string `json:"pilarNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	PesoPadrao     int    `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Entidade struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	ChefeId        int64  `json:"chefeId"`
	ChefeName      string `json:"chefeName"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Plano struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	EntidadeId     int64  `json:"entidadeId"`
	EntidadeNome   string `json:"entidadeNome"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Carteira struct {
	Order    int
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	AuthorId int64  `json:"authorId"`
	StatusId int64  `json:"statusId"`
	CStatus  string `json:"cStatus"`
}

type Escritorio struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	ChefeId        int64  `json:"chefeId"`
	ChefeNome      string `json:"chefeNome"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool   `json:"selected"`
}

type Jurisdicao struct {
	Order          int
	Id             int64  `json:"id"`
	EscritorioId   int64  `json:"escritorioId"`
	EntidadeId     int64  `json:"entidadeId"`
	EntidadeNome   string `json:"entidadeNome"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Membro struct {
	Order          int
	Id             int64  `json:"id"`
	EscritorioId   int64  `json:"escritorioId"`
	UsuarioId      int64  `json:"usuarioId"`
	UsuarioNome    string `json:"usuarioNome"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type NullTime struct {
	pq.NullTime
}

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type User struct {
	Order          int       `json:"order"`
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Mobile         string    `json:"mobile"`
	Escritorio     int64     `json:"escritorio"`
	EscritorioNome string    `json:"escritorioNome"`
	Role           int64     `json:"role"`
	RoleName       string    `json:"roleName"`
	AuthorId       int64     `json:"authorId"`
	AuthorName     string    `json:"authorName"`
	CriadoEm       string    `json:"criadoEm"`
	C_CriadoEm     string    `json:"c_criadoEm"`
	IdVersaoOrigem int64     `json:"idVersaoOrigem"`
	StatusId       int64     `json:"statusId"`
	CStatus        string    `json:"cStatus"`
	Features       []Feature `json:"features"`
	Selected       bool      `json:"selected"`
}

type PageElementos struct {
	AppName    string
	Title      string
	Elementos  []Elemento
	Users      []User
	LoggedUser LoggedUser
}

type PageComponentes struct {
	AppName     string
	Title       string
	Componentes []Componente
	Elementos   []Elemento
	Users       []User
	LoggedUser  LoggedUser
}

type PagePilares struct {
	AppName     string
	Title       string
	Pilares     []Pilar
	Componentes []Componente
	Users       []User
	LoggedUser  LoggedUser
}

type PageCiclos struct {
	AppName    string
	Title      string
	Ciclos     []Ciclo
	Pilares    []Pilar
	Users      []User
	LoggedUser LoggedUser
}

type PageEntidades struct {
	AppName    string
	Title      string
	Entidades  []Entidade
	Ciclos     []Ciclo
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PagePlanos struct {
	AppName    string
	Title      string
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageEscritorios struct {
	AppName     string
	Title       string
	Escritorios []Escritorio
	Entidades   []Entidade
	Users       []User
	LoggedUser  LoggedUser
}
type PageCarteiras struct {
	AppName    string
	Title      string
	Carteiras  []Carteira
	Users      []User
	LoggedUser LoggedUser
}

type PageUsers struct {
	AppName     string
	Title       string
	Users       []User
	Escritorios []Escritorio
	Roles       []Role
	LoggedUser  LoggedUser
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
	Statuss    []Status
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
