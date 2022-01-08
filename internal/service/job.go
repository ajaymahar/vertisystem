package service

import (
	"context"
	"fmt"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/google/uuid"
)

var jobs map[string]string

func init() {
	// temp in memory data store
	jobs = make(map[string]string)

}

// JobRepository defines the datastore handling job records
// JobRepository Port
type JobRepository interface {
	Create(context.Context, string) (internal.Job, error)
}

// StubRepository implementation for the JobRepo
type StubRepository struct {
	repo JobRepository
}

// NewStubRepository is factory function to create new
func NewStubRepository(repo JobRepository) *StubRepository {
	return &StubRepository{
		repo: repo,
	}
}

func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	id := getNewID()
	fmt.Println("id", id)
	jobs[id] = text
	return internal.Job{
		ID:   id,
		Text: text,
	}, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.New()

}
