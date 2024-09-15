package repository

import (
	"context"
	"database/sql"

	"github.com/mhaatha/simple-do/internal/database/entity"
)

type taskRepositoryImpl struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepositoryImpl{DB: db}
}

func (repository *taskRepositoryImpl) AddTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	sqlCommand := "INSERT INTO task(description) VALUES(?)"
	result, err := repository.DB.ExecContext(ctx, sqlCommand, task.Description)
	if err != nil {
		return task, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return task, err
	}

	task.Id = uint16(id)
	return task, nil
}
