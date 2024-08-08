# Task Management REST API Documentation (MongoDB Version)

## Overview

This updated version of the Task Management REST API uses MongoDB for persistent data storage. The MongoDB Go Driver is used to interact with the database, allowing CRUD operations on tasks.

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
cd tasks\'task managment api with mongo db'
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

## Endpoints

1. Get All Tasks

- Endpoint: `GET /tasks`
- Description: Retrieves a list of all tasks stored in MongoDB.
- Response:
  - Status: `200 OK`
  - Body:
  ```json
  {
      "data": [
          {
              "id": "64cd68f0988d4e034f9b013c",
              "title": "Task 1",
              "description": "Description for Task 1",
              "due_date": "2024-08-10",
              "status": "Pending"
          },
          ...
      ]
  }
  ```

2. Get Task by ID

- Endpoint: `GET /tasks/:id`
- Description: Retrieves details of a specific task by its MongoDB ID.
- Response:
  - Status: `200 OK`
  - Body:
  ```json
  {
    "data": {
      "id": "64cd68f0988d4e034f9b013c",
      "title": "Task 1",
      "description": "Description for Task 1",
      "due_date": "2024-08-10",
      "status": "Pending"
    }
  }
  ```
  - Status: `400 Bad Request`
  - Body:
  ```json
  {
    "error": "Invalid task ID"
  }
  ```
  - Status: `404 Not Found`
  - Body:
  ```json
  {
    "error": "task not found"
  }
  ```

3. Create a New Task

- Endpoint: `POST /tasks`
- Description: Creates a new task in MongoDB.
- Request:
  - Body: JSON object with task details
  ```json
  {
    "title": "New Task",
    "description": "Description for the new task",
    "due_date": "2024-08-15",
    "status": "Pending"
  }
  ```
- Response:
  - Status: `201 Created`
  - Body:
  ```json
  {
    "message": "task added successfully",
    "data": {
      "id": "64cd68f0988d4e034f9b013c",
      "title": "New Task",
      "description": "Description for the new task",
      "due_date": "2024-08-15",
      "status": "Pending"
    }
  }
  ```

4. Update a Task by ID

- Endpoint: `PUT /tasks/:id`
- Description: Updates details of a specific task by its MongoDB ID.
- Request:
  - Body: JSON object with updated task details.
- Response:
  - Status: `200 OK`
  - Body:
  ```json
  {
    "message": "Task updated successfully"
  }
  ```
  - Status: `400 Bad Request`
  - Body:
  ```json
  {
    "error": "Invalid task ID"
  }
  ```
  - Status: `404 Not Found`
  - Body:
  ```json
  {
    "error": "task not found"
  }
  ```

5. Delete a Task by ID

- Endpoint: `DELETE /tasks/:id`
- Description: Deletes a specific task by its MongoDB ID.
- Response:
  - Status: `200 OK`
  - Body:
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```
  - Status: `400 Bad Request`
  - Body:
  ```json
  {
    "error": "Invalid task ID"
  }
  ```
  - Status: `404 Not Found`
  - Body:
  ```json
  {
    "error": "task not found"
  }
  ```

## Testing

Use Postman to test each endpoint. Additionally, you can verify data correctness by querying MongoDB directly using the MongoDB shell or MongoDB Compass.
