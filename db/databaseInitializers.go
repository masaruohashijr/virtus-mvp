package db

import (
	"database/sql"
	"log"
	hd "virtus/handlers"
)

var db *sql.DB

func Initialize() {
	db = hd.Db
	createSeq()
	createTable()
	createFeatures()
	createRoles()
	createRoleFeatures()
	createUsers()
	createEscritorios()
	createStatusZERO()
	createEntidades()
	updateRoles()
	updateFeatures()
	createPKey()
	createFKey()
	createUniqueKey()
	createCicloCompleto(1, 3, 3, 3, 3)
}

func createStatusZERO() {
	query := "INSERT INTO status (id, name, stereotype, description, author_id, created_at)" +
		" SELECT 0, '-', '', '', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM status WHERE id = 0)"
	log.Println(query)
	db.Exec(query)
}
