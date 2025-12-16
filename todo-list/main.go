package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

const fileName = "todolist.json"

func loadTodos() ([]Todo, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var todos []Todo

	// We pass '&todos' (memory address) so Unmarshal can fill the slice with data.
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func saveTodos(todos []Todo) error {
	// MarshalIndent formats the JSON with newlines and 2 spaces so it is readable by humans in the text file.
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func main() {
	// os.Args holds all command line arguments.
	if len(os.Args) < 2 {
		fmt.Println("Subcommands: 'add', 'list', 'done'")
		os.Exit(1)
	}

	command := os.Args[1]
	todos, _ := loadTodos()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add '<TODO>'")
			return
		}
		title := os.Args[2]

		newTodo := Todo{
			ID:     len(todos) + 1,
			Title:  title,
			Status: false,
		}

		todos = append(todos, newTodo)
		saveTodos(todos)
		fmt.Println("Added:", newTodo.Title)

	case "list":
		fmt.Println("-- LIST ---")
		for _, t := range todos {
			status := "[ ]"
			if t.Status {
				status = "[x]"
			}
			fmt.Printf("%s %d: %s\n", status, t.ID, t.Title)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go done <ID>")
		}

		// Convert "1" to int 1.
		id, _ := strconv.Atoi(os.Args[2])

		// We use 'i' (index) to modify the slice directly.
		// If we tried 't.Status = true', it would only modify the local copy 't'
		for i, t := range todos {
			if t.ID == id {
				todos[i].Status = true
				saveTodos(todos)
				fmt.Println("Done:", t.Title)
				return
			}
		}
		fmt.Println("Task not found")

	default:
		fmt.Println("Unknown command")
	}
}
