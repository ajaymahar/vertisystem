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
	jobs map[string]string
}

// NewStubRepository is factory function to create new
func NewStubRepository() *StubRepository {
	return &StubRepository{
		jobs: make(map[string]string),
	}
}

func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	id := getNewID()
	fmt.Println("id", id)
	sr.jobs[id] = text
	fmt.Printf("%#v", sr.jobs)
	return internal.Job{
		ID:   id,
		Text: text,
	}, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.NewString()
}
