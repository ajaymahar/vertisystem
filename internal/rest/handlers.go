package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/gorilla/mux"
)

// JobService
type JobService interface {
	Create(context.Context, string) (internal.Job, error)
}

// ################################
// CreateJobRequest to hold the request payload data
type CreateJobRequest struct {
	Text string `json:"text"`
}

// CreateJobResponse defines the response returned back to the client
type CreateJobResponse struct {
	Job Job `json:"job"`
}

type Job struct {
	ID string `json:"id"`
}

// ################################

// ################################
type JobHandler struct {
	svc JobService
	l   *log.Logger
}

func NewJobHandler(svc JobService, l *log.Logger) *JobHandler {
	return &JobHandler{
		svc: svc,
		l:   l,
	}
}

// ################################

// ####################################
func (rh *JobHandler) Register(r *mux.Router) {

	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/submit", rh.createJob).Methods(http.MethodPost)
}

// ##################################

func (rh *JobHandler) createJob(w http.ResponseWriter, r *http.Request) {
	var req CreateJobRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		rh.l.Println("jobhandler: createJob: json.decode ", err)
		return
	}

	// create Service
	job, err := rh.svc.Create(r.Context(), req.Text)
	if err != nil {
		rh.l.Println("jobhandler: createJob: svc.Create ", err)
		return
	}
	fmt.Fprintln(w, job)
}

// // middleware to validate the comming request
// func (rh *JobHandler) validatePayload(next http.HandlerFunc) http.HandlerFunc {
//
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var p RequestPayload
// 		payload, err := ioutil.ReadAll(r.Body)
// 		defer r.Body.Close()
// 		if err != nil {
// 			rh.l.Println("validatePayload: readAll: ", err)
// 			return
// 		}
//
// 		// check if it's valid json data
// 		if !json.Valid(payload) {
// 			// not valid json
// 			rh.l.Println("validatePayload: json.Valid: ")
// 			return
// 		}
//
// 		if p.Data == "" {
// 			rh.l.Println("data must be provided")
// 			return
// 		}
//
// 		// call next handler
// 		next.ServeHTTP(w, r)
// 	}
// }

//
// func (rh *JobHandler) getWords(w http.ResponseWriter, r *http.Request) {
// 	var data RequestPayload
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	// pass request to next service
// 	rh.svc.Create(data)
// }
