package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ajaymahar/vertisystem/internal"
)

// JobRepository defines the datastore handling job records
// JobRepository Port
type JobRepository interface {
	Create(context.Context, string) (internal.Job, error)
}

// JobService which will implement the service port
type DefaultJobService struct {
	repo JobRepository
}

// NewDefaultJobService is factory function to create new NewDefaultJobService
func NewDefaultJobService(repo JobRepository) DefaultJobService {
	return DefaultJobService{repo: repo}
}

func (djs DefaultJobService) Create(ctx context.Context, text string) (internal.Job, error) {
	job, err := djs.repo.Create(ctx, text)
	if err != nil {
		log.Println("defaultJobService: Create: repo.Create", err)
		return internal.Job{}, fmt.Errorf("repo.create: %w", err)
	}
	return job, nil
}
