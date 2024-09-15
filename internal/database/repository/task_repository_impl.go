package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (repository *taskRepositoryImpl) ShowTasks(ctx context.Context) ([]entity.Task, error) {
	sqlCommand := "SELECT id, description, created_at, is_done FROM task"
	rows, err := repository.DB.QueryContext(ctx, sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.IsDone)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *taskRepositoryImpl) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	sqlCommand := "UPDATE task SET description = ?, is_done = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, sqlCommand, task.Description, task.IsDone, task.Id)
	if err != nil {
		return task, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Printf("Error: Task with ID %d not updated. It could be that the task with that ID was not found or was not updated.\n", task.Id)
		return task, err
	} else {
		fmt.Printf("Data with Id %v successfully updated.\n", task.Id)
		return task, nil
	}

}
