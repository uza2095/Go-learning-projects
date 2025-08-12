# Task Manager API Documentation (MongoDB)

## Authentication & Authorization

### Register
- **POST** `/api/register`
- Request: `{ "username": "string", "password": "string" }`
- Response: `{ "id": string, "username": "string", "role": "user|admin" }`

### Login
- **POST** `/api/login`
- Request: `{ "username": "string", "password": "string" }`
- Response: `{ "token": "JWT token" }`

### Promote User (Admin Only)
- **POST** `/api/promote`
- Header: `Authorization: Bearer <token>`
- Request: `{ "username": "string" }`
- Response: `{ "message": "User promoted to admin" }`

## Tasks
- **GET** `/api/tasks/` - List all tasks (auth required)
- **GET** `/api/tasks/:id` - Get task by ID (auth required)
- **POST** `/api/tasks/` - Create task (admin only)
- **PUT** `/api/tasks/:id` - Update task (admin only)
- **DELETE** `/api/tasks/:id` - Delete task (admin only)

## Usage
- Register a user, then login to receive a JWT token.
- Use the token in the `Authorization` header for all protected endpoints.
- Only admins can create, update, delete tasks, and promote users.
- The first registered user is automatically an admin.

## Security
- Passwords are hashed using bcrypt.
- JWT tokens are used for authentication and must be included in the `Authorization` header.
