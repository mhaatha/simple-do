package repository

import (
	"context"

	"github.com/mhaatha/simple-do/internal/database/entity"
)

type TaskRepository interface {
	AddTask(ctx context.Context, task entity.Task) (entity.Task, error)
}
