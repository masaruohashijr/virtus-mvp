package handlers

import (
	"log"
	"net/http"
	"strconv"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreatePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		nome := r.FormValue("Nome")
		sqlStatement := "INSERT INTO planos(nome) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, nome).Scan(&id)
		log.Println(sqlStatement + " :: " + nome)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Nome: " + nome)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdatePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		nome := r.FormValue("Nome")
		sqlStatement := "UPDATE planos SET nome=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(nome, id)
		log.Println("UPDATE: Id: " + id + " | Nome: " + nome)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdatePlanosHandler(planosPage []mdl.Plano, planosDB []mdl.Plano) {
	for i := range planosPage {
		id := planosPage[i].Id
		log.Println("id: " + strconv.FormatInt(id, 10))
		for j := range planosDB {
			log.Println("planosDB[j].Id: " + strconv.FormatInt(planosDB[j].Id, 10))
			if strconv.FormatInt(planosDB[j].Id, 10) == strconv.FormatInt(id, 10) {
				log.Println("Entrei")
				fieldsChanged := hasSomeFieldChangedPlano(planosPage[i], planosDB[j]) //DONE
				log.Println(fieldsChanged)
				if fieldsChanged {
					updatePlanoHandler(planosPage[i], planosDB[j])
				}
				planosDB = removePlano(planosDB, planosPage[i])
				break
			}
		}
	}
	DeletePlanosHandler(planosDB)
}

func hasSomeFieldChangedPlano(planoPage mdl.Plano, planoDB mdl.Plano) bool {
	log.Println("planoPage.Nome: " + planoPage.Nome)
	log.Println("planoDB.Nome: " + planoDB.Nome)
	if planoPage.Nome != planoDB.Nome {
		return true
	} else if planoPage.Descricao != planoDB.Descricao {
		return true
	} else {
		return false
	}
}

func updatePlanoHandler(p mdl.Plano, planoDB mdl.Plano) {
	sqlStatement := "UPDATE planos SET " +
		"nome=$1, descricao=$2 WHERE id=$3"
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	log.Println(p.Nome)
	log.Println(p.Descricao)
	log.Println(p.Id)
	_, err := updtForm.Exec(p.Nome, p.Descricao, p.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Statement: " + sqlStatement)
}

func DeletePlanoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Plano")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM planos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.PlanosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

// AJAX
func ListPlanosByEntidadeId(entidadeId string) []mdl.Plano {
	log.Println("List Planos By Entidade Id")
	log.Println("entidadeId: " + entidadeId)
	sql := "SELECT " +
		" a.id, " +
		" a.entidade_id, " +
		" a.nome," +
		" coalesce(a.descricao,''), " +
		" a.author_id, " +
		" coalesce(b.name,'') as author_name, " +
		" coalesce(to_char(a.criado_em,'DD/MM/YYYY')) as criado_em," +
		" a.status_id, " +
		" coalesce(c.name,'') as status_name " +
		" FROM planos a LEFT JOIN users b ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" WHERE a.entidade_id = $1 " +
		" ORDER BY a.nome ASC"
	log.Println(sql)
	rows, _ := Db.Query(sql, entidadeId)
	var planos []mdl.Plano
	var plano mdl.Plano
	var i = 1
	for rows.Next() {
		rows.Scan(
			&plano.Id,
			&plano.EntidadeId,
			&plano.Nome,
			&plano.Descricao,
			&plano.AuthorId,
			&plano.AuthorName,
			&plano.CriadoEm,
			&plano.StatusId,
			&plano.CStatus)
		plano.Order = i
		i++
		planos = append(planos, plano)
		log.Println(plano)
	}
	return planos
}

func DeletePlanosByEntidadeId(entidadeId string) {
	sqlStatement := "DELETE FROM Planos WHERE entidade_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(entidadeId)
	log.Println("DELETE Planos in Order Id: " + entidadeId)
}

func DeletePlanosHandler(diffDB []mdl.Plano) {
	sqlStatement := "DELETE FROM Planos WHERE id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Plano Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func containsPlano(planos []mdl.Plano, planoCompared mdl.Plano) bool {
	for n := range planos {
		if planos[n].Id == planoCompared.Id {
			return true
		}
	}
	return false
}

func removePlano(planos []mdl.Plano, planoToBeRemoved mdl.Plano) []mdl.Plano {
	var newPlanos []mdl.Plano
	for i := range planos {
		if planos[i].Id != planoToBeRemoved.Id {
			newPlanos = append(newPlanos, planos[i])
		}
	}
	return newPlanos
}
