package repository

import (
	"context"

	"github.com/mhaatha/simple-do/internal/entity"
)

type TaskRepository interface {
	Add(ctx context.Context, task entity.Task) (entity.Task, error)
}
