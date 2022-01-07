package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RequestPayload to hold the request data
type RequestPayload struct {
	Text string `json:"text"`
}

type RequestHandler struct{}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (rh *RequestHandler) Register(r *mux.Router) {

	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/words", rh.getWords).Methods(http.MethodGet)
}

func (rh *RequestHandler) getWords(w http.ResponseWriter, r *http.Request) {
	var data RequestPayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
