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

func CreateRoleHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create Role")
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("Name")
		features := r.Form["FeaturesForInsert"]
		sqlStatement := "INSERT INTO roles(name) VALUES ($1) RETURNING id"
		roleId := 0
		err := Db.QueryRow(sqlStatement, name).Scan(&roleId)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		for _, featureId := range features {
			sqlStatement := "INSERT INTO features_roles(feature_id,role_id) VALUES ($1,$2) RETURNING id"
			featureRoleId := 0
			err = Db.QueryRow(sqlStatement, featureId, roleId).Scan(&featureRoleId)
			sec.CheckInternalServerError(err, w)
			if err != nil {
				panic(err.Error())
			}
			sec.CheckInternalServerError(err, w)
		}
		log.Println("INSERT: Id: " + strconv.Itoa(roleId) + " | Name: " + name)
	}
	http.Redirect(w, r, route.RolesRoute, 301)
}

func UpdateRoleHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Update Role")
	if r.Method == "POST" {
		roleId := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE roles SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, roleId)
		log.Println("UPDATE: Id: " + roleId + " | Name: " + name)

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
	log.Println("Delete Role")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM features_roles WHERE role_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM roles WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.RolesRoute, 301)
}

func ListRolesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Roles")
	sec.IsAuthenticated(w, r)
	rows, err := Db.Query("SELECT id, name FROM roles order by id asc")
	sec.CheckInternalServerError(err, w)
	var roles []mdl.Role
	var role mdl.Role
	var i = 1
	for rows.Next() {
		err = rows.Scan(&role.Id, &role.Name)
		sec.CheckInternalServerError(err, w)
		role.Order = i
		i++
		roles = append(roles, role)
	}
	rows, err = Db.Query("SELECT id, name FROM features order by name asc")
	sec.CheckInternalServerError(err, w)
	var features []mdl.Feature
	var feature mdl.Feature
	i = 1
	for rows.Next() {
		err = rows.Scan(&feature.Id, &feature.Name)
		sec.CheckInternalServerError(err, w)
		feature.Order = i
		i++
		features = append(features, feature)
	}
	var page mdl.PageRoles
	page.Roles = roles
	page.Features = features
	page.AppName = mdl.AppName
	page.Title = "Papéis"
	page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
	var tmpl = template.Must(template.ParseGlob("tiles/roles/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Roles", page)
	sec.CheckInternalServerError(err, w)
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
func ListRolesByActionIdHandler(actionId string) []mdl.Role {
	log.Println("List Roles By Action Id")
	sql := "SELECT role_id" +
		" FROM actions_roles WHERE action_id= $1"
	log.Println(sql)
	rows, _ := Db.Query(sql, actionId)
	var roles []mdl.Role
	var role mdl.Role
	for rows.Next() {
		rows.Scan(&role.Id)
		roles = append(roles, role)
	}
	return roles
}

func DeleteRolesByActionHandler(actionId string) {
	sqlStatement := "DELETE FROM actions_roles WHERE action_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(actionId)
	log.Println("DELETE actions_roles in Action Id: " + actionId)
}

func DeleteRolesHandler(diffDB []mdl.Role) {
	sqlStatement := "DELETE FROM actions_roles WHERE role_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Role Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}
