# RESTful Task API with Gin

A simple REST API for task management built with Go and the Gin framework. This project demonstrates how to create a CRUD API with proper routing, JSON serialization, and API documentation using Swagger.

## Features

- Complete RESTful API implementation (GET, POST, PUT, DELETE)
- In-memory data storage with thread-safe operations
- JSON request/response handling
- Swagger integration for API documentation
- CORS support for cross-origin requests
- Proper error handling with appropriate HTTP status codes

## Installation

### Requirements

- Go 1.16 or higher
- Git

### Build and Install

```bash
# Clone the repository
git clone https://github.com/hikaru-tomizawa/restful-task-api.git
cd restful-task-api

# Install dependencies
go mod download

# Install Swagger (optional, for API documentation)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs (if Swagger is installed)
swag init

# Run the application
go run main.go
```

## Usage

Once the server is running, you can interact with the API using curl or any HTTP client.

```bash
# Get all tasks
curl -X GET http://localhost:8080/api/v1/tasks

# Create a new task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Study RESTful API development with Gin"}'

# Get a specific task (replace {id} with an actual task ID)
curl -X GET http://localhost:8080/api/v1/tasks/{id}

# Update a task
curl -X PUT http://localhost:8080/api/v1/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Study RESTful API development with Gin","status":"completed"}'

# Delete a task
curl -X DELETE http://localhost:8080/api/v1/tasks/{id}
```

Additionally, you can access the Swagger UI documentation at `http://localhost:8080/swagger/index.html`

## Project Structure

```
restful-task-api/
├── main.go                  # Application entry point
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── docs/                    # Swagger generated documentation
├── handlers/
│   └── task_handler.go      # Task-related request handlers
├── models/
│   └── task.go              # Task model definition
└── routes/
    └── routes.go            # API route configuration
```

## Learning Objectives

This project was created to learn the following Golang concepts:

- **Gin Framework**: Routing, middleware, and HTTP request/response handling
- **JSON Serialization**: Converting between Go structs and JSON
- **API Documentation**: Using Swagger to document API endpoints
- **Concurrency Safety**: Using mutexes to protect shared resources
- **RESTful API Design**: Proper resource-oriented API design using HTTP methods
- **Error Handling**: Returning appropriate HTTP status codes and error messages
- **Cross-Origin Resource Sharing (CORS)**: Handling requests from different origins

## Possible Extensions

Ideas for extending this project:
- Add persistent storage using a database (SQLite, MySQL, PostgreSQL)
- Implement user authentication and authorization
- Add pagination and filtering for task listing
- Create a simple frontend UI to interact with the API
- Implement task categories or tags
- Add unit and integration tests
- Containerize the application using Docker
- Set up CI/CD pipeline

## License

MIT License

## Contributing

While this project was created for learning purposes, improvement suggestions and bug reports are welcome. Feel free to create an Issue or Pull Request.

## Author

Hikaru Tomizawa
