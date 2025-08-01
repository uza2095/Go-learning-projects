package main

import (
	"log"
	"os"
	"task_manager/data"
	"task_manager/models"
	"task_manager/router"
	"time"

	"github.com/joho/godotenv"
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
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// MongoDB configuration
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGODB_DBNAME")
	if dbName == "" {
		dbName = "taskmanager"
	}

	// Initialize MongoDB
	if err := data.InitMongoDB(mongoURI, dbName); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer data.CloseMongoDB()

	// Initialize sample data
	initializeSampleTasks()

	// Start server
	r := router.SetupRouter()
	log.Println("Server running on :8080")
	log.Println("Connected to MongoDB and initialized with sample tasks")
	r.Run(":8080")
}
