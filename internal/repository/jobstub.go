package repository

import (
	"context"
	"fmt"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/google/uuid"
)

// StubRepository implementation for the JobRepo
type StubRepository struct {
	// temp in memory data store
	// jobs map[string]string
	jobs map[internal.Job]internal.JobResult
}

// NewStubRepository is factory function to create new
func NewStubRepository() *StubRepository {
	return &StubRepository{
		jobs: make(map[internal.Job]internal.JobResult),
	}
}

func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	j := internal.Job{
		ID:   getNewID(),
		Text: text,
	}
	sr.jobs[j] = internal.JobResult{}
	fmt.Printf("%+v", sr.jobs)
	return j, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.NewString()
}
