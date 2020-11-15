package handlers

import (
	"log"
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
		"	motivacao_peso,  " +
		"	motivacao_nota,  " +
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
		"	motivacao_peso,  " +
		"	motivacao_nota,  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	$1,  " +
		"	$2,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_componentes " +
		"	FROM entidade_id = $3, " +
		"	ciclo_id = $4, " +
		"	pilar_id = $5, " +
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
		"	motivacao_peso,  " +
		"	motivacao_nota,  " +
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
		"	tipo_nota_id,  " +
		"	elemento_id,  " +
		"	tipo_pontuacao_id,  " +
		"	peso,  " +
		"	nota,  " +
		"	motivacao_peso,  " +
		"	motivacao_nota,  " +
		"	justificativa,  " +
		"	supervisor_id,  " +
		"	auditor_id,  " +
		"	$1,  " +
		"	$2,  " +
		"	id,  " +
		"	status_id " +
		"	FROM produtos_elementos " +
		"	FROM entidade_id = $3, " +
		"	ciclo_id = $4, " +
		"	pilar_id = $5, " +
		"	componente_id = $6 " +
		"	tipo_nota_id = $7 " +
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
