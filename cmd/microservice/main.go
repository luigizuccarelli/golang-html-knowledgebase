package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luigizuccarelli/golang-html-knowledgebase/pkg/connectors"
	"github.com/luigizuccarelli/golang-html-knowledgebase/pkg/handlers"
	"github.com/luigizuccarelli/golang-html-knowledgebase/pkg/validator"
	"github.com/microlib/simple"
)

var (
	logger *simple.Logger
)

func startHttpServer(con connectors.Clients) *http.Server {
	srv := &http.Server{Addr: ":" + os.Getenv("SERVER_PORT")}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/search", func(w http.ResponseWriter, req *http.Request) {
		handlers.SearchHandler(w, req, con)
	}).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/v1/sys/info/isalive", handlers.IsAlive).Methods("GET")

	sh := http.StripPrefix("/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/").Handler(sh)

	http.Handle("/", r)

	if err := srv.ListenAndServe(); err != nil {
		con.Error("Httpserver: ListenAndServe() error: " + err.Error())
	}

	return srv
}

func main() {
	if os.Getenv("LOG_LEVEL") == "" {
		logger = &simple.Logger{Level: "info"}
	} else {
		logger = &simple.Logger{Level: os.Getenv("LOG_LEVEL")}
	}
	err := validator.ValidateEnvars(logger)
	if err != nil {
		os.Exit(-1)
	}
	conn := connectors.NewClientConnections(logger)
	srv := startHttpServer(conn)
	logger.Info("Starting server on port " + srv.Addr)
}
