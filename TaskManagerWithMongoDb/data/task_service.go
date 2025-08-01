package data

import (
	"context"
	"errors"
	"log"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func InitMongoDB(uri, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	collection = client.Database(dbName).Collection("tasks")
	log.Println("Connected to MongoDB!")
	return nil
}

func CloseMongoDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}
}

func CreateTask(taskInput models.TaskInput) (models.Task, error) {
	task := models.Task{
		Title:       taskInput.Title,
		Description: taskInput.Description,
		DueDate:     taskInput.DueDate,
		Status:      taskInput.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return models.Task{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid.Hex()
		return task, nil
	}

	return models.Task{}, errors.New("failed to get inserted ID")
}

func GetAllTasks() []models.Task {
	var tasks []models.Task

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return tasks
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &tasks); err != nil {
		log.Println("Error decoding tasks:", err)
	}

	return tasks
}

func GetTaskByID(id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID format")
	}

	var task models.Task
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	task.ID = objID.Hex()
	return task, nil
}

func UpdateTask(id string, taskInput models.TaskInput) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID format")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       taskInput.Title,
			"description": taskInput.Description,
			"due_date":    taskInput.DueDate,
			"status":      taskInput.Status,
			"updated_at":  time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedTask models.Task
	err = collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": objID},
		update,
		opts,
	).Decode(&updatedTask)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	updatedTask.ID = objID.Hex()
	return updatedTask, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
