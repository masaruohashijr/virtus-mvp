package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(actionId, id)
	log.Println("UPDATE: Id: " + actionId)

	sqlStatement = "SELECT a.status_id, b.name FROM " + tableName + " a LEFT JOIN status b ON a.status_id = b.id WHERE a.id = $1"
	log.Println("Query: " + sqlStatement)
	rows, _ := Db.Query(sqlStatement, id)
	defer rows.Close()
	var status mdl.Status
	for rows.Next() {
		err = rows.Scan(&status.Id, &status.Name)
	}
	log.Println("Retornando o Status: " + strconv.FormatInt(status.Id, 10) + " - " + status.Name)
	jsonStatus, _ := json.Marshal(status)
	w.Write([]byte(jsonStatus))
	log.Println("JSON ExecuteAction")
}

func CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Action")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
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
		description := r.FormValue("DescriptionForInsert")
		sqlStatement := "INSERT INTO actions(name, origin_status_id, destination_status_id, other_than, description, author_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
		actionId := 0
		err := Db.QueryRow(
			sqlStatement,
			name,
			originStatus[0],
			destinationStatus[0],
			otherThan,
			description,
			currentUser.Id,
			time.Now()).Scan(&actionId)
		if err != nil {
			log.Println(err.Error())
		}
		sqlStatement = "INSERT INTO actions_status(action_id,origin_status_id,destination_status_id) VALUES ($1,$2,$3)"
		Db.QueryRow(sqlStatement, actionId, originStatus[0], destinationStatus[0])
		if err != nil {
			log.Println(err.Error())
		}
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
		description := r.FormValue("Description")
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
		rows, _ := Db.Query(query, actionId)
		defer rows.Close()
		originStatusDB := ""
		destinationStatusDB := ""
		for rows.Next() {
			rows.Scan(&originStatusDB, &destinationStatusDB)
		}

		if originStatus[0] != originStatusDB || destinationStatus[0] != destinationStatusDB {
			sqlStatement := "DELETE FROM actions_status WHERE action_id=$1"
			deleteForm, _ := Db.Prepare(sqlStatement)
			_, err := deleteForm.Exec(actionId)
			if err != nil && strings.Contains(err.Error(), "violates foreign key") {
				http.Redirect(w, r, route.ActionsRoute+"?errMsg=Action está vinculada e não foi removida.", 301)
			} else {
				http.Redirect(w, r, route.ActionsRoute+"?msg=Action removida com sucesso.", 301)
			}
			log.Println("DELETE Action_Status: Id: " + actionId)
		}
		sqlStatement := "UPDATE actions SET name=$1, origin_status_id=$2, destination_status_id=$3, other_than=$4, description=$5 WHERE id=$6"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		updtForm.Exec(name, originStatus[0], destinationStatus[0], otherThan, description, actionId)
		log.Println("UPDATE: Id: " + actionId + " | Name: " + name)
		sqlStatement = "INSERT INTO actions_status(action_id,origin_status_id,destination_status_id) VALUES ($1,$2,$3)"
		Db.QueryRow(sqlStatement, actionId, originStatus[0], destinationStatus[0])
		if err != nil {
			log.Println(err.Error())
		}
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
		deleteForm, _ := Db.Prepare(sqlStatement)
		_, err := deleteForm.Exec(id)
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			http.Redirect(w, r, route.ActionsRoute+"?errMsg=A Ação está vinculada e não foi removida.", 301)
		} else {
			http.Redirect(w, r, route.ActionsRoute+"?msg=Ação removida com sucesso.", 301)
		}
		sqlStatement = "DELETE FROM actions WHERE id=$1"
		deleteForm, _ = Db.Prepare(sqlStatement)
		deleteForm.Exec(id)
		if err != nil && strings.Contains(err.Error(), "violates foreign key") {
			http.Redirect(w, r, route.ActionsRoute+"?errMsg=A Ação está vinculada e não foi removida.", 301)
		} else {
			http.Redirect(w, r, route.ActionsRoute, 301)
		}
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListActionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Actions")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listActions") {
		errMsg := r.FormValue("errMsg")
		sql := " SELECT " +
			" a.id, " +
			" a.name, " +
			" a.description, " +
			" a.origin_status_id, " +
			" b.name as origin_name, " +
			" a.destination_status_id, " +
			" c.name as destination_name, " +
			" a.other_than, " +
			" a.author_id, " +
			" d.name, " +
			" to_char(a.created_at,'DD/MM/YYYY HH24:MI:SS') as created_at, " +
			" coalesce(e.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM actions a " +
			" LEFT JOIN status b ON a.origin_status_id = b.id " +
			" LEFT JOIN status c ON a.destination_status_id = c.id " +
			" LEFT JOIN users d ON a.author_id = d.id " +
			" LEFT JOIN status e ON a.status_id = c.id " +
			" ORDER BY a.id asc "
		log.Println("List Action -> SQL: " + sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var actions []mdl.Action
		var action mdl.Action
		var i = 1
		for rows.Next() {
			rows.Scan(
				&action.Id,
				&action.Name,
				&action.Description,
				&action.OriginId,
				&action.Origin,
				&action.DestinationId,
				&action.Destination,
				&action.OtherThan,
				&action.AuthorId,
				&action.AuthorName,
				&action.C_CreatedAt,
				&action.CStatus,
				&action.StatusId,
				&action.IdVersaoOrigem)
			action.Order = i
			i++
			actions = append(actions, action)
		}
		sql = "SELECT id, name, stereotype FROM status ORDER BY name asc"
		log.Println("List Action -> Query: " + sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var statuss []mdl.Status
		var status mdl.Status
		i = 1
		for rows.Next() {
			rows.Scan(&status.Id, &status.Name, &status.Stereotype)
			status.Order = i
			i++
			statuss = append(statuss, status)
		}
		var page mdl.PageActions
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
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
	roles := ListPerfisByActionIdHandler(actionId)
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
	defer rows.Close()
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
