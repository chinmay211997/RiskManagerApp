package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chinmay211997/RiskManagerApp/api"
	"github.com/chinmay211997/RiskManagerApp/config"
	"github.com/chinmay211997/RiskManagerApp/utils"
)

func main() {

	utils.Logger.Info(fmt.Sprintf("Starting %s", config.AppName))
	srv := &http.Server{
		Handler:      api.API(),
		Addr:         fmt.Sprintf("%s:%s", config.HOST, config.PORT),
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	utils.Logger.Info(fmt.Sprintf("Started %s on %s:%s", config.AppName, config.HOST, config.PORT))

	log.Fatal(srv.ListenAndServe())
}
