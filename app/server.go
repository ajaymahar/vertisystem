package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// basic server type for service
type Server struct {
	router *mux.Router
	logger *log.Logger
}

// NewServer - factory function to create new server
func NewServer(r *mux.Router, l *log.Logger) *Server {
	return &Server{
		router: r,
		logger: l,
	}
}

// Start - will start the server
func (s *Server) Start() {

	// custom server to provide some timeout
	svr := &http.Server{
		Addr:              ":8080",
		Handler:           s.router,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       3 * time.Second,
		ErrorLog:          s.logger,
	}

	// create new gorotine which will listen to the server
	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			svr.ErrorLog.Fatal(err.Error())
		}
	}()

	s.logger.Println("server started ListenAndServe on port: 8080")
	// setup gracefull shutdown for the server
	// create a chanel to capture the os signal for any intruption
	sigChan := make(chan os.Signal, 1)

	// capture the os signal and notify the sigChan
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// block the call till we get the signal
	sig := <-sigChan
	s.logger.Println("received os signal: ", sig)
	// keep 30sec grace time period to do graceful shutdown the server
	// this will make sure if there is any request peding to process
	// will process it and then shutdown the server
	cxt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := svr.Shutdown(cxt)
	if err != nil {
		svr.ErrorLog.Println("error while shutdonw the server: ", err.Error())
	}
}
