package main

import (
	"os"
	"os/signal"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/crackerplace/ems/store"
	"github.com/gorilla/mux"
)

// Config provides basic server configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	repo := store.NewRepo()
	r.HandleFunc("/departments", createDepartmentHandler(repo)).Methods("POST")
	r.HandleFunc("/employees", createEmployeeHandler(repo)).Methods("POST")
	r.HandleFunc("/employees/{name}/{departmentName}", getEmployeesHandler(repo)).Methods("GET")
	return r
}

func main() {
	serverCfg := Config{
		Host:        "localhost:8080",
		ReadTimeout: 5 * time.Second,
	}
	server := Start(newRouter(), serverCfg)
	defer server.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Info("shutting down")
}
