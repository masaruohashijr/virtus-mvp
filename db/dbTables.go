package db

import ()

func createTable() {
	// Table ACTIONS
	db.Exec(" CREATE TABLE IF NOT EXISTS public.actions (" +
		" id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL, " +
		" name character varying(255) NOT NULL, " +
		" origin_status_id integer, " +
		" destination_status_id integer, " +
		" other_than boolean, " +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)")
	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")

	// Table ACTIVITIES
	db.Exec(" CREATE TABLE public.activities (" +
		" id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass)," +
		" workflow_id integer," +
		" action_id integer," +
		" expiration_action_id integer," +
		" expiration_time_days integer," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone)")

	// Table ACTIVITIES_ROLES
	db.Exec(
		" CREATE TABLE public.activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activity_id integer," +
			" role_id integer)")

	// Table CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.ciclos (" +
			" id integer DEFAULT nextval('public.ciclos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table CICLOS_ENTIDADES
	db.Exec(
		" CREATE TABLE public.ciclos_entidades (" +
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
		" CREATE TABLE IF NOT EXISTS public.componentes (" +
			" id integer DEFAULT nextval('public.componentes_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table COMPONENTES_PILARES
	db.Exec(
		" CREATE TABLE public.componentes_pilares (" +
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
		" CREATE TABLE IF NOT EXISTS public.elementos (" +
			" id integer DEFAULT nextval('public.elementos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" peso integer DEFAULT 1 NOT NULL," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" status_id integer)")

	// Table ELEMENTOS_COMPONENTES
	db.Exec(
		" CREATE TABLE public.elementos_componentes (" +
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
		" CREATE TABLE IF NOT EXISTS public.entidades (" +
			" id integer DEFAULT nextval('public.entidades_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.escritorios (" +
			" id integer DEFAULT nextval('public.escritorios_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.features  (" +
			" id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE public.features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")

	// Table FEATURES_ACTIVITIES
	db.Exec(
		" CREATE TABLE public.features_activities (" +
			" id integer DEFAULT nextval('features_activities_id_seq'::regclass)," +
			" feature_id integer," +
			" activity_id integer)")

	// Table INTEGRANTES
	db.Exec(
		" CREATE TABLE public.integrantes (" +
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
		" CREATE TABLE IF NOT EXISTS public.itens (" +
			" id integer DEFAULT nextval('public.itens_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" criado_em timestamp without time zone," +
			" author_id integer," +
			" status_id integer)")

	// Table JURISDICOES
	db.Exec(
		" CREATE TABLE public.jurisdicoes (" +
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
		" CREATE TABLE public.membros (" +
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
		" CREATE TABLE IF NOT EXISTS public.pilares (" +
			" id integer DEFAULT nextval('public.pilares_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PILARES_CICLOS
	db.Exec(
		" CREATE TABLE public.pilares_ciclos (" +
			" id integer DEFAULT nextval('pilares_ciclos_id_seq'::regclass)," +
			" pilar_id integer," +
			" ciclo_id integer," +
			" tipo_media integer," +
			" peso_padrao integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PLANOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.planos (" +
			" id integer DEFAULT nextval('public.planos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PRODUTOS_CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.produtos_ciclos (" +
			" id integer DEFAULT nextval('public.produtos_ciclos_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.produtos_pilares (" +
			" id integer DEFAULT nextval('public.produtos_pilares_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.produtos_componentes (" +
			" id integer DEFAULT nextval('public.produtos_componentes_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.produtos_elementos (" +
			" id integer DEFAULT nextval('public.produtos_elementos_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.produtos_itens (" +
			" id integer DEFAULT nextval('public.produtos_itens_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.roles  (" +
			" id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.status  (" +
			" id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer," +
			" stereotype character varying(255))")
	// Table TIPOS DE NOTAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.tipos_notas (" +
			" id integer DEFAULT nextval('public.tipos_notas_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(255)," +
			" letra character(1) NOT NULL," +
			" cor_letra character(6)," +
			" dominio_componente bool," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")
	// Table USERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.users (" +
			" id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL," +
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
		" CREATE TABLE IF NOT EXISTS public.workflows  (" +
			" id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL," +
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
