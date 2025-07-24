# Task Management API Documentation

## Overview
This is a REST API for managing tasks, built with Go and the Gin framework. The API provides endpoints for creating, reading, updating, and deleting tasks with in-memory storage.

## Base URL
```
http://localhost:8080
```

## API Endpoints

### 1. Health Check
**GET** `/health`

Check if the API is running.

**Response:**
```json
{
  "status": "healthy",
  "message": "Task Management API is running",
  "version": "1.0.0"
}
```

### 2. Get All Tasks
**GET** `/tasks` or **GET** `/api/v1/tasks`

Retrieve all tasks from the system.

**Response:**
```json
{
  "success": true,
  "message": "Tasks retrieved successfully",
  "data": [
    {
      "id": "1",
      "title": "Complete project documentation",
      "description": "Write comprehensive documentation for the task management API",
      "due_date": "2025-01-16T10:30:00Z",
      "status": "pending",
      "created_at": "2025-01-13T10:30:00Z",
      "updated_at": "2025-01-13T10:30:00Z"
    }
  ]
}
```

### 3. Get Task by ID
**GET** `/tasks/{id}` or **GET** `/api/v1/tasks/{id}`

Retrieve a specific task by its ID.

**Parameters:**
- `id` (string, required): The task ID

**Response (Success - 200):**
```json
{
  "success": true,
  "message": "Task retrieved successfully",
  "data": {
    "id": "1",
    "title": "Complete project documentation",
    "description": "Write comprehensive documentation for the task management API",
    "due_date": "2025-01-16T10:30:00Z",
    "status": "pending",
    "created_at": "2025-01-13T10:30:00Z",
    "updated_at": "2025-01-13T10:30:00Z"
  }
}
```

**Response (Not Found - 404):**
```json
{
  "success": false,
  "message": "Task not found",
  "error": "task not found"
}
```

### 4. Create Task
**POST** `/tasks` or **POST** `/api/v1/tasks`

Create a new task.

**Request Body:**
```json
{
  "title": "New Task Title",
  "description": "Task description here",
  "due_date": "2025-01-20T15:30:00Z",
  "status": "pending"
}
```

**Field Validation:**
- `title` (string, required): Task title
- `description` (string, required): Task description
- `due_date` (datetime, required): Task due date in ISO 8601 format
- `status` (string, required): Must be one of: "pending", "in_progress", "completed"

**Response (Success - 201):**
```json
{
  "success": true,
  "message": "Task created successfully",
  "data": {
    "id": "4",
    "title": "New Task Title",
    "description": "Task description here",
    "due_date": "2025-01-20T15:30:00Z",
    "status": "pending",
    "created_at": "2025-01-13T10:30:00Z",
    "updated_at": "2025-01-13T10:30:00Z"
  }
}
```

**Response (Validation Error - 400):**
```json
{
  "success": false,
  "message": "Invalid request payload",
  "error": "Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

### 5. Update Task
**PUT** `/tasks/{id}` or **PUT** `/api/v1/tasks/{id}`

Update an existing task. All fields are optional.

**Parameters:**
- `id` (string, required): The task ID

**Request Body:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2025-01-25T15:30:00Z",
  "status": "in_progress"
}
```

**Field Validation:**
- `title` (string, optional): Task title
- `description` (string, optional): Task description
- `due_date` (datetime, optional): Task due date in ISO 8601 format
- `status` (string, optional): Must be one of: "pending", "in_progress", "completed"

**Response (Success - 200):**
```json
{
  "success": true,
  "message": "Task updated successfully",
  "data": {
    "id": "1",
    "title": "Updated Task Title",
    "description": "Updated description",
    "due_date": "2025-01-25T15:30:00Z",
    "status": "in_progress",
    "created_at": "2025-01-13T10:30:00Z",
    "updated_at": "2025-01-13T10:35:00Z"
  }
}
```

**Response (Not Found - 404):**
```json
{
  "success": false,
  "message": "Failed to update task",
  "error": "task not found"
}
```

### 6. Delete Task
**DELETE** `/tasks/{id}` or **DELETE** `/api/v1/tasks/{id}`

Delete a specific task.

**Parameters:**
- `id` (string, required): The task ID

**Response (Success - 200):**
```json
{
  "success": true,
  "message": "Task deleted successfully"
}
```

**Response (Not Found - 404):**
```json
{
  "success": false,
  "message": "Failed to delete task",
  "error": "task not found"
}
```

### 7. Get Task Statistics
**GET** `/api/v1/tasks/stats`

Get statistics about tasks in the system.

**Response:**
```json
{
  "success": true,
  "message": "Task statistics retrieved successfully",
  "data": {
    "total_tasks": 5,
    "pending": 2,
    "in_progress": 2,
    "completed": 1
  }
}
```

## Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Task Status Values

The API supports three task statuses:
- `pending`: Task is waiting to be started
- `in_progress`: Task is currently being worked on
- `completed`: Task has been finished

## Testing with Postman

### Collection Structure
Create a new Postman collection named "Task Management API" with the following requests:

1. **Health Check**
   - Method: GET
   - URL: `http://localhost:8080/health`

2. **Get All Tasks**
   - Method: GET
   - URL: `http://localhost:8080/tasks`

3. **Get Task by ID**
   - Method: GET
   - URL: `http://localhost:8080/tasks/1`

4. **Create Task**
   - Method: POST
   - URL: `http://localhost:8080/tasks`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "title": "Test Task",
       "description": "This is a test task created via Postman",
       "due_date": "2025-01-20T15:30:00Z",
       "status": "pending"
     }
     ```

5. **Update Task**
   - Method: PUT
   - URL: `http://localhost:8080/tasks/1`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "status": "in_progress"
     }
     ```

6. **Delete Task**
   - Method: DELETE
   - URL: `http://localhost:8080/tasks/1`

7. **Get Task Statistics**
   - Method: GET
   - URL: `http://localhost:8080/api/v1/tasks/stats`

### Environment Variables
Create a Postman environment with:
- `base_url`: `http://localhost:8080`

## Error Handling

The API implements comprehensive error handling:

1. **Validation Errors**: Returns 400 status code with detailed validation messages
2. **Not Found Errors**: Returns 404 status code when resources don't exist
3. **ID Format Validation**: Ensures task IDs are valid integers
4. **Status Validation**: Ensures task status values are valid
5. **Required Field Validation**: Ensures all required fields are provided

## Sample Data

The API initializes with sample data for testing:
- Complete project documentation (pending)
- Code review (in_progress)
- Database optimization (pending)

## Running the API

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. The API will be available at `http://localhost:8080`

## Architecture

The API follows a clean architecture pattern:
- **main.go**: Application entry point
- **router/**: Route definitions and configuration
- **controllers/**: HTTP request handlers
- **data/**: Business logic and data operations
- **models/**: Data structures and validation

This architecture ensures separation of concerns and makes the code maintainable and testable.