package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func ExecuteActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Execute Action")
	r.ParseForm()
	var entityType = r.FormValue("entityType")
	var id = r.FormValue("id")
	var actionId = r.FormValue("actionId")
	log.Println("id: " + id)
	log.Println("actionId: " + actionId)
	// atualiza o status do Pedido e retorna o novo statusId
	sec.IsAuthenticated(w, r)
	log.Println("Update Action")
	var tableName = ""
	if entityType == "elemento" {
		tableName = "elementos"
	} else if entityType == "item" {
		tableName = "itens"
	}
	// verificar brecha de segurança aqui acesso GET com parametros.
	sqlStatement := "update " + tableName + " set status_id = " +
		" (select destination_status_id from actions " +
		" where id = $1) where id = $2"
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	sec.CheckInternalServerError(err, w)
	if err != nil {
		panic(err.Error())
	}
	sec.CheckInternalServerError(err, w)
	updtForm.Exec(actionId, id)
	log.Println("UPDATE: Id: " + actionId)

	sqlStatement = "SELECT a.status_id, b.name FROM " + tableName + " a LEFT JOIN status b ON a.status_id = b.id WHERE a.id = $1"
	log.Println("Query: " + sqlStatement)
	rows, _ := Db.Query(sqlStatement, id)
	var status mdl.Status
	for rows.Next() {
		err = rows.Scan(&status.Id, &status.Name)
		sec.CheckInternalServerError(err, w)
	}
	log.Println("Retornando o Status: " + strconv.FormatInt(status.Id, 10) + " - " + status.Name)
	jsonStatus, _ := json.Marshal(status)
	w.Write([]byte(jsonStatus))
	log.Println("JSON ExecuteAction")
}

func CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Action")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		name := r.FormValue("Name")
		except := r.FormValue("Except")
		otherThan := false
		if except != "" {
			otherThan = true
		}
		log.Println(except)
		originStatus := r.Form["OriginStatusForInsert"]
		log.Println(originStatus)
		destinationStatus := r.Form["DestinationStatusForInsert"]
		log.Println(destinationStatus)
		sqlStatement := "INSERT INTO actions(name, origin_status_id, destination_status_id, other_than) VALUES ($1, $2, $3, $4) RETURNING id"
		actionId := 0
		err := Db.QueryRow(sqlStatement, name, originStatus[0], destinationStatus[0], otherThan).Scan(&actionId)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		sqlStatement = "INSERT INTO actions_status(action_id,origin_status_id,destination_status_id) VALUES ($1,$2,$3)"
		Db.QueryRow(sqlStatement, actionId, originStatus[0], destinationStatus[0])
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(actionId) + " | Name: " + name)
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Action")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		actionId := r.FormValue("Id")
		name := r.FormValue("Name")
		except := r.FormValue("ExceptForUpdate")
		otherThan := false
		if except != "" {
			otherThan = true
		}
		log.Println(except)
		originStatus := r.Form["OriginStatusForUpdate"]
		log.Println(originStatus)
		destinationStatus := r.Form["DestinationStatusForUpdate"]
		log.Println(destinationStatus)
		query := "SELECT origin_status_id, destination_status_id FROM actions_status WHERE action_id = $1 "
		log.Println("List Action -> Query: " + query)
		rows, err := Db.Query(query, actionId)
		sec.CheckInternalServerError(err, w)
		originStatusDB := ""
		destinationStatusDB := ""
		for rows.Next() {
			rows.Scan(&originStatusDB, &destinationStatusDB)
		}

		if originStatus[0] != originStatusDB || destinationStatus[0] != destinationStatusDB {
			sqlStatement := "DELETE FROM actions_status WHERE action_id=$1"
			deleteForm, err := Db.Prepare(sqlStatement)
			if err != nil {
				panic(err.Error())
			}
			deleteForm.Exec(actionId)
			sec.CheckInternalServerError(err, w)
			log.Println("DELETE Action_Status: Id: " + actionId)
		}
		sqlStatement := "UPDATE actions SET name=$1, origin_status_id=$2, destination_status_id=$3, other_than=$4 WHERE id=$5"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, originStatus[0], destinationStatus[0], otherThan, actionId)
		log.Println("UPDATE: Id: " + actionId + " | Name: " + name)
		sqlStatement = "INSERT INTO actions_status(action_id,origin_status_id,destination_status_id) VALUES ($1,$2,$3)"
		Db.QueryRow(sqlStatement, actionId, originStatus[0], destinationStatus[0])
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func containsRole(roles []mdl.Role, roleCompared mdl.Role) bool {
	for n := range roles {
		if roles[n].Id == roleCompared.Id {
			return true
		}
	}
	return false
}

func removeRole(roles []mdl.Role, roleToBeRemoved mdl.Role) []mdl.Role {
	var newRoles []mdl.Role
	for i := range roles {
		if roles[i].Id != roleToBeRemoved.Id {
			newRoles = append(newRoles, roles[i])
		}
	}
	return newRoles
}

func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Action")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM actions_status WHERE action_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM actions WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListActionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Actions")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT a.id, a.name, a.origin_status_id, b.name as origin_name, a.destination_status_id, c.name as destination_name, a.other_than " +
			"FROM actions a, status b, status c where a.origin_status_id = b.id and a.destination_status_id = c.id order by id asc"
		log.Println("List Action -> Query: " + query)
		rows, err := Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var actions []mdl.Action
		var action mdl.Action
		var i = 1
		for rows.Next() {
			err = rows.Scan(&action.Id, &action.Name, &action.OriginId, &action.Origin, &action.DestinationId, &action.Destination, &action.OtherThan)
			sec.CheckInternalServerError(err, w)
			action.Order = i
			i++
			actions = append(actions, action)
		}
		query = "SELECT id, name, stereotype FROM status order by name asc"
		log.Println("List Action -> Query: " + query)
		rows, err = Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var statuss []mdl.Status
		var status mdl.Status
		i = 1
		for rows.Next() {
			err = rows.Scan(&status.Id, &status.Name, &status.Stereotype)
			sec.CheckInternalServerError(err, w)
			status.Order = i
			i++
			statuss = append(statuss, status)
		}
		var page mdl.PageActions
		page.Statuss = statuss
		page.Actions = actions
		page.AppName = mdl.AppName
		page.Title = "Ação"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/actions/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Actions", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadRolesByActionId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Roles By Action Id")
	r.ParseForm()
	var actionId = r.FormValue("actionId")
	log.Println("actionId: " + actionId)
	roles := ListRolesByActionIdHandler(actionId)
	jsonRoles, _ := json.Marshal(roles)
	w.Write([]byte(jsonRoles))
	log.Println("JSON")
}

func LoadAllowedActions(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Allowed Actions")
	r.ParseForm()
	savedUser := GetUserInCookie(w, r)
	roleId := savedUser.Role
	var statusId = r.FormValue("statusId")
	var entityType = r.FormValue("entityType")
	log.Println("entityType: " + entityType)
	log.Println("statusId: " + statusId)
	sql := " select id, name from actions where " +
		" (not other_than and origin_status_id = $1 " +
		" and id in ( select a.action_id from activities a, activities_roles b " +
		" where a.workflow_id = ( select id from workflows where entity_type = $2 and end_at is null) " +
		" and a.id = b.activity_id and b.role_id = $3 ) ) " +
		" or " +
		" (other_than and origin_status_id != $4 " +
		" and id in ( select a.action_id from activities a, activities_roles b " +
		" where a.workflow_id = ( select id from workflows where entity_type = $5 and end_at is null) ) ) " +
		" order by other_than asc "
	log.Println("Query: " + sql)
	rows, _ := Db.Query(sql, statusId, entityType, roleId, statusId, entityType)
	var actions []mdl.Action
	var action mdl.Action
	for rows.Next() {
		rows.Scan(&action.Id, &action.Name)
		actions = append(actions, action)
		log.Println(actions)
	}
	jsonActions, _ := json.Marshal(actions)
	w.Write([]byte(jsonActions))
	log.Println("JSON Load Actions")
}
