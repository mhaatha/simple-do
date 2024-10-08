package repository

import (
	"context"

	"github.com/mhaatha/simple-do/internal/database/entity"
)

type TaskRepository interface {
	AddTask(ctx context.Context, task entity.Task) (entity.Task, error)
	ShowTasks(ctx context.Context) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	DeleteTask(ctx context.Context, task entity.Task) error
}
