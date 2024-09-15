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

func ShowTasks() {
	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"

	taskRepository := repository.NewTaskRepository(database.GetConnection())

	ctx := context.Background()

	result, err := taskRepository.ShowTasks(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--------------------------------------------------------------------------------")
	for _, task := range result {
		fmt.Println("Id:", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Created At:", task.CreatedAt.Format("2006-01-02 15:04:05"))
		if task.IsDone {
			fmt.Println("Completed:", string(green), task.IsDone, string(reset))
		} else {
			fmt.Println("Completed:", string(red), task.IsDone, string(reset))
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}
}
