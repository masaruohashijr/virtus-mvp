package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	dpk "virtus/db"
	hd "virtus/handlers"
	route "virtus/routes"
	sec "virtus/security"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func dbConn() *sql.DB {
	dbase, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	log.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	return dbase
}

func main() {
	sec.Store = sessions.NewCookieStore([]byte("vindixit123581321"))
	hd.Db = dbConn()
	// injetando a variável Authenticated
	dpk.Initialize()
	r := mux.NewRouter()

	r.HandleFunc("/", hd.IndexHandler).Methods("GET")
	r.HandleFunc("/login", hd.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", hd.LogoutHandler).Methods("GET")
	// ----------------- WORKFLOWS
	r.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler).Methods("GET")
	r.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler).Methods("POST")
	r.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler).Methods("POST")
	r.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler).Methods("POST")
	// ----------------- ACTIONS
	r.HandleFunc(route.ActionsRoute, hd.ListActionsHandler).Methods("GET")
	r.HandleFunc("/createAction", hd.CreateActionHandler).Methods("POST")
	r.HandleFunc("/updateAction", hd.UpdateActionHandler).Methods("POST")
	r.HandleFunc("/deleteAction", hd.DeleteActionHandler).Methods("POST")
	// ----------------- STATUS
	r.HandleFunc(route.StatusRoute, hd.ListStatusHandler).Methods("GET")
	r.HandleFunc("/createStatus", hd.CreateStatusHandler).Methods("POST")
	r.HandleFunc("/updateStatus", hd.UpdateStatusHandler).Methods("POST")
	r.HandleFunc("/deleteStatus", hd.DeleteStatusHandler).Methods("POST")
	// ----------------- FEATURES
	r.HandleFunc(route.FeaturesRoute, hd.ListFeaturesHandler).Methods("GET")
	r.HandleFunc("/createFeature", hd.CreateFeatureHandler).Methods("POST")
	r.HandleFunc("/updateFeature", hd.UpdateFeatureHandler).Methods("POST")
	r.HandleFunc("/deleteFeature", hd.DeleteFeatureHandler).Methods("POST")
	// ----------------- ROLES
	r.HandleFunc(route.RolesRoute, hd.ListPerfisHandler).Methods("GET")
	r.HandleFunc("/createRole", hd.CreateRoleHandler).Methods("POST")
	r.HandleFunc("/updateRole", hd.UpdateRoleHandler).Methods("POST")
	r.HandleFunc("/deleteRole", hd.DeleteRoleHandler).Methods("POST")
	// ----------------- USERS
	r.HandleFunc(route.UsersRoute, hd.ListUsersHandler).Methods("GET")
	r.HandleFunc("/createUser", hd.CreateUserHandler).Methods("POST")
	r.HandleFunc("/updateUser", hd.UpdateUserHandler).Methods("POST")
	r.HandleFunc("/deleteUser", hd.DeleteUserHandler).Methods("POST")
	// ----------------- ENTIDADES
	r.HandleFunc(route.EntidadesRoute, hd.ListEntidadesHandler).Methods("GET")
	r.HandleFunc("/createEntidade", hd.CreateEntidadeHandler).Methods("POST")
	r.HandleFunc("/updateEntidade", hd.UpdateEntidadeHandler).Methods("POST")
	r.HandleFunc("/deleteEntidade", hd.DeleteEntidadeHandler).Methods("POST")
	// ----------------- CICLOS
	r.HandleFunc(route.CiclosRoute, hd.ListCiclosHandler).Methods("GET")
	r.HandleFunc("/createCiclo", hd.CreateCicloHandler).Methods("POST")
	r.HandleFunc("/updateCiclo", hd.UpdateCicloHandler).Methods("POST")
	r.HandleFunc("/deleteCiclo", hd.DeleteCicloHandler).Methods("POST")
	// ----------------- PILARES
	r.HandleFunc(route.PilaresRoute, hd.ListPilaresHandler).Methods("GET")
	r.HandleFunc("/createPilar", hd.CreatePilarHandler).Methods("POST")
	r.HandleFunc("/updatePilar", hd.UpdatePilarHandler).Methods("POST")
	r.HandleFunc("/deletePilar", hd.DeletePilarHandler).Methods("POST")
	// ----------------- ESCRITORIOS
	r.HandleFunc(route.EscritoriosRoute, hd.ListEscritoriosHandler).Methods("GET")
	r.HandleFunc("/createEscritorio", hd.CreateEscritorioHandler).Methods("POST")
	r.HandleFunc("/updateEscritorio", hd.UpdateEscritorioHandler).Methods("POST")
	r.HandleFunc("/deleteEscritorio", hd.DeleteEscritorioHandler).Methods("POST")
	// ----------------- COMPONENTES
	r.HandleFunc(route.ComponentesRoute, hd.ListComponentesHandler).Methods("GET")
	r.HandleFunc("/createComponente", hd.CreateComponenteHandler).Methods("POST")
	r.HandleFunc("/updateComponente", hd.UpdateComponenteHandler).Methods("POST")
	r.HandleFunc("/deleteComponente", hd.DeleteComponenteHandler).Methods("POST")
	// ----------------- ELEMENTOS
	r.HandleFunc(route.ElementosRoute, hd.ListElementosHandler).Methods("GET")
	r.HandleFunc("/createElemento", hd.CreateElementoHandler).Methods("POST")
	r.HandleFunc("/updateElemento", hd.UpdateElementoHandler).Methods("POST")
	r.HandleFunc("/deleteElemento", hd.DeleteElementoHandler).Methods("POST")
	// ----------------- Loads
	r.HandleFunc("/loadPlanosByEntidadeId", hd.LoadPlanosByEntidadeId).Methods("GET")
	r.HandleFunc("/loadCiclosByEntidadeId", hd.LoadCiclosByEntidadeId).Methods("GET")
	r.HandleFunc("/loadPilaresByCicloId", hd.LoadPilaresByCicloId).Methods("GET")
	r.HandleFunc("/loadItensByElementoId", hd.LoadItensByElementoId).Methods("GET")
	r.HandleFunc("/loadFeaturesByRoleId", hd.LoadFeaturesByRoleId).Methods("GET")
	r.HandleFunc("/loadRolesByActionId", hd.LoadRolesByActionId).Methods("GET")
	r.HandleFunc("/loadActivitiesByWorkflowId", hd.LoadActivitiesByWorkflowId).Methods("GET")
	r.HandleFunc("/loadAllowedActions", hd.LoadAllowedActions).Methods("GET")
	r.HandleFunc("/loadAvailableFeatures", hd.LoadAvailableFeatures).Methods("GET")
	r.HandleFunc("/executeAction", hd.ExecuteActionHandler).Methods("GET")
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.Handle("/", r)
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	defer hd.Db.Close()
}
