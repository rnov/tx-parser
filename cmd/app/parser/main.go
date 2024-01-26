package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"tx-parser/internal/parser/handler"
	"tx-parser/internal/parser/service"
	"tx-parser/pkg/config"
)

func main() {
	// Get the CONFIG_PATH environment variable, default to "config/config.yaml" if not set
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// load parser config
	cfg, err := config.LoadParserConfig(configPath)
	if err != nil {
		log.Fatalf("error loading parser config: %v", err)
	}

	nSrv := service.NewService(cfg.NodeAddress, cfg.Storage.Type, cfg.IntervalCheck)
	// start service checker
	nSrv.StartChecker()

	// register parser's handler
	h := handler.NewParserHandler(nSrv)
	r := mux.NewRouter()
	r.HandleFunc("/block", h.GetCurrentBlock).Methods("GET")
	r.HandleFunc("/subscribe", h.SubscribeAddress).Methods("PUT")
	r.HandleFunc("/transactions/{address}", h.GetTransactions).Methods("GET")
	fmt.Println("starting server")
	// Fire up the server ":8080"
	log.Fatal(http.ListenAndServe(cfg.URL, r))

}
