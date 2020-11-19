package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	mdl "virtus/models"
)

func registrarHistoricoAuditorComponente(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_componentes_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	author_id,  " +
		"	criado_em,  " +
		"	id_versao_origem,  " +
		"	status_id) " +
		"SELECT  " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	$1,  " +
		"	$2,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_componentes " +
		"	WHERE entidade_id = $3 AND " +
		"	ciclo_id = $4 AND " +
		"	pilar_id = $5 AND " +
		"	componente_id = $6 " +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoComponenteId := 0
	err := Db.QueryRow(
		sqlStatement,
		currentUser.Id,
		time.Now(),
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId).Scan(&historicoProdutoComponenteId)
	if err != nil {
		log.Println(err)
	}
}

func registrarHistoricoNotaElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_elementos_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"   tipo_nota_id," +
		"   elemento_id," +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	tipo_alteracao,  " +
		"	motivacao_nota,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	author_id,  " +
		"	criado_em,  " +
		"	id_versao_origem,  " +
		"	status_id) " +
		"SELECT  " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"	tipo_nota_id,  " +
		"	elemento_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	'N',  " +
		"	motivacao_nota,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	$1,  " +
		"	$2,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_elementos " +
		"	WHERE entidade_id = $3 AND " +
		"	ciclo_id = $4 AND " +
		"	pilar_id = $5 AND " +
		"	componente_id = $6 AND " +
		"	tipo_nota_id = $7 AND " +
		"	elemento_id = $8 " +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoComponenteId := 0
	err := Db.QueryRow(
		sqlStatement,
		currentUser.Id,
		time.Now(),
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.TipoNotaId,
		produto.ElementoId).Scan(&historicoProdutoComponenteId)
	if err != nil {
		log.Println(err)
	}
}

func registrarHistoricoPesoElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_elementos_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"   tipo_nota_id," +
		"   elemento_id," +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	tipo_alteracao,  " +
		"	motivacao_peso,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	author_id,  " +
		"	criado_em,  " +
		"	id_versao_origem,  " +
		"	status_id) " +
		"SELECT  " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"	tipo_nota_id,  " +
		"	elemento_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	'P',  " +
		"	motivacao_peso,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	$1,  " +
		"	$2,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_elementos " +
		"	WHERE entidade_id = $3 AND " +
		"	ciclo_id = $4 AND " +
		"	pilar_id = $5 AND " +
		"	componente_id = $6 AND " +
		"	tipo_nota_id = $7 AND " +
		"	elemento_id = $8" +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoComponenteId := 0
	err := Db.QueryRow(
		sqlStatement,
		currentUser.Id,
		time.Now(),
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.TipoNotaId,
		produto.ElementoId).Scan(&historicoProdutoComponenteId)
	if err != nil {
		log.Println(err)
	}
}

func LoadHistoricos(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Historicos")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var componenteId = r.FormValue("componenteId")
	var elementoId = r.FormValue("elementoId")
	var filtro mdl.Historico
	filtro.EntidadeId = entidadeId
	filtro.CicloId = cicloId
	filtro.PilarId = pilarId
	filtro.ComponenteId = componenteId
	filtro.ElementoId = elementoId
	historicos := ListHistoricos(filtro)
	jsonHistoricos, _ := json.Marshal(historicos)
	w.Write([]byte(jsonHistoricos))
	log.Println("JSON Históricos")
}

func ListHistoricos(filtro mdl.Historico) []mdl.Historico {
	log.Println("List Históricos")
	sql := "SELECT " +
		"a.id, " +
		"a.entidade_id, " +
		"a.ciclo_id, " +
		"a.pilar_id, " +
		"a.componente_id, " +
		"a.elemento_id, " +
		"a.peso, " +
		"a.tipo_pontuacao_id, " +
		"a.nota, " +
		"a.author_id, " +
		"coalesce(b.name,''), " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MM:SS')) as alterado_em, " +
		"case when tipo_alteracao = 'P' then a.motivacao_peso else a.motivacao_nota end, " +
		"case when tipo_alteracao = 'P' then 'Peso' else 'Nota' end " +
		"FROM produtos_elementos_historicos a " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"WHERE a.entidade_id = " + filtro.EntidadeId + " AND " +
		"a.ciclo_id = " + filtro.CicloId + " AND " +
		"a.pilar_id = " + filtro.PilarId + " AND " +
		"a.componente_id = " + filtro.ComponenteId + " AND " +
		"a.elemento_id = " + filtro.ElementoId
	log.Println(sql)
	rows, _ := Db.Query(sql)
	var historicos []mdl.Historico
	var historico mdl.Historico
	for rows.Next() {
		rows.Scan(
			&historico.Id,
			&historico.EntidadeId,
			&historico.CicloId,
			&historico.PilarId,
			&historico.ComponenteId,
			&historico.ElementoId,
			&historico.Peso,
			&historico.Metodo,
			&historico.Nota,
			&historico.AutorId,
			&historico.AutorNome,
			&historico.AlteradoEm,
			&historico.Motivacao,
			&historico.TipoAlteracao)
		switch historico.Metodo {
		case "1":
			historico.Metodo = "Manual"
		case "2":
			historico.Metodo = "Calculada"
		case "3":
			historico.Metodo = "Ajustada"
		}
		historicos = append(historicos, historico)
		log.Println(historico)
	}
	return historicos
}
