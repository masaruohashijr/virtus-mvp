package db

import (
//"log"
)

func createTiposNotas() {
	stmtTiposNotas := " INSERT INTO tipos_notas ( " +
		" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
		" SELECT 'Nota de Risco', 'Descrição da Nota de Risco', 'R', 'FF6961', 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'R')"
	//log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
	stmtTiposNotas = " INSERT INTO tipos_notas ( " +
		" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
		" SELECT 'Nota de Controle', 'Descrição da Nota de Controle', 'C', 'FDF6C6', 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'C')"
	//log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
}
