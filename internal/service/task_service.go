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

	counter := 0

	for _, task := range result {
		if counter == 0 {
			fmt.Println("--------------------------------------------------------------------------------")
		}
		fmt.Println("Id:", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Created At:", task.CreatedAt.Format("2006-01-02 15:04:05"))
		if task.IsDone {
			fmt.Println("Completed:", string(green), task.IsDone, string(reset))
		} else {
			fmt.Println("Completed:", string(red), task.IsDone, string(reset))
		}
		fmt.Println("--------------------------------------------------------------------------------")
		counter++
	}
	if counter == 0 {
		fmt.Println("Task is empty. You can add task using 'simple-do add --help' for more information.")
	}
	counter = 0
}

func UpdateTask(id uint16, description string, isDone bool) {
	taskRepository := repository.NewTaskRepository(database.GetConnection())

	ctx := context.Background()
	task := entity.Task{
		Id:          id,
		Description: description,
		IsDone:      isDone,
	}

	_, err := taskRepository.UpdateTask(ctx, task)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteTask(id uint16) {
	taskRepository := repository.NewTaskRepository(database.GetConnection())

	ctx := context.Background()
	task := entity.Task{
		Id: id,
	}

	err := taskRepository.DeleteTask(ctx, task)
	if err != nil {
		fmt.Println(err)
	}
}
