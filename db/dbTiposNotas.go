package db

import (
	"log"
)

func createTiposNotas() {
	stmtTiposNotas := " INSERT INTO public.tipos_notas_componentes( " +
		" nome, descricao, letra, cor_letra, dominio_componente, author_id, criado_em, status_id) " +
		" SELECT 'Nota de Risco', 'Descrição da Nota de Risco', 'R', 'C90000', true, 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'R')"
	log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
	stmtTiposNotas = " INSERT INTO public.tipos_notas_componentes( " +
		" nome, descricao, letra, cor_letra, dominio_componente, author_id, criado_em, status_id) " +
		" SELECT 'Nota de Controle', 'Descrição da Nota de Controle', 'C', '0000C9', true, 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'C')"
	log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
}
