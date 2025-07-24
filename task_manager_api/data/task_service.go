package data

import (
	"errors"
	"task_manager/models"
	"time"
)

var tasks = make(map[string]models.Task)

func CreateTask(taskInput models.TaskInput) (models.Task, error) {
	id := time.Now().Format("20060102150405")
	task := models.Task{
		ID:          id,
		Title:       taskInput.Title,
		Description: taskInput.Description,
		DueDate:     taskInput.DueDate,
		Status:      taskInput.Status,
	}
	tasks[id] = task
	return task, nil
}

func GetAllTasks() []models.Task {
	var result []models.Task
	for _, task := range tasks {
		result = append(result, task)
	}
	return result
}

func GetTaskByID(id string) (models.Task, error) {
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func UpdateTask(id string, taskInput models.TaskInput) (models.Task, error) {
	if _, exists := tasks[id]; !exists {
		return models.Task{}, errors.New("task not found")
	}

	updatedTask := models.Task{
		ID:          id,
		Title:       taskInput.Title,
		Description: taskInput.Description,
		DueDate:     taskInput.DueDate,
		Status:      taskInput.Status,
	}
	tasks[id] = updatedTask
	return updatedTask, nil
}

func DeleteTask(id string) error {
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
