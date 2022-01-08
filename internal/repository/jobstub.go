package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/google/uuid"
)

// StubRepository implementation for the JobRepo
type StubRepository struct {
	// temp in memory data store
	jobs map[string]string
	// jobs map[internal.Job]internal.JobResult

	result map[string]internal.JobResult
}

// NewStubRepository is factory function to create new
func NewStubRepository() *StubRepository {
	return &StubRepository{
		// jobs: make(map[internal.Job]internal.JobResult),
		jobs: make(map[string]string),
	}
}

func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	j := internal.Job{
		ID:   getNewID(),
		Text: text,
	}
	// sr.jobs[j] = internal.JobResult{}
	sr.jobs[j.ID] = j.Text
	// fmt.Printf("%+v", sr.jobs)

	processText(j)
	return j, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.NewString()
}

func (sr *StubRepository) processText(j internal.Job) {
	words := strings.Split(j.Text, " ")
	fmt.Println("all words here: ", words)
	for _, word := range words {
		var jr internal.JobResult
		c := strings.Count(words, word)
		sr.result[j.ID] = jr

	}
}
