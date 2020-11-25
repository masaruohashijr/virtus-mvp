package db

import (
//"log"
)

func createTiposNotas() {
	stmtTiposNotas := " INSERT INTO tipos_notas ( " +
		" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
		" SELECT 'Risco', 'Descrição da Nota de Risco', 'R', 'ED7864', 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'R')"
	//log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
	stmtTiposNotas = " INSERT INTO tipos_notas ( " +
		" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
		" SELECT 'Controle', 'Descrição da Nota de Controle', 'C', 'EDBC64', 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'C')"
	//log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
	stmtTiposNotas = " INSERT INTO tipos_notas ( " +
		" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
		" SELECT 'Avaliação', 'Descrição da Avaliação', 'A', '6495ED', 1, now()::timestamp, 0 " +
		" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = 'A')"
	//log.Println(stmtTiposNotas)
	db.Exec(stmtTiposNotas)
}
