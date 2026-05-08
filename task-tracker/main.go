package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/service"
	"task-tracker/model"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: app <arg1> <arg2> ...")
		return
	}

	tracker, err := service.NewTracker()
	if err != nil {
		fmt.Println("Error initializing tracker:", err)
		return
	}

	switch os.Args[1] {
	case "add":
		tracker.AddTask(model.Task{
			Description: os.Args[2],
			Status:      "todo",
			UpdatedAt:   time.Now().Format(time.RFC3339),
			CreatedAt:   time.Now().Format(time.RFC3339),
		})
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		task := tracker.GetTask(id)
		if task == nil {
			fmt.Println("Task not found:", id)
			return
		}

		// update description of the task and updatedAt
		task.Description = os.Args[3]
		task.UpdatedAt = time.Now().Format(time.RFC3339)

		tracker.UpdateTask(id, *task)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		tracker.DeleteTask(id)
	case "mark-in-progress":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		task := tracker.GetTask(id)
		if task == nil {
			fmt.Println("Task not found:", id)
			return
		}

		task.Status = "in-progress"
		task.UpdatedAt = time.Now().Format(time.RFC3339)
		tracker.UpdateTask(id, *task)
	case "mark-done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		task := tracker.GetTask(id)
		if task == nil {
			fmt.Println("Task not found:", id)
			return
		}

		task.Status = "done"
		task.UpdatedAt = time.Now().Format(time.RFC3339)
		tracker.UpdateTask(id, *task)
	case "list":
		if len(os.Args) > 2 {
			status := os.Args[2]
			tasks := tracker.ListTasksByStatus(status)
			for _, task := range tasks {
				fmt.Printf("ID: %d, Description: %s, Status: %s, UpdatedAt: %s, CreatedAt: %s\n",
					task.ID, task.Description, task.Status, task.UpdatedAt, task.CreatedAt)
			}
			return
		}

		tasks := tracker.ListTasks()
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, UpdatedAt: %s, CreatedAt: %s\n",
				task.ID, task.Description, task.Status, task.UpdatedAt, task.CreatedAt)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

	tracker.Save()
}
