package app

import (
	"log"
	"os"

	"github.com/ajaymahar/vertisystem/internal/repository"
	"github.com/ajaymahar/vertisystem/internal/rest"
	"github.com/ajaymahar/vertisystem/internal/service"
	"github.com/gorilla/mux"
)

func StartService() {
	// logger
	serviceLogger := log.New(os.Stdout, "service-log ", log.LstdFlags)

	// mux router
	r := mux.NewRouter()

	// repository
	repo := repository.NewStubRepository()

	// service
	svc := service.NewDefaultJobService(repo)

	// register new request handler
	jh := rest.NewJobHandler(svc, serviceLogger)
	jh.Register(r)

	// create new server and inject logger and mux router
	s := NewServer(r, serviceLogger)

	// start server
	s.Start()

}
