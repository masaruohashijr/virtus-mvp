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

func GetStartStatus(entityType string) int {
	log.Println("List Workflows")
	query := "SELECT id FROM status where id in (select origin_status_id from actions_status where action_id in " +
		" ( select action_id from activities where workflow_id in (select id from workflows where " +
		" entity_type = $1 and end_at is null))) " +
		" and stereotype = 'Start' LIMIT 1"
	log.Println("List WF -> Query: " + query)
	log.Println("entityType: " + entityType)
	rows, _ := Db.Query(query, entityType)
	startStatusId := 0
	log.Println("startStatusId: " + strconv.Itoa(startStatusId))
	for rows.Next() {
		rows.Scan(&startStatusId)
	}
	//	log.Println("Saindo!")
	return startStatusId
}

func CreateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Workflow")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		name := r.FormValue("Name")
		description := r.FormValue("Description")
		entityType := r.FormValue("EntityTypeForInsert")
		sqlStatement := "UPDATE workflows SET end_at = $1 WHERE entity_type = $2"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(time.Now(), entityType)
		sqlStatement = "INSERT INTO " +
			" workflows(name, entity_type, start_at, description, author_id, created_at) " +
			" VALUES ($1,$2,$3,$4,$5,$6) RETURNING id "
		wId := 0
		err = Db.QueryRow(
			sqlStatement,
			name,
			entityType,
			time.Now(),
			description,
			currentUser.Id,
			time.Now()).Scan(&wId)
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Id: " + strconv.Itoa(wId) + " | Name: " + name + " | Entitity: " + entityType)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "activity") {
				array := strings.Split(value[0], "#")
				log.Println(value[0])
				activityId := 0
				actionId := strings.Split(array[3], ":")[1]
				startAt, _ := time.Parse("yyyy-mm-dd", strings.Split(array[8], ":")[1])
				endAt, _ := time.Parse("yyyy-mm-dd", strings.Split(array[9], ":")[1])
				expTime := strings.Split(array[7], ":")[1]
				if expTime == "" {
					expTime = "0"
				}
				expActionId := strings.Split(array[5], ":")[1]
				strRoles := strings.Split(array[10], ":")[1]
				log.Println("actionId: " + actionId)
				sqlStatement := "INSERT INTO " +
					" activities(workflow_id, action_id, start_at, end_at, expiration_time_days, expiration_action_id) " +
					" VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
				log.Println(sqlStatement)
				log.Println("wId: " + strconv.Itoa(wId) + " | Action: " + actionId + " | ExpDays: " + expTime + " | ExpAction: " + expActionId)
				if expActionId == "" {
					Db.QueryRow(sqlStatement, wId, actionId, startAt, endAt, expTime, nil).Scan(&activityId)
				} else {
					Db.QueryRow(sqlStatement, wId, actionId, startAt, endAt, expTime, expActionId).Scan(&activityId)
				}
				if len(strRoles) > 0 {
					log.Println("Roles: " + strRoles)
					roles := strings.Split(strRoles, ".")
					for _, roleId := range roles {
						sqlStatement := "INSERT INTO " +
							" activities_roles(activity_id, role_id) " +
							" VALUES ($1,$2)"
						log.Println(sqlStatement + " - " + strconv.Itoa(activityId) + " - " + roleId)
						Db.QueryRow(sqlStatement, activityId, roleId)
					}
				}
			}
		}
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Workflow")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		// Workflow
		wId := r.FormValue("Id")
		log.Println("Workflow Id: " + wId)
		name := r.FormValue("NameForUpdate")
		description := r.FormValue("DescriptionForUpdate")
		entity := r.FormValue("EntityTypeForUpdate")
		sqlStatement := "UPDATE workflows SET name=$1, entity_type=$2, description=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(name, entity, description, wId)
		log.Println("UPDATE: Id: " + wId + " | Name: " + name + " | Entity: " + entity + " | Description: " + description)
		// Atividades
		var actsDB = ListActivitiesHandler(wId)
		var actsPage []mdl.Activity
		var actPage mdl.Activity
		for key, value := range r.Form {
			if strings.HasPrefix(key, "activity") {
				log.Println(value[0])
				array := strings.Split(value[0], "#")
				id := strings.Split(array[1], ":")[1]
				log.Println("Id -------- " + id)
				actPage.Id, _ = strconv.ParseInt(id, 10, 64)
				wId := strings.Split(array[2], ":")[1]
				log.Println("wId -------- " + wId)
				actPage.WorkflowId, _ = strconv.ParseInt(wId, 10, 64)
				actionId := strings.Split(array[3], ":")[1]
				log.Println("actionId -------- " + actionId)
				actPage.ActionId, _ = strconv.ParseInt(actionId, 10, 64)
				actionName := strings.Split(array[4], ":")[1]
				log.Println("actionName -------- " + actionName)
				actPage.ActionName = actionName
				//				log.Println("Id -------- " + id)
				expActionId := strings.Split(array[5], ":")[1]
				//				log.Println("Id -------- " + id)
				actPage.ExpirationActionId, _ = strconv.ParseInt(expActionId, 10, 64)
				//				log.Println("Id -------- " + id)
				expActionName := strings.Split(array[6], ":")[1]
				//				log.Println("Id -------- " + id)
				actPage.ExpirationActionName = expActionName
				//				log.Println("Id -------- " + id)
				expTime := strings.Split(array[7], ":")[1]
				//				log.Println("Id -------- " + id)
				actPage.ExpirationTimeDays, _ = strconv.Atoi(expTime)
				//				log.Println("Id -------- " + id)
				startAt := strings.Split(array[8], ":")[1]
				//				log.Println("Id -------- " + id)
				actPage.CStartAt = startAt
				//				log.Println("Id -------- " + id)
				endAt := strings.Split(array[9], ":")[1]
				//				log.Println("Id -------- " + id)
				actPage.CEndAt = endAt
				roles := strings.Split(array[10], ":")[1]
				actPage.CRoles = roles
				log.Println("Roles -------- " + roles)
				features := strings.Split(array[13], ":")[1]
				actPage.CFeatures = features
				log.Println("Features -------- " + features)
				actsPage = append(actsPage, actPage)
			}
		}
		if len(actsPage) < len(actsDB) {
			log.Println("Quantidade de Activities da Página: " + strconv.Itoa(len(actsPage)))
			if len(actsPage) == 0 {
				DeleteActivitiesByWorkflowIdHandler(wId) //DONE
			} else {
				var diffDB []mdl.Activity = actsDB
				for n := range actsPage {
					if containsAct(diffDB, actsPage[n]) {
						diffDB = removeAct(diffDB, actsPage[n])
					}
				}
				DeleteActivitiesHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Activity = actsPage
			for n := range actsPage {
				log.Println("Page Action: " + actsPage[n].ActionName)
			}
			for n := range actsDB {
				if containsAct(diffPage, actsDB[n]) {
					log.Println(n)
					log.Println("actsDB[n]: " + actsDB[n].ActionName)
					diffPage = removeAct(diffPage, actsDB[n])
				}
			}
			log.Println("Tamamho: " + strconv.Itoa(len(diffPage)))
			for n := range diffPage {
				log.Println("Action Name Incluida agora " + diffPage[n].ActionName)
			}

			var act mdl.Activity
			for i := range diffPage {
				act = diffPage[i]
				log.Println("Workflow Id: " + strconv.FormatInt(act.WorkflowId, 10))
				sqlStatement := "INSERT INTO " +
					"activities(workflow_id, action_id, start_at, end_at, expiration_time_days, expiration_action_id) " +
					"VALUES ($1,$2,TO_TIMESTAMP($3, 'YYYY-MM-DD HH24:MI:SS'),TO_TIMESTAMP($4, 'YYYY-MM-DD HH24:MI:SS'),$5,$6) RETURNING id"
				log.Println(sqlStatement)
				var activityId int
				log.Println("wId: " + wId + " | Action: " + strconv.FormatInt(act.ActionId, 10) + " | ExpDays: " + strconv.Itoa(act.ExpirationTimeDays) + " | ExpAction: " + strconv.FormatInt(act.ExpirationActionId, 10))
				if act.ExpirationActionId == 0 {
					log.Println("entrei aqui")
					err := Db.QueryRow(sqlStatement, wId, act.ActionId, act.CStartAt, act.CEndAt, act.ExpirationTimeDays, nil).Scan(&activityId)
					if err != nil {
						panic(err.Error())
					}
				} else {
					log.Println("entrei acolá")
					err := Db.QueryRow(sqlStatement, wId, act.ActionId, act.CStartAt, act.CEndAt, act.ExpirationTimeDays, act.ExpirationActionId).Scan(&activityId)
					if err != nil {
						panic(err.Error())
					}
				}
				log.Println("Papel: " + act.CRoles)
				strRoles := strings.Split(act.CRoles, ".")
				for _, roleId := range strRoles {
					sqlStatement := "INSERT INTO " +
						"activities_roles(activity_id, role_id) " +
						"VALUES ($1,$2)"
					log.Println(sqlStatement + " - " + strconv.Itoa(activityId) + " - " + roleId)
					Db.QueryRow(sqlStatement, activityId, roleId)
				}
			}
		}
		UpdateActivitiesHandler(actsPage, actsDB) // TODO
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func containsAct(acts []mdl.Activity, actCompared mdl.Activity) bool {
	for n := range acts {
		if acts[n].Id == actCompared.Id {
			return true
		}
	}
	return false
}

func removeAct(acts []mdl.Activity, actToBeRemoved mdl.Activity) []mdl.Activity {
	var newActs []mdl.Activity
	for i := range acts {
		if acts[i].Id != actToBeRemoved.Id {
			newActs = append(newActs, acts[i])
		}
	}
	return newActs
}

func DeleteWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Workflow")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM activities_roles " +
			" WHERE activity_id IN (" +
			" SELECT id FROM activities WHERE workflow_id = $1)"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM activities WHERE workflow_id = $1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM workflows WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListWorkflowsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Workflows")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listWorkflows") {
		sql := "SELECT " +
			" a.id, " +
			" a.name, " +
			" a.entity_type, " +
			" coalesce(to_char(a.start_at,'DD/MM/YYYY'),'') as c_start_at, " +
			" coalesce(to_char(a.end_at,'DD/MM/YYYY'),'') as c_end_at, " +
			" coalesce(a.description,'') as dsc, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.created_at,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM " +
			" workflows a " +
			" LEFT JOIN users b ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" ORDER BY a.id ASC"

		log.Println("List WF -> SQL: " + sql)
		rows, _ := Db.Query(sql)
		var workflows []mdl.Workflow
		var workflow mdl.Workflow
		var i = 1
		for rows.Next() {
			rows.Scan(
				&workflow.Id,
				&workflow.Name,
				&workflow.EntityType,
				&workflow.StartAt,
				&workflow.EndAt,
				&workflow.Description,
				&workflow.AuthorId,
				&workflow.AuthorName,
				&workflow.C_CreatedAt,
				&workflow.CStatus,
				&workflow.StatusId,
				&workflow.IdVersaoOrigem)
			workflow.Order = i
			i++
			workflows = append(workflows, workflow)
		}
		sql = " SELECT " +
			" a.id, " +
			" a.name, " +
			" a.origin_status_id, " +
			" b.name as origin_status, " +
			" a.destination_status_id, " +
			" c.name as destination_status, " +
			" a.other_than " +
			" FROM " +
			" actions a " +
			" LEFT JOIN status b ON a.origin_status_id = b.id " +
			" LEFT JOIN status c ON a.destination_status_id = c.id " +
			" ORDER BY a.id asc"
		log.Println("List WF -> sql: " + sql)
		rows, _ = Db.Query(sql)
		var actions []mdl.Action
		var action mdl.Action
		i = 1
		for rows.Next() {
			rows.Scan(
				&action.Id,
				&action.Name,
				&action.OriginId,
				&action.Origin,
				&action.DestinationId,
				&action.Destination,
				&action.OtherThan)
			action.Order = i
			i++
			actions = append(actions, action)
		}
		sql = "SELECT id, name FROM roles order by name asc"
		log.Println("List WF -> Query: " + sql)
		rows, _ = Db.Query(sql)
		var roles []mdl.Role
		var role mdl.Role
		i = 1
		for rows.Next() {
			rows.Scan(&role.Id, &role.Name)
			role.Order = i
			i++
			roles = append(roles, role)
		}

		sql = "SELECT id, name " +
			" FROM features order by id desc"
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var features []mdl.Feature
		var feature mdl.Feature
		for rows.Next() {
			rows.Scan(&feature.Id, &feature.Name)
			features = append(features, feature)
		}

		var page mdl.PageWorkflows
		page.Actions = actions
		page.Features = features
		page.Roles = roles
		page.Workflows = workflows
		page.AppName = mdl.AppName
		page.Title = "Workflows"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/workflows/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Workflows", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadActivitiesByWorkflowId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Activities By Workflow Id")
	r.ParseForm()
	var idWF = r.FormValue("idWF")
	log.Println("idWF: " + idWF)
	activities := ListActivitiesHandler(idWF)
	jsonActivities, _ := json.Marshal(activities)
	w.Write([]byte(jsonActivities))
	log.Println("JSON")
}
