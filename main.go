package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	dpk "virtus/db"
	hd "virtus/handlers"
	route "virtus/routes"
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

	hd.Db = dbConn()
	// injeta	ndo a variável Authenticated
	dpk.Initialize()
	http.HandleFunc("/", hd.IndexHandler)
	http.HandleFunc("/login", hd.LoginHandler)
	// ----------------- WORKFLOWS
	http.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler)
	http.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler)
	http.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler)
	http.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler)
	// ----------------- ACTIONS
	http.HandleFunc(route.ActionsRoute, hd.ListActionsHandler)
	http.HandleFunc("/createAction", hd.CreateActionHandler)
	http.HandleFunc("/updateAction", hd.UpdateActionHandler)
	http.HandleFunc("/deleteAction", hd.DeleteActionHandler)
	// ----------------- STATUS
	http.HandleFunc(route.StatusRoute, hd.ListStatusHandler)
	http.HandleFunc("/createStatus", hd.CreateStatusHandler)
	http.HandleFunc("/updateStatus", hd.UpdateStatusHandler)
	http.HandleFunc("/deleteStatus", hd.DeleteStatusHandler)
	// ----------------- FEATURES
	http.HandleFunc(route.FeaturesRoute, hd.ListFeaturesHandler)
	http.HandleFunc("/createFeature", hd.CreateFeatureHandler)
	http.HandleFunc("/updateFeature", hd.UpdateFeatureHandler)
	http.HandleFunc("/deleteFeature", hd.DeleteFeatureHandler)
	// ----------------- ROLES
	http.HandleFunc(route.RolesRoute, hd.ListRolesHandler)
	http.HandleFunc("/createRole", hd.CreateRoleHandler)
	http.HandleFunc("/updateRole", hd.UpdateRoleHandler)
	http.HandleFunc("/deleteRole", hd.DeleteRoleHandler)
	// ----------------- BEERS
	//	http.HandleFunc(route.BeersRoute, hd.ListBeersHandler)
	//	http.HandleFunc("/createBeer", hd.CreateBeerHandler)
	//	http.HandleFunc("/updateBeer", hd.UpdateBeerHandler)
	//	http.HandleFunc("/deleteBeer", hd.DeleteBeerHandler)
	// ----------------- USERS
	http.HandleFunc(route.UsersRoute, hd.ListUsersHandler)
	http.HandleFunc("/createUser", hd.CreateUserHandler)
	http.HandleFunc("/updateUser", hd.UpdateUserHandler)
	http.HandleFunc("/deleteUser", hd.DeleteUserHandler)
	// ----------------- ORDERS
	http.HandleFunc(route.OrdersRoute, hd.ListOrdersHandler)
	http.HandleFunc("/createOrder", hd.CreateOrderHandler)
	http.HandleFunc("/updateOrder", hd.UpdateOrderHandler)
	http.HandleFunc("/deleteOrder", hd.DeleteOrderHandler)
	// ----------------- ITEMS
	http.HandleFunc("/loadItemsByOrderId", hd.LoadItemsByOrderId)
	http.HandleFunc("/loadFeaturesByRoleId", hd.LoadFeaturesByRoleId)
	http.HandleFunc("/loadRolesByActionId", hd.LoadRolesByActionId)
	http.HandleFunc("/loadActivitiesByWorkflowId", hd.LoadActivitiesByWorkflowId)
	http.HandleFunc("/loadAllowedActions", hd.LoadAllowedActions)
	http.HandleFunc("/loadAvailableFeatures", hd.LoadAvailableFeatures)
	http.HandleFunc("/executeAction", hd.ExecuteActionHandler)
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
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
