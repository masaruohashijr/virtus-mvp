package db

import (
//"log"
)

func createTablesHistoricos() {
	// Table PRODUTOS_CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_ciclos_historicos (" +
			" id integer DEFAULT nextval('produtos_ciclos_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" tipo_pontuacao_id integer," +
			" nota double precision," +
			" motivacao character varying(4000)," +
			" supervisor_id integer," +
			" auditor_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_PILARES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_pilares_historicos (" +
			" id integer DEFAULT nextval('produtos_pilares_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" tipo_pontuacao_id integer," +
			" peso double precision," +
			" nota double precision," +
			" motivacao_peso character varying(4000)," +
			" motivacao_nota character varying(4000)," +
			" supervisor_id integer," +
			" auditor_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_COMPONENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_componentes_historicos (" +
			" id integer DEFAULT nextval('produtos_componentes_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" tipo_pontuacao_id integer," +
			" peso double precision," +
			" nota double precision," +
			" motivacao_peso character varying(4000)," +
			" motivacao_nota character varying(4000)," +
			" justificativa character varying(4000)," +
			" supervisor_id integer," +
			" auditor_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_ELEMENTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_elementos_historicos (" +
			" id integer DEFAULT nextval('produtos_elementos_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" tipo_nota_id integer," +
			" elemento_id integer," +
			" tipo_pontuacao_id integer," +
			" peso double precision," +
			" nota double precision," +
			" motivacao_peso character varying(4000)," +
			" motivacao_nota character varying(4000)," +
			" justificativa character varying(4000)," +
			" supervisor_id integer," +
			" auditor_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_ITENS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_itens_historicos (" +
			" id integer DEFAULT nextval('produtos_itens_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" tipo_nota_id integer," +
			" elemento_id integer," +
			" item_id integer," +
			" avaliacao character varying(4000)," +
			" anexo character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_TIPOS_NOTAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_tipos_notas_historicos (" +
			" id integer DEFAULT nextval('produtos_tipos_notas_historicos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" tipo_nota_id integer," +
			" tipo_pontuacao_id integer," +
			" peso double precision," +
			" nota double precision," +
			" anexo character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")
}
