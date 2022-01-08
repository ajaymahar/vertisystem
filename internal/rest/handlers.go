package rest

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/ajaymahar/vertisystem/internal/repository"
	"github.com/gorilla/mux"
)

// JobService
type JobService interface {
	Create(context.Context, string) (internal.Job, error)
	Get(context.Context, string) (internal.JobResult, error)
}

// ################################
// CreateJobRequest to hold the request payload data
type CreateJobRequest struct {
	Text string `json:"text"`
}

// CreateJobResponse defines the response returned back to the client
type CreateJobResponse struct {
	Job CreateJob `json:"job"`
}

type CreateJob struct {
	ID string `json:"id"`
}

// ################################
// GetJobRequest to hold the request payload data
type GetJobRequest struct {
	ID string `json:"id"`
}

// GetJobResponse defines the response returned back to the client
type GetJobResponse struct {
	Job GetJob `json:"job"`
}

//GetJob
type GetJob struct {
	ID        string
	Frequency map[string]int // word with it's occurance
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
	sr.HandleFunc("/text", rh.createJob).Methods(http.MethodPost)
	sr.HandleFunc("/text/{id}", rh.getWords).Methods(http.MethodGet)
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
	resp := CreateJobResponse{
		Job: CreateJob{
			ID: job.ID,
		},
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
func (rh *JobHandler) getWords(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// fmt.Println("id: ", id)

	req := GetJobRequest{
		ID: id,
	}
	jobResult, err := rh.svc.Get(r.Context(), req.ID)
	if err != nil {
		// handle wrong id
		if errors.Is(err, repository.ErrJobNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// any other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println(jobResult)
	// fmt.Println("job result")

	resp := GetJobResponse{
		Job: GetJob{
			ID:        id,
			Frequency: jobResult.Frequency,
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
