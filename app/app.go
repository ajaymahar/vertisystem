package app

import (
	"log"
	"os"

	"github.com/ajaymahar/vertisystem/internal/rest"
	"github.com/gorilla/mux"
)

func StartService() {
	// logger
	serviceLogger := log.New(os.Stdout, "service-log ", log.LstdFlags)

	// mux router
	r := mux.NewRouter()

	// register new request handler
	rh := rest.NewRequestHandler()
	rh.Register(r)

	// create new server and inject logger and mux router
	s := NewServer(r, serviceLogger)

	// start server
	s.Start()

}
