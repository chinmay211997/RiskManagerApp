package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/chinmay211997/RiskManagerApp/config"
)

func GetHealth(w http.ResponseWriter, req *http.Request) {
	response := struct {
		AppName string
	}{
		AppName: config.AppName,
	}
	json.NewEncoder(w).Encode(response)
}
