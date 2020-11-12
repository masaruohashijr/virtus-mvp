package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	e "virtus/errors"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

func CreateFeatureHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Feature")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		name := r.FormValue("Name")
		code := r.FormValue("Code")
		description := r.FormValue("Description")
		sqlStatement := "INSERT INTO features(name, code, description) VALUES ($1, $2, $3) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, code, description).Scan(&id)
		if err != nil && strings.Contains(err.Error(), "duplicate key") {
			page := listFeatures(e.ErroChaveDuplicada)
			page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
			var tmpl = template.Must(template.ParseGlob("tiles/features/*"))
			tmpl.ParseGlob("tiles/*")
			tmpl.ExecuteTemplate(w, "Main-Status", page)
		}
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Code: " + code + " | Description: " + description)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateFeatureHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Feature")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		code := r.FormValue("Code")
		description := r.FormValue("Description")
		sqlStatement := "UPDATE features SET name=$1, code=$2, description=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(name, code, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Code: " + code + " | Description: " + description)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteFeatureHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Feature")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM features WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteFeaturesByRoleHandler(roleId string) {
	sqlStatement := "DELETE FROM features_roles WHERE role_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(roleId)
	log.Println("DELETE features_roles in Role Id: " + roleId)
}

func DeleteFeaturesHandler(diffDB []mdl.Feature) {
	sqlStatement := "DELETE FROM features_roles WHERE feature_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Feature Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func ListFeaturesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Features")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listFeatures") {
		page := listFeatures("")
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/features/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Features", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func listFeatures(errorMsg string) mdl.PageFeatures {
	sql := "SELECT " +
		" a.id, " +
		" a.name, " +
		" a.code, " +
		" coalesce(a.description,'') as dsc, " +
		" a.author_id, " +
		" b.name, " +
		" to_char(a.created_at,'DD/MM/YYYY HH24:MI:SS'), " +
		" coalesce(c.name,'') as cstatus, " +
		" a.status_id, " +
		" a.id_versao_origem " +
		" FROM features a LEFT JOIN users b " +
		" ON a.author_id = b.id " +
		" LEFT JOIN status c ON a.status_id = c.id " +
		" order by a.id asc"
	log.Println(sql)
	rows, _ := Db.Query(sql)
	var features []mdl.Feature
	var feature mdl.Feature
	var i = 1
	for rows.Next() {
		rows.Scan(
			&feature.Id,
			&feature.Name,
			&feature.Code,
			&feature.Description,
			&feature.AuthorId,
			&feature.AuthorName,
			&feature.C_CreatedAt,
			&feature.CStatus,
			&feature.StatusId,
			&feature.IdVersaoOrigem)
		log.Println(feature.AuthorName)
		feature.Order = i
		i++
		features = append(features, feature)
	}
	var page mdl.PageFeatures
	page.Features = features
	page.AppName = mdl.AppName
	page.Title = "Funcionalidades"
	if errorMsg != "" {
		page.ErrMsg = errorMsg
	}
	return page
}

// AJAX
func ListFeaturesByRoleIdHandler(roleId string) []mdl.Feature {
	log.Println("List Features By Role Id")
	sql := "SELECT feature_id" +
		" FROM features_roles WHERE role_id= $1"
	log.Println(sql)
	rows, _ := Db.Query(sql, roleId)
	var features []mdl.Feature
	var feature mdl.Feature
	for rows.Next() {
		rows.Scan(&feature.Id)
		features = append(features, feature)
	}
	return features
}

func LoadAvailableFeatures(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Load Available Features")
	r.ParseForm()
	savedUser := GetUserInCookie(w, r)
	var statusId = r.FormValue("statusId")
	var entityType = r.FormValue("entityType")
	log.Println("entityType: " + entityType)
	log.Println("statusId: " + statusId)
	sql := " SELECT a.id, a.name, a.code " +
		" FROM features a INNER JOIN features_activities b ON a.id = b.feature_id " +
		" INNER JOIN activities c ON c.id = b.activity_id " +
		" INNER JOIN actions d ON c.action_id = d.id " +
		" INNER JOIN workflows e ON c.workflow_id = e.id " +
		" WHERE e.end_at IS null " +
		" AND e.entity_type = $1 " +
		" AND d.origin_status_id = $2 " +
		" AND a.id in ( SELECT feature_id from features_roles where role_id = $3 ) "
	log.Println("Query Available Features: " + sql)
	rows, _ := Db.Query(sql, entityType, statusId, savedUser.Role)
	var features []mdl.Feature
	var feature mdl.Feature
	for rows.Next() {
		rows.Scan(&feature.Id, &feature.Name, &feature.Code)
		features = append(features, feature)
		log.Println(features)
	}
	jsonFeatures, _ := json.Marshal(features)
	w.Write([]byte(jsonFeatures))
	log.Println("JSON Load Features")
}
