package db

import ()

func createSeq() {
	// Sequence ACTIONS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIVITIES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_roles_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.ciclos_entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence COMPONENTES_PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.componentes_pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ELEMENTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.elementos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ELEMENTOS_COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.elementos_componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ESCRITORIOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.escritorios_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_id_seq " +
		" START WITH 18" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ACTIVITIES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_roles_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ITENS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.itens_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence JURISDICOES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.jurisdicoes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence MEMBROS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.membros_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PLANOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.planos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.pilares_ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.roles_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence TIPOS_NOTAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.tipos_notas_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.users_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence WORKFLOWS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.workflows_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
}
