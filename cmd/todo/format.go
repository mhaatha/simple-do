package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Menu() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the CLI To-Do List App!")
	fmt.Println("----------------------------------")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Add a Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Mark Task as Done")
	fmt.Println("4. Delete a Task")
	fmt.Println("5. Exit")
	fmt.Println("----------------------------------")

	fmt.Print("Enter your choice: ")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	return scanner.Text()
}
