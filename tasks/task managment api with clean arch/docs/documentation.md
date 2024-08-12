# Task Management REST API Documentation (with Authentication and Authorization)

## Overview

This updated version of the Task Management REST API uses MongoDB for persistent data storage with authentication. The MongoDB Go Driver is used to interact with the database, allowing CRUD operations on tasks. The API also includes user management endpoints for user registration, login, and role-based access control. JWT tokens are used for authentication and authorization.

## MongoDB Integration

### Configuration

MongoDB Connection String:
The API connects to a MongoDB instance running locally or in the cloud. The connection string is configured in the `dbInit` function within the `data/task_service.go` file. (If you are using mongoDB atlas or a cloud provider, you can replace the connection string with the one provided by the cloud provider)

```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
```

Database and Collection:

- Database: `task_manager`
- Collection: `tasks`

### MongoDB Installation

- Install MongoDB locally or use a cloud provider like MongoDB Atlas.
- Install the MongoDB Go Driver:

```bash
go get go.mongodb.org/mongo-driver/mongo
```

### Running the API

- Ensure MongoDB is running on your machine or accessible via a cloud service.
- Run the API server

To run the API server, follow these steps:

1. Clone the repository to your local machine:

```bash
git clone https://github.com/dagmaros27/backend-with-go.git
```

2. Change into the project directory:

```bash
cd tasks/'task managment api with mongo db'
```

3. Install the required dependencies:

```bash
go mod download
```

4. Start the API server:

```bash
go run main.go
```

The API server will be up and running at http://localhost:3000.

API Endpoints

1. User Management

   Register a New User

   Endpoint: POST /register

   Description: Create a new user account with a unique username and password.
   Request Body:

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

Responses:

- 201 Created: User registration successful.
- 400 Bad Request: Invalid input data or username already exists.

Login
Endpoint: POST /login
Description: Authenticate the user and generate a JWT token upon successful login.
Request Body:

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

Responses:

- 200 OK: Login successful. Returns a JWT token.
- 401 Unauthorized: Invalid username or password.

Promote User to Admin (Admin Only)
Endpoint: POST /promote
Description: Promote an existing user to an admin role. Only admins can perform this operation.
Headers: Authorization: Bearer <JWT token>
Request Body:

```json
{
  "username": "user_to_promote"
}
```

Responses:

- 200 OK: User promotion successful.
- 403 Forbidden: Unauthorized access. Only admins can promote users.

2. Task Management
   Create a Task (Admin Only)
   Endpoint: POST /tasks
   Description: Create a new task. Only admins can create tasks.
   Headers: Authorization: Bearer <JWT token>
   Request Body:

```json
{
  "title": "Task title",
  "description": "Task description"
}
```

Responses:

- 201 Created: Task created successfully.
- 403 Forbidden: Unauthorized access. Only admins can create tasks.

Update a Task (Admin Only)
Endpoint: PUT /tasks/:id
Description: Update an existing task. Only admins can update tasks.
Headers: Authorization: Bearer <JWT token>
Request Body:

```json
{
  "title": "Updated task title",
  "description": "Updated task description"
}
```

Responses:

- 200 OK: Task updated successfully.
- 403 Forbidden: Unauthorized access. Only admins can update tasks.

Delete a Task (Admin Only)
Endpoint: DELETE /tasks/:id
Description: Delete an existing task. Only admins can delete tasks.
Headers: Authorization: Bearer <JWT token>
Responses:

- 200 OK: Task deleted successfully.
- 403 Forbidden: Unauthorized access. Only admins can delete tasks.

Retrieve All Tasks
Endpoint: GET /tasks
Description: Retrieve a list of all tasks. Both admins and regular users can access this endpoint.
Headers: Authorization: Bearer <JWT token>
Responses:

- 200 OK: Returns a list of all tasks.

Retrieve a Task by ID
Endpoint: GET /tasks/:id
Description: Retrieve a task by its ID. Both admins and regular users can access this endpoint.
Headers: Authorization: Bearer <JWT token>
Responses:

- 200 OK: Returns the task details.
- 404 Not Found: Task with the specified ID not found.

Authentication & Authorization
JWT Token
After successful login, the server generates a JWT token. This token must be included in the Authorization header of subsequent requests to protected endpoints.
Format: Authorization: Bearer <JWT token>

User Roles
Admin: Has full access to all endpoints, including creating, updating, and deleting tasks.
Regular User: Can only retrieve tasks.

Middleware
JWT tokens are validated through middleware before granting access to protected routes. If the token is invalid or expired, the request is denied with a 401 Unauthorized response.
Admin routes are further protected by role-checking middleware, returning a 403 Forbidden response for unauthorized users.

Security
Password Storage: Passwords are hashed using a secure hashing algorithm (e.g., bcrypt) before being stored in the database.
Token Security: JWT tokens are signed with a secret key, ensuring their authenticity.

Testing
Use Postman or a similar tool to test the API endpoints.
Test both with and without valid JWT tokens to verify the proper enforcement of authentication and authorization.
Attempt to access admin-protected endpoints as a regular user to confirm that access control rules are correctly implemented.

Example Usage
Register a New User

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username": "newuser", "password": "password123"}'
```

Login

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "newuser", "password": "password123"}'
```

Create a Task (Admin Only)

```bash
curl -X POST http://localhost:8080/tasks \
-H "Authorization: Bearer <JWT token>" \
-H "Content-Type: application/json" \
-d '{"title": "New Task", "description": "Task description"}'
```

## Testing

Use Postman to test each endpoint. Additionally, you can verify data correctness by querying MongoDB directly using the MongoDB shell or MongoDB Compass.

## Postman documentation

The complete postman documentation can be found [here](https://documenter.getpostman.com/view/25928149/2sA3s3HB55)
