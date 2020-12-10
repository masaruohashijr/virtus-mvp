package db

import (
	"database/sql"
	//"log"
	hd "virtus/handlers"
)

var db *sql.DB

func Initialize() {
	db = hd.Db
	createSeq()
	createSeqHistoricos()
	createTable()
	createTablesHistoricos()
	createFeatures()
	createRoles()
	createRoleFeatures()
	createEscritorios()
	createUsers()
	createMembros()
	createStatusZERO()
	createEntidades()
	createPlanos()
	createJurisdicoes()
	updateRoles()
	updateFeatures()
	createPKey()
	createFKey()
	createUniqueKey()
	createCicloCompleto()
}

func createStatusZERO() {
	query := "INSERT INTO status (id, name, stereotype, description, author_id, created_at)" +
		" SELECT 0, '-', '', '', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM status WHERE id = 0)"
	//log.Println(query)
	db.Exec(query)
}
