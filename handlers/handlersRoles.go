package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateRoleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Role")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		r.ParseForm()
		name := r.FormValue("Name")
		description := r.FormValue("Description")
		features := r.Form["FeaturesForInsert"]
		sqlStatement := "INSERT INTO roles(name, description, author_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
		roleId := 0
		err := Db.QueryRow(sqlStatement, name, description, currentUser.Id, time.Now()).Scan(&roleId)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			log.Println(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		for _, featureId := range features {
			sqlStatement := "INSERT INTO features_roles(feature_id,role_id) VALUES ($1,$2) RETURNING id"
			featureRoleId := 0
			err = Db.QueryRow(sqlStatement, featureId, roleId).Scan(&featureRoleId)
			sec.CheckInternalServerError(err, w)
			if err != nil {
				log.Println(err.Error())
			}
			sec.CheckInternalServerError(err, w)
		}
		log.Println("INSERT: Id: " + strconv.Itoa(roleId) + " | Name: " + name)
	}
	http.Redirect(w, r, route.RolesRoute, 301)
}

func UpdateRoleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Role")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		roleId := r.FormValue("Id")
		name := r.FormValue("Name")
		description := r.FormValue("Description")
		sqlStatement := "UPDATE roles SET name=$1, description=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		updtForm.Exec(name, description, roleId)
		log.Println("UPDATE: Id: " + roleId + " | Name: " + name + " | Description: " + description)

		var featuresDB = ListFeaturesByRoleIdHandler(roleId)
		var featuresPage []mdl.Feature
		var featurePage mdl.Feature
		for _, featureId := range r.Form["FeaturesForUpdate"] {
			featurePage.Id, _ = strconv.ParseInt(featureId, 10, 64)
			featuresPage = append(featuresPage, featurePage)
		}
		if len(featuresPage) < len(featuresDB) {
			log.Println("Quantidade de Features da Página: " + strconv.Itoa(len(featuresPage)))
			if len(featuresPage) == 0 {
				DeleteFeaturesByRoleHandler(roleId) //DONE
			} else {
				var diffDB []mdl.Feature = featuresDB
				for n := range featuresPage {
					if containsFeature(diffDB, featuresPage[n]) {
						diffDB = removeFeature(diffDB, featuresPage[n])
					}
				}
				DeleteFeaturesHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Feature = featuresPage
			for n := range featuresDB {
				if containsFeature(diffPage, featuresDB[n]) {
					diffPage = removeFeature(diffPage, featuresDB[n])
				}
			}
			var feature mdl.Feature
			for i := range diffPage {
				feature = diffPage[i]
				log.Println("Role Id: " + roleId)
				sqlStatement := "INSERT INTO features_roles(role_id, feature_id) VALUES ($1,$2)"
				log.Println(sqlStatement)
				Db.QueryRow(sqlStatement, roleId, feature.Id)
			}
		}
	}
	http.Redirect(w, r, route.RolesRoute, 301)
}

func containsFeature(features []mdl.Feature, featureCompared mdl.Feature) bool {
	for n := range features {
		if features[n].Id == featureCompared.Id {
			return true
		}
	}
	return false
}

func removeFeature(features []mdl.Feature, featureToBeRemoved mdl.Feature) []mdl.Feature {
	var newFeatures []mdl.Feature
	for i := range features {
		if features[i].Id != featureToBeRemoved.Id {
			newFeatures = append(newFeatures, features[i])
		}
	}
	return newFeatures
}

func DeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Perfil")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM features_roles WHERE role_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM roles WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			log.Println(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.RolesRoute, 301)
}

func ListPerfisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Perfis")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listRoles") {
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, " +
			" a.name, " +
			" a.description, " +
			" a.author_id, " +
			" b.name, " +
			" to_char(a.created_at,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(c.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM roles a LEFT JOIN users b " +
			" ON a.author_id = b.id " +
			" LEFT JOIN status c ON a.status_id = c.id " +
			" order by id asc"
		log.Println("sql: " + sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var roles []mdl.Role
		var role mdl.Role
		var i = 1
		for rows.Next() {
			rows.Scan(
				&role.Id,
				&role.Name,
				&role.Description,
				&role.AuthorId,
				&role.AuthorName,
				&role.C_CreatedAt,
				&role.CStatus,
				&role.StatusId,
				&role.IdVersaoOrigem)
			role.Order = i
			i++
			roles = append(roles, role)
		}
		rows, _ = Db.Query("SELECT id, name FROM features order by name asc")
		defer rows.Close()
		var features []mdl.Feature
		var feature mdl.Feature
		i = 1
		for rows.Next() {
			rows.Scan(&feature.Id, &feature.Name)
			feature.Order = i
			i++
			features = append(features, feature)
		}
		var page mdl.PageRoles
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		page.Roles = roles
		page.Features = features
		page.AppName = mdl.AppName
		page.Title = "Perfis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/perfis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Perfis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadFeaturesByRoleId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Features By Role Id")
	r.ParseForm()
	var roleId = r.FormValue("roleId")
	log.Println("roleId: " + roleId)
	features := ListFeaturesByRoleIdHandler(roleId)
	jsonFeatures, _ := json.Marshal(features)
	w.Write([]byte(jsonFeatures))
	log.Println("JSON")
}

// AJAX
func ListPerfisByActionIdHandler(actionId string) []mdl.Role {
	log.Println("List Perfis By Action Id")
	sql := "SELECT role_id" +
		" FROM actions_roles WHERE action_id= $1"
	log.Println(sql)
	rows, _ := Db.Query(sql, actionId)
	defer rows.Close()
	var roles []mdl.Role
	var role mdl.Role
	for rows.Next() {
		rows.Scan(&role.Id)
		roles = append(roles, role)
	}
	return roles
}

func DeletePerfisByActionHandler(actionId string) {
	sqlStatement := "DELETE FROM actions_roles WHERE action_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	deleteForm.Exec(actionId)
	log.Println("DELETE actions_roles in Action Id: " + actionId)
}

func DeletePerfisHandler(diffDB []mdl.Role) {
	sqlStatement := "DELETE FROM actions_roles WHERE role_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Role Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}
