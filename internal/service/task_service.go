package service

import (
	"context"
	"fmt"

	"github.com/mhaatha/simple-do/internal/database"
	"github.com/mhaatha/simple-do/internal/database/entity"
	"github.com/mhaatha/simple-do/internal/database/repository"
)

func AddTask(description string) {
	taskRepository := repository.NewTaskRepository(database.GetConnection())

	ctx := context.Background()
	task := entity.Task{
		Description: description,
	}

	result, err := taskRepository.AddTask(ctx, task)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Data inserted successfully with ID: %d\n", result.Id)
}
