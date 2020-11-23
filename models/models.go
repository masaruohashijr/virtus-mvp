package models

import (
	"github.com/lib/pq"
)

var AppName = "Virtus"

type TipoPontuacao int

const (
	Manual TipoPontuacao = iota
	Calculada
	Ajustada
	Definitiva
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ColSpan struct {
	CicloId      int64 `json:"cicloId"`
	PilarId      int64 `json:"pilarId"`
	ComponenteId int64 `json:"componenteId"`
	Qtd          int   `json:"qtd"`
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
	Sonda          string `json:"sonda"`
	PesoPadrao     string `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
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
	TipoNotaId     int64  `json:"tipoNotaId"`
	TipoNotaNome   string `json:"tipoNotaNome"`
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
	Id             int64           `json:"id"`
	Nome           string          `json:"nome"`
	Descricao      string          `json:"descricao"`
	Sigla          string          `json:"sigla"`
	Codigo         string          `json:"codigo"`
	Situacao       string          `json:"situacao"`
	ESI            bool            `json:"esi"`
	Municipio      string          `json:"municipio"`
	SiglaUF        string          `json:"siglaUF"`
	ChefeId        int64           `json:"chefeId"`
	ChefeName      string          `json:"chefeName"`
	AuthorId       int64           `json:"authorId"`
	AuthorName     string          `json:"authorName"`
	CriadoEm       string          `json:"criadoEm"`
	C_CriadoEm     string          `json:"c_criadoEm"`
	IdVersaoOrigem int64           `json:"idVersaoOrigem"`
	StatusId       int64           `json:"statusId"`
	CStatus        string          `json:"cStatus"`
	CiclosEntidade []CicloEntidade `json:"ciclos"`
}

type Escritorio struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Abreviatura    string `json:"abreviatura"`
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

type Historico struct {
	Id            string `json:"id"`
	EntidadeId    string `json:"entidadeId"`
	CicloId       string `json:"cicloId"`
	PilarId       string `json:"pilarId"`
	ComponenteId  string `json:"componenteId"`
	ElementoId    string `json:"elementoId"`
	Nota          string `json:"nota"`
	Metodo        string `json:"metodo"`
	Peso          string `json:"peso"`
	AutorId       string `json:"autorId"`
	AutorNome     string `json:"autorNome"`
	AlteradoEm    string `json:"alteradoEm"`
	Motivacao     string `json:"motivacao"`
	TipoAlteracao string `json:"tipoAlteracao"`
}

type Item struct {
	Order          int    `json:"order"`
	Id             int64  `json:"id"`
	ElementoId     int64  `json:"elementoId"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"status"`
	CStatus        string `json:"cStatus"`
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

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type Membro struct {
	Order          int
	Id             int64  `json:"id"`
	EscritorioId   int64  `json:"escritorioId"`
	UsuarioId      int64  `json:"usuarioId"`
	UsuarioNome    string `json:"usuarioNome"`
	UsuarioPerfil  string `json:"usuarioPerfil"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Integrante struct {
	Order          int
	Id             int64  `json:"id"`
	EntidadeId     int64  `json:"entidadeId"`
	CicloId        int64  `json:"cicloId"`
	UsuarioId      int64  `json:"usuarioId"`
	UsuarioNome    string `json:"usuarioNome"`
	UsuarioPerfil  string `json:"usuarioPerfil"`
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

type Plano struct {
	Order             int
	Id                int64  `json:"id"`
	Nome              string `json:"nome"`
	Descricao         string `json:"descricao"`
	EntidadeId        int64  `json:"entidadeId"`
	EntidadeNome      string `json:"entidadeNome"`
	CNPB              string `json:"cnpb"`
	RecursoGarantidor string `json:"recursoGarantidor"`
	Modalidade        string `json:"modalidade"`
	AuthorId          int64  `json:"authorId"`
	AuthorName        string `json:"authorName"`
	CriadoEm          string `json:"criadoEm"`
	C_CriadoEm        string `json:"c_criadoEm"`
	IdVersaoOrigem    int64  `json:"idVersaoOrigem"`
	StatusId          int64  `json:"statusId"`
	CStatus           string `json:"cStatus"`
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

type PilarCiclo struct {
	Order          int
	Id             int64  `json:"id"`
	CicloId        int64  `json:"cicloId"`
	PilarId        int64  `json:"pilarId"`
	PilarNome      string `json:"pilarNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	PesoPadrao     string `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ProdutoCiclo struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    int64   `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}
type ProdutoPilar struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    int64   `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoComponente struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    string  `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	ComponenteId    int64   `json:"componenteId"`
	ComponenteNome  string  `json:"componenteNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoElemento struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    string  `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	ComponenteId    int64   `json:"componenteId"`
	ComponenteNome  string  `json:"componenteNome"`
	ElementoId      int64   `json:"elementoId"`
	ElementoNome    string  `json:"elementoNome"`
	TipoNotaId      int64   `json:"tipoNotaId"`
	TipoNotaNome    string  `json:"tipoNotaNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Peso            float64 `json:"peso"`
	Nota            int     `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoItem struct {
	Order            int
	Id               int64  `json:"id"`
	EntidadeId       int64  `json:"entidadeId"`
	EntidadeNome     string `json:"entidadeNome"`
	CicloId          int64  `json:"cicloId"`
	CicloNome        string `json:"cicloNome"`
	CicloNota        string `json:"cicloNota"`
	PilarId          int64  `json:"pilarId"`
	PilarNome        string `json:"pilarNome"`
	PilarPeso        string `json:"pilarPeso"`
	PilarNota        string `json:"pilarNota"`
	ComponenteId     int64  `json:"componenteId"`
	ComponenteNome   string `json:"componenteNome"`
	ComponentePeso   string `json:"componentePeso"`
	ComponenteNota   string `json:"componenteNota"`
	ElementoId       int64  `json:"elementoId"`
	ElementoNome     string `json:"elementoNome"`
	ElementoPeso     string `json:"elementoPeso"`
	ElementoNota     string `json:"elementoNota"`
	TipoPontuacaoId  string `json:"tipoPontuacaoId"`
	TipoNotaId       int    `json:"tipoNotaId"`
	TipoNotaNome     string `json:"tipoNotaNome"`
	TipoNotaLetra    string `json:"tipoNotaLetra"`
	TipoNotaCorLetra string `json:"tipoNotaCorLetra"`
	TipoNotaPeso     string `json:"tipoNotaPeso"`
	TipoNotaNota     string `json:"tipoNotaNota"`
	PesoPadraoEC     string `json:"pesoPadraoEC"`
	TipoMediaCPId    int    `json:"tipoMediaCPId"`
	TipoMediaCP      string `json:"tipoMediaCP"`
	PesoPadraoCP     string `json:"pesoPadraoCP"`
	TipoMediaPCId    int    `json:"tipoMediaPCId"`
	TipoMediaPC      string `json:"tipoMediaPC"`
	PesoPadraoPC     string `json:"pesoPadraoPC"`
	TipoMediaCEId    int    `json:"tipoMediaCEId"`
	TipoMediaCE      string `json:"tipoMediaCE"`
	IniciaEm         string `json:"iniciaEm"`
	TerminaEm        string `json:"terminaEm"`
	ItemId           int64  `json:"itemId"`
	ItemNome         string `json:"itemNome"`
	AuditorId        int64  `json:"auditorId"`
	AuditorName      string `json:"auditorName"`
	SupervisorId     int64  `json:"supervisorId"`
	SupervisorName   string `json:"supervisorName"`
	AuthorId         int64  `json:"autorId"`
	AuthorName       string `json:"autorNome"`
	CriadoEm         string `json:"criadoEm"`
	IdVersaoOrigem   int64  `json:"idVersaoOrigem"`
	StatusId         int64  `json:"statusId"`
	CStatus          string `json:"cStatus"`
}

type ElementoDaMatriz struct {
	Order                   int
	Id                      int64  `json:"id"`
	EntidadeId              int64  `json:"entidadeId"`
	EntidadeNome            string `json:"entidadeNome"`
	CicloId                 int64  `json:"cicloId"`
	CicloNome               string `json:"cicloNome"`
	CicloNota               string `json:"cicloNota"`
	CicloQtdPilares         string `json:"cicloQtdPilares"`
	CicloColSpan            int    `json:"cicloColSpan"`
	PilarId                 int64  `json:"pilarId"`
	PilarNome               string `json:"pilarNome"`
	PilarPeso               string `json:"pilarPeso"`
	PilarNota               string `json:"pilarNota"`
	PilarColSpan            int    `json:"pilarColSpan"`
	PilarQtdComponentes     string `json:"pilarQtdComponentes"`
	ComponenteId            int64  `json:"componenteId"`
	ComponenteNome          string `json:"componenteNome"`
	ComponentePeso          string `json:"componentePeso"`
	ComponenteNota          string `json:"componenteNota"`
	ComponenteColSpan       int    `json:"componenteColSpan"`
	ComponenteQtdTiposNotas string `json:"componenteQtdTiposNotas"`
	ElementoId              int64  `json:"elementoId"`
	ElementoNome            string `json:"elementoNome"`
	ElementoPeso            string `json:"elementoPeso"`
	ElementoNota            string `json:"elementoNota"`
	TipoNotaId              int    `json:"tipoNotaId"`
	TipoNotaNome            string `json:"tipoNotaNome"`
	TipoNotaLetra           string `json:"tipoNotaLetra"`
	TipoNotaCorLetra        string `json:"tipoNotaCorLetra"`
	TipoNotaPeso            string `json:"tipoNotaPeso"`
	TipoNotaNota            string `json:"tipoNotaNota"`
	PesoPadraoEC            string `json:"pesoPadraoEC"`
	TipoMediaCPId           int    `json:"tipoMediaCPId"`
	TipoMediaCP             string `json:"tipoMediaCP"`
	PesoPadraoCP            string `json:"pesoPadraoCP"`
	TipoMediaPCId           int    `json:"tipoMediaPCId"`
	TipoMediaPC             string `json:"tipoMediaPC"`
	PesoPadraoPC            string `json:"pesoPadraoPC"`
	TipoMediaCEId           int    `json:"tipoMediaCEId"`
	TipoMediaCE             string `json:"tipoMediaCE"`
	IniciaEm                string `json:"iniciaEm"`
	TerminaEm               string `json:"terminaEm"`
	ItemId                  int64  `json:"itemId"`
	ItemNome                string `json:"itemNome"`
	AuditorId               int64  `json:"auditorId"`
	AuditorName             string `json:"auditorName"`
	SupervisorId            int64  `json:"supervisorId"`
	SupervisorName          string `json:"supervisorName"`
	AuthorId                int64  `json:"autorId"`
	AuthorName              string `json:"autorNome"`
	CriadoEm                string `json:"criadoEm"`
	IdVersaoOrigem          int64  `json:"idVersaoOrigem"`
	StatusId                int64  `json:"statusId"`
	CStatus                 string `json:"cStatus"`
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

type TipoNota struct {
	Order          int
	Id             int64  `json:"id"`
	TipoNotaId     string `json:"tipoNotaId"'`
	ComponenteId   string `json:"componenteId"'`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Letra          string `json:"letra"`
	CorLetra       string `json:"corLetra"`
	PesoPadrao     string `json:"pesoPadrao"'`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
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

// PÁGINAS
type PageActions struct {
	ErrMsg     string
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageCiclos struct {
	ErrMsg     string
	AppName    string
	Title      string
	Ciclos     []Ciclo
	Entidades  []Entidade
	Pilares    []Pilar
	Users      []User
	LoggedUser LoggedUser
}

type PageComponentes struct {
	ErrMsg      string
	AppName     string
	Title       string
	TiposNota   []TipoNota
	Componentes []Componente
	Elementos   []Elemento
	Users       []User
	LoggedUser  LoggedUser
}

type PageElementos struct {
	ErrMsg     string
	AppName    string
	Title      string
	Elementos  []Elemento
	Users      []User
	LoggedUser LoggedUser
}

type PageEntidades struct {
	ErrMsg     string
	AppName    string
	Title      string
	Entidades  []Entidade
	Ciclos     []Ciclo
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageFeatures struct {
	ErrMsg     string
	AppName    string
	Title      string
	Features   []Feature
	LoggedUser LoggedUser
}

type PageEscritorios struct {
	ErrMsg      string
	AppName     string
	Title       string
	Escritorios []Escritorio
	Entidades   []Entidade
	Users       []User
	LoggedUser  LoggedUser
}

type PageEntidadesCiclos struct {
	ErrMsg       string
	AppName      string
	Title        string
	Entidades    []Entidade
	Membros      []Membro
	Supervisores []User
	LoggedUser   LoggedUser
}

type PageMatriz struct {
	ErrMsg            string
	AppName           string
	Title             string
	ElementosDaMatriz []ElementoDaMatriz
	Supervisores      []User
	Auditores         []User
	LoggedUser        LoggedUser
}

type PageProdutosItens struct {
	ErrMsg       string
	AppName      string
	Title        string
	Inc          func(i int) int
	Produtos     []ProdutoItem
	Supervisores []User
	Auditores    []User
	LoggedUser   LoggedUser
}

type PageProdutosComponentes struct {
	ErrMsg       string
	AppName      string
	Title        string
	Produtos     []ProdutoComponente
	Supervisores []User
	Auditores    []User
	LoggedUser   LoggedUser
}

type PagePilares struct {
	ErrMsg      string
	AppName     string
	Title       string
	Pilares     []Pilar
	Componentes []Componente
	Users       []User
	LoggedUser  LoggedUser
}

type PagePlanos struct {
	ErrMsg     string
	AppName    string
	Title      string
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageRoles struct {
	ErrMsg     string
	AppName    string
	Title      string
	Roles      []Role
	Features   []Feature
	LoggedUser LoggedUser
}

type PageStatus struct {
	ErrMsg     string
	AppName    string
	Title      string
	Statuss    []Status
	LoggedUser LoggedUser
}

type PageTiposNotas struct {
	ErrMsg     string
	AppName    string
	Title      string
	TiposNotas []TipoNota
	Users      []User
	LoggedUser LoggedUser
}

type PageUsers struct {
	ErrMsg      string
	AppName     string
	Title       string
	Users       []User
	Escritorios []Escritorio
	Roles       []Role
	LoggedUser  LoggedUser
}

type PageWorkflows struct {
	ErrMsg     string
	AppName    string
	Title      string
	Features   []Feature
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
