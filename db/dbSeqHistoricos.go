package db

import ()

func createSeqHistoricos() {
	// Sequence PRODUTOS_CICLOS_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_ciclos_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_PILARES_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_pilares_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_COMPONENTES_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_componentes_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_TIPOS_NOTAS_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_tipos_notas_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_ELEMENTOS_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_elementos_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PRODUTOS_ITENS_HISTORICOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_itens_historicos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
}
