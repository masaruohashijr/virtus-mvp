package db

import ()

func createSeq() {
	// Sequence ACTIONS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS actions_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS actions_status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIVITIES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS activities_roles_id_seq " +
		" START WITH 6" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CHAMADOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS chamados_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS ciclos_entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence COMPONENTES_PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS componentes_pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ELEMENTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS elementos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ELEMENTOS_COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS elementos_componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ESCRITORIOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS escritorios_id_seq " +
		" START WITH 6" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS features_id_seq " +
		" START WITH 50" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ACTIVITIES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS features_activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS features_roles_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence INTEGRANTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS integrantes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ITENS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS itens_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence JURISDICOES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS jurisdicoes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence MEMBROS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS membros_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PLANOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS planos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS pilares_ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_PLANOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_planos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_TIPOS_NOTAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_tipos_notas_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_ELEMENTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_elementos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_ITENS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS produtos_itens_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ANOTACOES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS anotacoes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ANOTACOES_RADARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS anotacoes_radares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence RADARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS radares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS roles_id_seq " +
		" START WITH 6" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence VERSOES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS versoes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence TIPOS_NOTAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS tipos_notas_id_seq " +
		" START WITH 4" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence TIPOS_NOTAS_COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS tipos_notas_componentes_id_seq " +
		" START WITH 16" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS users_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence WORKFLOWS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS workflows_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
}
