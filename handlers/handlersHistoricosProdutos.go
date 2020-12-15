package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
)

func registrarHistoricoAuditorComponente(produto mdl.ProdutoComponente, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_componentes_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	componente_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	tipo_alteracao,  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	auditor_anterior_id,  " +
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
		"	'R',  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		strconv.FormatInt(produto.AuditorAnteriorId, 10) + ",  " +
		strconv.FormatInt(currentUser.Id, 10) + ",  " +
		"	$1,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_componentes " +
		"	WHERE entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) + " AND " +
		"	ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) + " AND " +
		"	pilar_id = " + strconv.FormatInt(produto.PilarId, 10) + " AND " +
		"	componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoComponenteId := 0
	err := Db.QueryRow(sqlStatement, time.Now()).Scan(&historicoProdutoComponenteId)
	if err != nil {
		log.Println(err)
	}
}

func registrarHistoricoNotaElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_elementos_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	plano_id,  " +
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
		"	plano_id,  " +
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
		"	" + strconv.FormatInt(currentUser.Id, 10) + ",  " +
		"	now()::timestamp,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_elementos " +
		"	WHERE entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) + " AND " +
		"	ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) + " AND " +
		"	pilar_id = " + strconv.FormatInt(produto.PilarId, 10) + " AND " +
		"	plano_id = " + strconv.FormatInt(produto.PlanoId, 10) + " AND " +
		"	componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) + " AND " +
		"	tipo_nota_id = " + strconv.FormatInt(produto.TipoNotaId, 10) + " AND " +
		"	elemento_id = " + strconv.FormatInt(produto.ElementoId, 10) +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoElementoId := 0
	err := Db.QueryRow(sqlStatement).Scan(&historicoProdutoElementoId)
	if err != nil {
		log.Println(err)
	}
}

func registrarHistoricoPesoElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "INSERT INTO produtos_elementos_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
		"	plano_id,  " +
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
		"	plano_id,  " +
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
		"	" + strconv.FormatInt(currentUser.Id, 10) + ",  " +
		"	now()::timestamp,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_elementos " +
		"	WHERE entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) + " AND " +
		"	ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) + " AND " +
		"	pilar_id = " + strconv.FormatInt(produto.PilarId, 10) + " AND " +
		"	plano_id = " + strconv.FormatInt(produto.PlanoId, 10) + " AND " +
		"	componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) + " AND " +
		"	tipo_nota_id = " + strconv.FormatInt(produto.TipoNotaId, 10) + " AND " +
		"	elemento_id = " + strconv.FormatInt(produto.ElementoId, 10) +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoElementoId := 0
	err := Db.QueryRow(
		sqlStatement).Scan(&historicoProdutoElementoId)
	if err != nil {
		log.Println(err)
	}
}

func LoadHistoricosElemento(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Historicos do Elemento")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var planoId = r.FormValue("planoId")
	var componenteId = r.FormValue("componenteId")
	var elementoId = r.FormValue("elementoId")
	var filtro mdl.Historico
	filtro.EntidadeId = entidadeId
	filtro.CicloId = cicloId
	filtro.PilarId = pilarId
	filtro.PlanoId = planoId
	filtro.ComponenteId = componenteId
	filtro.ElementoId = elementoId
	historicos := ListHistoricosElemento(filtro)
	jsonHistoricos, _ := json.Marshal(historicos)
	w.Write([]byte(jsonHistoricos))
	log.Println("JSON Históricos do Elemento")
}

func ListHistoricosElemento(filtro mdl.Historico) []mdl.Historico {
	log.Println("List Históricos do Elemento")
	sql := "SELECT " +
		"a.id, " +
		"a.entidade_id, " +
		"a.ciclo_id, " +
		"a.pilar_id, " +
		"a.plano_id, " +
		"a.componente_id, " +
		"a.elemento_id, " +
		"a.peso, " +
		"a.tipo_pontuacao_id, " +
		"a.nota, " +
		"a.author_id, " +
		"coalesce(b.name,''), " +
		"coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS')) as alterado_em, " +
		"case when tipo_alteracao = 'P' then a.motivacao_peso else a.motivacao_nota end, " +
		"case when tipo_alteracao = 'P' then 'Peso' else 'Nota' end " +
		"FROM produtos_elementos_historicos a " +
		"LEFT JOIN users b ON a.author_id = b.id " +
		"WHERE a.entidade_id = " + filtro.EntidadeId + " AND " +
		"a.ciclo_id = " + filtro.CicloId + " AND " +
		"a.pilar_id = " + filtro.PilarId + " AND " +
		"a.plano_id = " + filtro.PlanoId + " AND " +
		"a.componente_id = " + filtro.ComponenteId + " AND " +
		"a.elemento_id = " + filtro.ElementoId + " ORDER BY a.criado_em DESC "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var historicos []mdl.Historico
	var historico mdl.Historico
	for rows.Next() {
		rows.Scan(
			&historico.Id,
			&historico.EntidadeId,
			&historico.CicloId,
			&historico.PilarId,
			&historico.PlanoId,
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

func LoadHistoricosComponente(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Historicos do Componente")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var componenteId = r.FormValue("componenteId")
	var filtro mdl.Historico
	filtro.EntidadeId = entidadeId
	filtro.CicloId = cicloId
	filtro.PilarId = pilarId
	filtro.ComponenteId = componenteId
	historicos := ListHistoricosComponente(filtro)
	jsonHistoricos, _ := json.Marshal(historicos)
	w.Write([]byte(jsonHistoricos))
	log.Println("JSON Históricos do Componente")
}

func ListHistoricosComponente(filtro mdl.Historico) []mdl.Historico {
	log.Println("List Históricos do Componente")
	sql :=
		"SELECT  " +
			"	a.id,  " +
			"	entidade_id,  " +
			"	ciclo_id,  " +
			"	pilar_id,  " +
			"	componente_id,  " +
			"	coalesce(peso,0),  " +
			"	tipo_pontuacao_id,  " +
			"	coalesce(nota,0),  " +
			"	tipo_alteracao,  " +
			"	supervisor_id,  " +
			"	auditor_id,  " +
			"	coalesce(auditor_anterior_id,0),  " +
			"	a.author_id,  " +
			"	coalesce(b.name,'') as author_name, " +
			"	coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS')) as alterado_em,  " +
			"	case " +
			" 		when tipo_alteracao = 'R' then justificativa " +
			"       when tipo_alteracao = 'I' then motivacao_cronograma " +
			"       when tipo_alteracao = 'T' then motivacao_cronograma " +
			"       when tipo_alteracao = 'P' then motivacao_config " +
			"	end, " +
			"	case " +
			" 		when tipo_alteracao = 'R' then 'Remoção' " +
			"       when tipo_alteracao = 'I' then 'Inicia Em' " +
			"       when tipo_alteracao = 'T' then 'Termina Em' " +
			"       when tipo_alteracao = 'P' then 'Planos' " +
			"	end " +
			"	FROM produtos_componentes_historicos a " +
			"	LEFT JOIN users b ON a.author_id = b.id " +
			"	WHERE a.entidade_id = " + filtro.EntidadeId + " AND " +
			"   a.ciclo_id = " + filtro.CicloId + " AND " +
			"	a.pilar_id = " + filtro.PilarId + " AND " +
			"	a.componente_id = " + filtro.ComponenteId +
			" 	ORDER BY a.criado_em DESC "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var historicos []mdl.Historico
	var historico mdl.Historico
	for rows.Next() {
		rows.Scan(
			&historico.Id,
			&historico.EntidadeId,
			&historico.CicloId,
			&historico.PilarId,
			&historico.ComponenteId,
			&historico.Peso,
			&historico.Metodo,
			&historico.Nota,
			&historico.TipoAlteracao,
			&historico.SupervisorId,
			&historico.AuditorNovoId,
			&historico.AuditorAnteriorId,
			&historico.AutorId,
			&historico.AutorNome,
			&historico.AlteradoEm,
			&historico.Motivacao)
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

func registrarHistoricoPesoPilar(produto mdl.ProdutoPilar, currentUser mdl.User) {
	log.Println("========== registrarHistoricoPesoPilar ===========")
	sqlStatement := "INSERT INTO produtos_pilares_historicos( " +
		"	entidade_id,  " +
		"	ciclo_id,  " +
		"	pilar_id,  " +
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
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	'P',  " +
		"	motivacao_peso,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	" + strconv.FormatInt(currentUser.Id, 10) + ",  " +
		"	now()::timestamp,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_pilares " +
		"	WHERE entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) + " AND " +
		"	ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) + " AND " +
		"	pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		" RETURNING id"
	log.Println(sqlStatement)
	historicoProdutoPilarId := 0
	err := Db.QueryRow(
		sqlStatement).Scan(&historicoProdutoPilarId)
	if err != nil {
		log.Println(err)
	}
}

func LoadHistoricosPilar(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Historicos do Pilar")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var filtro mdl.Historico
	filtro.EntidadeId = entidadeId
	filtro.CicloId = cicloId
	filtro.PilarId = pilarId
	historicos := ListHistoricosPilar(filtro)
	jsonHistoricos, _ := json.Marshal(historicos)
	w.Write([]byte(jsonHistoricos))
	log.Println("JSON Históricos do Pilar")
}

func ListHistoricosPilar(filtro mdl.Historico) []mdl.Historico {
	log.Println("List Históricos do Pilar")
	sql :=
		"SELECT  " +
			"	a.id,  " +
			"	entidade_id,  " +
			"	ciclo_id,  " +
			"	pilar_id,  " +
			"	coalesce(peso,0),  " +
			"	tipo_pontuacao_id,  " +
			"	coalesce(nota,0),  " +
			"	tipo_alteracao,  " +
			"	a.author_id,  " +
			"	coalesce(b.name,'') as author_name, " +
			"	coalesce(to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS')) as alterado_em,  " +
			"	motivacao_peso  " +
			"	FROM produtos_pilares_historicos a " +
			"	LEFT JOIN users b ON a.author_id = b.id " +
			"	WHERE a.entidade_id = " + filtro.EntidadeId + " AND " +
			"   a.ciclo_id = " + filtro.CicloId + " AND " +
			"	a.pilar_id = " + filtro.PilarId +
			" 	ORDER BY a.criado_em DESC "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var historicos []mdl.Historico
	var historico mdl.Historico
	for rows.Next() {
		rows.Scan(
			&historico.Id,
			&historico.EntidadeId,
			&historico.CicloId,
			&historico.PilarId,
			&historico.Peso,
			&historico.Metodo,
			&historico.Nota,
			&historico.TipoAlteracao,
			&historico.AutorId,
			&historico.AutorNome,
			&historico.AlteradoEm,
			&historico.Motivacao)
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
