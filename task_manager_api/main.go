package main

import (
	"log"
	"task_manager/data"
	"task_manager/models"
	"task_manager/router"
	"time"
)

func initializeSampleTasks() {

	sampleTasks := []models.TaskInput{
		{
			Title:       "Complete API Development",
			Description: "Finish the Task Manager API endpoints",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "in_progress",
		},
		{
			Title:       "Write Documentation",
			Description: "Create API documentation for users",
			DueDate:     time.Now().Add(72 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Deploy to Test Server",
			Description: "Setup CI/CD pipeline for testing",
			DueDate:     time.Now().Add(168 * time.Hour),
			Status:      "not_started",
		},
	}

	for _, task := range sampleTasks {
		if _, err := data.CreateTask(task); err != nil {
			log.Printf("Failed to initialize task '%s': %v", task.Title, err)
		}
	}
}

func main() {
	initializeSampleTasks()

	r := router.SetupRouter()
	log.Println("Server running on :8080")
	log.Println("Initialized with 3 sample tasks")
	r.Run(":8080")
}
