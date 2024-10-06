package api

import (
	"github.com/chinmay211997/RiskManagerApp/controllers"
	"github.com/gorilla/mux"
)

func API() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(Health, controllers.GetHealth).Methods("GET")
	router.HandleFunc(GetPreparedURLPath(GetRisksList), controllers.GetRisksList).Methods("GET")
	router.HandleFunc(GetPreparedURLPath(GetRisksById), controllers.GetRisksById).Methods("GET")
	router.HandleFunc(GetPreparedURLPath(CreateRisk), controllers.CreateRisk).Methods("POST")
	return router
}
