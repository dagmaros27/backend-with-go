# Task Management REST API Documentation

## Overview

This Task Management REST API allows users to perform CRUD (Create, Read, Update, Delete) operations on tasks. The API is developed using the Go programming language and the Gin Framework. The data is stored in-memory for simplicity.

## Endpoints

1. Get All Tasks

   - Endpoint: `GET /tasks`
   - Description: Retrieves a list of all tasks.
   - Request:
     - Method: `GET`
     - URL: `/tasks`
   - Response:
     - Status: `200 OK`
     - Body: Array of tasks (JSON)
       ```json
       [
         {
           "id": "1",
           "title": "Task 1",
           "description": "Description for Task 1",
           "due_date": "2024-08-10",
           "status": "Pending"
         },
         ...
       ]
       ```
     - Status: `200 OK` with message (JSON)
       ```json
       {
         "Message": "No task is added yet"
       }
       ```

2. Get Task by ID

   - Endpoint: `GET /tasks/:id`
   - Description: Retrieves details of a specific task by its ID.
   - Request:
     - Method: `GET`
     - URL: `/tasks/:id`
     - Parameters:
       - `id`: The ID of the task to retrieve (path parameter).
   - Response:
     - Status: `200 OK`
     - Body: Task details (JSON)
       ```json
       {
         "id": "1",
         "title": "Task 1",
         "description": "Description for Task 1",
         "due_date": "2024-08-10",
         "status": "Pending"
       }
       ```
     - Status: `404 Not Found` (JSON)
       ```json
       {
         "error": "task not found"
       }
       ```

3. Create a New Task

   - Endpoint: `POST /tasks`
   - Description: Creates a new task.
   - Request:
     - Method: `POST`
     - URL: `/tasks`
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
     - Status: `201 Created` (JSON)
       ```json
       {
         "message": "task added successfully"
       }
       ```
     - Status: `400 Bad Request` (JSON)
       ```json
       {
         "error": "Bad request message"
       }
       ```

4. Update a Task by ID

   - Endpoint: `PUT /tasks/:id`
   - Description: Updates details of a specific task by its ID.
   - Request:
     - Method: `PUT`
     - URL: `/tasks/:id`
     - Parameters:
       - `id`: The ID of the task to update (path parameter).
     - Body: JSON object with updated task details
       ```json
       {
         "title": "Updated Task Title",
         "description": "Updated Description",
         "due_date": "2024-08-20",
         "status": "Completed"
       }
       ```
   - Response:
     - Status: `200 OK` (JSON)
       ```json
       {
         "message": "Task updated successfully"
       }
       ```
     - Status: `400 Bad Request` (JSON)
       ```json
       {
         "error": "Bad request message"
       }
       ```
     - Status: `404 Not Found` (JSON)
       ```json
       {
         "error": "task not found"
       }
       ```

5. Delete a Task by ID
   - Endpoint: `DELETE /tasks/:id`
   - Description: Deletes a specific task by its ID.
   - Request:
     - Method: `DELETE`
     - URL: `/tasks/:id`
     - Parameters:
       - `id`: The ID of the task to delete (path parameter).
   - Response:
     - Status: `200 OK` (JSON)
       ```json
       {
         "message": "Task deleted successfully"
       }
       ```
     - Status: `404 Not Found` (JSON)
       ```json
       {
         "error": "task not found"
       }
       ```

## Testing

Use Postman to test each endpoint of the Task Management API by following these steps:

1. Create a new collection in Postman.
2. Add requests for each endpoint with the appropriate method, URL, and request body (if applicable).
3. Send requests to verify the responses match the expected output as documented.

## Running the API

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run the API server:
   ```bash
   go run main.go
   ```
   The server will start at `localhost:3000`. Use the base URL `http://localhost:3000` to interact with the API.

This documentation provides a comprehensive guide to understanding and using the Task Management REST API. Ensure that you follow the instructions and specifications closely for successful implementation and testing.

## Postman Documentation

- The postman documentation for this API is available [here](https://documenter.getpostman.com/view/25928149/2sA3rzJCSf#53f19b62-1b16-450a-b0ca-75427dba206f).
