package db

import (
	"log"
)

func createTable() {
	// Table ACTIONS
	stmt := " CREATE TABLE IF NOT EXISTS actions (" +
		" id integer DEFAULT nextval('actions_id_seq'::regclass) NOT NULL, " +
		" name character varying(255) NOT NULL, " +
		" origin_status_id integer, " +
		" destination_status_id integer, " +
		" other_than boolean, " +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}
	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")

	// Table ACTIVITIES
	db.Exec(" CREATE TABLE activities (" +
		" id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass)," +
		" workflow_id integer," +
		" action_id integer," +
		" expiration_action_id integer," +
		" expiration_time_days integer," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone)")

	// Table ACTIVITIES_ROLES
	db.Exec(
		" CREATE TABLE activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activity_id integer," +
			" role_id integer)")

	// Table CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS ciclos (" +
			" id integer DEFAULT nextval('ciclos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table CICLOS_ENTIDADES
	db.Exec(
		" CREATE TABLE ciclos_entidades (" +
			" id integer DEFAULT nextval('ciclos_entidades_id_seq'::regclass)," +
			" ciclo_id integer," +
			" entidade_id integer," +
			" tipo_media integer," +
			" supervisor_id integer," + // TODO FK
			" inicia_em timestamp without time zone," +
			" termina_em timestamp without time zone," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table COMPONENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS componentes (" +
			" id integer DEFAULT nextval('componentes_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table COMPONENTES_PILARES
	db.Exec(
		" CREATE TABLE componentes_pilares (" +
			" id integer DEFAULT nextval('componentes_pilares_id_seq'::regclass)," +
			" componente_id integer," +
			" pilar_id integer," +
			" tipo_media integer," +
			" peso_padrao integer," +
			" sonda character varying (255)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ELEMENTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS elementos (" +
			" id integer DEFAULT nextval('elementos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" peso integer DEFAULT 1 NOT NULL," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" status_id integer)")

	// Table ELEMENTOS_COMPONENTES
	db.Exec(
		" CREATE TABLE elementos_componentes (" +
			" id integer DEFAULT nextval('elementos_componentes_id_seq'::regclass)," +
			" componente_id integer," +
			" elemento_id integer," +
			" tipo_nota_id integer," +
			" peso_padrao integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ENTIDADES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS entidades (" +
			" id integer DEFAULT nextval('entidades_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" sigla character varying(25)," +
			" codigo character varying(4000)," +
			" situacao character varying(30)," +
			" ESI bool," +
			" municipio character varying(255)," +
			" sigla_uf character(2)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ESCRITORIOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS escritorios (" +
			" id integer DEFAULT nextval('escritorios_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" abreviatura character (4)," +
			" descricao character varying(4000)," +
			" chefe_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table FEATURES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS features  (" +
			" id integer DEFAULT nextval('features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")

	// Table FEATURES_ACTIVITIES
	db.Exec(
		" CREATE TABLE features_activities (" +
			" id integer DEFAULT nextval('features_activities_id_seq'::regclass)," +
			" feature_id integer," +
			" activity_id integer)")

	// Table INTEGRANTES
	db.Exec(
		" CREATE TABLE integrantes (" +
			" id integer DEFAULT nextval('integrantes_id_seq'::regclass)," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" usuario_id integer," +
			" inicia_em timestamp without time zone," +
			" termina_em timestamp without time zone," +
			" motivacao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ITENS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS itens (" +
			" id integer DEFAULT nextval('itens_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" criado_em timestamp without time zone," +
			" author_id integer," +
			" status_id integer)")

	// Table JURISDICOES
	db.Exec(
		" CREATE TABLE jurisdicoes (" +
			" id integer DEFAULT nextval('jurisdicoes_id_seq'::regclass)," +
			" escritorio_id integer," +
			" entidade_id integer," +
			" inicia_em timestamp without time zone," +
			" termina_em timestamp without time zone," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table MEMBROS
	db.Exec(
		" CREATE TABLE membros (" +
			" id integer DEFAULT nextval('membros_id_seq'::regclass)," +
			" escritorio_id integer," +
			" usuario_id integer," +
			" inicia_em timestamp without time zone," +
			" termina_em timestamp without time zone," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PILARES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS pilares (" +
			" id integer DEFAULT nextval('pilares_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PILARES_CICLOS
	db.Exec(
		" CREATE TABLE pilares_ciclos (" +
			" id integer DEFAULT nextval('pilares_ciclos_id_seq'::regclass)," +
			" pilar_id integer," +
			" ciclo_id integer," +
			" tipo_media integer," +
			" peso_padrao double precision," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PLANOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS planos (" +
			" id integer DEFAULT nextval('planos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS produtos_ciclos (" +
			" id integer DEFAULT nextval('produtos_ciclos_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS produtos_pilares (" +
			" id integer DEFAULT nextval('produtos_pilares_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS produtos_componentes (" +
			" id integer DEFAULT nextval('produtos_componentes_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS produtos_elementos (" +
			" id integer DEFAULT nextval('produtos_elementos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" elemento_id integer," +
			" tipo_pontuacao_id integer," +
			" tipo_nota_id integer," +
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
		" CREATE TABLE IF NOT EXISTS produtos_itens (" +
			" id integer DEFAULT nextval('produtos_itens_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" ciclo_id integer," +
			" pilar_id integer," +
			" componente_id integer," +
			" elemento_id integer," +
			" item_id integer," +
			" avaliacao character varying(4000)," +
			" anexo character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")
	// Table ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS roles  (" +
			" id integer DEFAULT nextval('roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS status  (" +
			" id integer DEFAULT nextval('status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer," +
			" stereotype character varying(255))")
	// Table TIPOS DE NOTAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS tipos_notas (" +
			" id integer DEFAULT nextval('tipos_notas_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(255)," +
			" letra character(1) NOT NULL," +
			" cor_letra character(6)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table TIPOS_NOTAS_COMPONENTES
	db.Exec(
		" CREATE TABLE tipos_notas_componentes (" +
			" id integer DEFAULT nextval('tipos_notas_componentes_id_seq'::regclass)," +
			" componente_id integer," +
			" tipo_nota_id integer," +
			" peso_padrao integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table USERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS users (" +
			" id integer DEFAULT nextval('users_id_seq'::regclass) NOT NULL," +
			" name character varying(255)," +
			" username character varying(255) NOT NULL," +
			" password character varying(255) NOT NULL," +
			" email character varying(255)," +
			" mobile character varying(255)," +
			" role_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table WORKFLOWS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS workflows  (" +
			" id integer DEFAULT nextval('workflows_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" entity_type character varying(50)," +
			" start_at timestamp without time zone," +
			" end_at timestamp without time zone," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")
}
