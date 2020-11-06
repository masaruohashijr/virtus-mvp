package db

import (
	"log"
)

func createRoles() {
	query := " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 1, 'Admin', 'Admin' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Admin')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 2, 'Chefe', 'Chefe' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Chefe')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 3, 'Supervisor', 'Supervisor' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Supervisor')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 4, 'Auditor', 'Auditor' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Auditor')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 5, 'Visualizador', 'Visualizador' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Visualizador')"
	db.Exec(query)
}

func updateRoles() {
	query := " UPDATE roles SET author_id = 1 WHERE name = 'Admin' AND (SELECT author_id FROM roles WHERE name = 'Admin') IS NULL "
	log.Println(query)
	db.Exec(query)
}
