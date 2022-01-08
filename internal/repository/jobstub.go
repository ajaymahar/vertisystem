package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/google/uuid"
)

var (
	ErrJobNotFound error = errors.New("job not found")
)

// StubRepository implementation for the JobRepo
type StubRepository struct {
	// temp in memory data store
	jobs map[string]string
	// jobs map[internal.Job]internal.JobResult

	// temp result datastre
	result map[string]internal.JobResult
}

// NewStubRepository is factory function to create new
func NewStubRepository() *StubRepository {
	return &StubRepository{
		jobs:   make(map[string]string),
		result: make(map[string]internal.JobResult),
	}
}

// Create JobRepository implementation
func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	j := internal.Job{
		ID:   getNewID(),
		Text: text,
	}
	// sr.jobs[j] = internal.JobResult{}
	sr.jobs[j.ID] = j.Text
	// fmt.Printf("%+v", sr.jobs)

	go func() {
		sr.processText(j)
	}()
	return j, nil
}

func (sr *StubRepository) Get(ctx context.Context, id string) (internal.JobResult, error) {
	result, ok := sr.result[id]
	if !ok {
		return internal.JobResult{}, ErrJobNotFound
	}
	// result, err := sr.getTopNWords()
	// if err != nil {
	// 	return internal.JobResult{}, ErrJobNotFound
	// }
	return result, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.NewString()
}

// processText is local func to process the provided text
// it will be responsable to extract the text
// get all the words from the req text
// count occurance of each word
// and store it to result datastore
func (sr *StubRepository) processText(j internal.Job) {
	words := strings.Split(j.Text, " ")
	jr := internal.NewJobResult()
	fmt.Println("jr:", jr)
	for _, word := range words {
		if len(word) > 1 {
			c := strings.Count(j.Text, word)
			jr.Frequency[word] = c
		} else {
			jr.Frequency[word] = 1
		}
	}
	fmt.Println(jr.Frequency)
	sr.result[j.ID] = *jr
}
