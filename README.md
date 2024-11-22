# Go CRUD API: A Personal Learning Project

![Project Image](images/go_crud_project.png)

## Overview
This project is a simple CRUD (Create, Read, Update, Delete) API built using Go (Golang) and PostgreSQL. It was created as a hands-on learning project to understand basic concepts of Go programming, web development, and working with relational databases. The goal was to explore how to build and connect a Go server with a PostgreSQL database to perform CRUD operations.

The project focuses on creating an API server that interacts with a PostgreSQL database to perform operations like creating users, retrieving user details, and maintaining basic server health checks.

The key sections of the project include:
- Setting up a PostgreSQL connection
- Building RESTful APIs with Go
- Creating a health check endpoint
- Creating and retrieving users

## Table of Contents
- Installation
- Setting Up the Database

  - Update User
- Usage
- Conclusion
- License

## Installation
To run this project locally, clone the repository and install the required dependencies.

```bash
git clone https://github.com/your-username/go-crud-api
cd go-crud-api
```

### Prerequisites
- Go installed (version 1.18 or higher).
- PostgreSQL installed and running.
- Properly configured `psql` role with username and password.

Install the required Go modules:
```bash
go mod tidy
```

## Setting Up the Database
Ensure that PostgreSQL is running locally and create a database named `mydb`. You also need to create a user role (`myuser`) with the correct permissions.

Connect to your PostgreSQL server and execute the following commands to set up the database and user role:

```sql
CREATE DATABASE mydb;
CREATE ROLE myuser WITH LOGIN PASSWORD 'mypassword';
ALTER ROLE myuser CREATEDB;
```

The Go application will automatically create the required `users` table if it does not exist.

![Database Setup](images/database_setup.png)

### Update User
- **URL**: `http://localhost:8080/user`
- **Method**: PUT
- **Headers**: `Content-Type: application/json`
- **Body**:
  ```json
  {
    "username": "johndoe",
    "password": "newpassword123",
    "age": 35
  }
  ```
- **Description**: Updates an existing user's information with the provided details.
- **Response**: Returns `200 OK` and a success message if the user is updated successfully.

![Update User Image](images/update_user.png)

## CRUD Endpoints

### Health Check
- **URL**: `http://localhost:8080/healthcheck`
- **Method**: GET
- **Description**: Checks the health of the server and database connection.
- **Response**: Returns `OK` if the server is healthy.

![Health Check Image](images/health_check.png)

### Create User
- **URL**: `http://localhost:8080/user`
- **Method**: POST
- **Headers**: `Content-Type: application/json`
- **Body**:
  ```json
  {
    "username": "johndoe",
    "password": "password123",
    "age": 30
  }
  ```
- **Description**: Creates a new user with the provided details.
- **Response**: Returns `201 Created` and a success message if the user is created successfully.

![Create User Image](images/create_user.png)

### Get User
- **URL**: `http://localhost:8080/getuser`
- **Method**: GET
- **Authentication**: Basic Auth (username and password of the user you wish to retrieve)
- **Description**: Retrieves user details if the correct credentials are provided.
- **Response**: Returns `200 OK` and user details if authenticated successfully.

![Get User Image](images/get_user.png)

## Usage
To run this project locally:

1. **Navigate to the Project Directory**:
   ```bash
   cd go-crud-api
   ```
2. **Initialize the Go Module** (if you haven't already):
   ```bash
   go mod init go-crud-api
   ```
3. **Run the Server**:
   ```bash
   go run main.go
   ```
4. **Testing the Endpoints**:
   Use Postman or cURL to send requests to the API.

   Example:
   - **Health Check**: `GET http://localhost:8080/healthcheck`
   - **Create User**: `POST http://localhost:8080/user` with JSON body `{ "username": "johndoe", "password": "password123", "age": 30 }`
   - **Get User**: `GET http://localhost:8080/getuser` using Basic Auth (`username` and `password`)

## Conclusion
This Go CRUD API project was a learning experience in building a backend server with Go and PostgreSQL. It demonstrates how to create a RESTful API, connect to a relational database, and perform basic CRUD operations in a scalable and reusable manner.

This project serves as a foundation for anyone interested in learning Go web development and working with relational databases.

![Project Conclusion Image](images/project_conclusion.png)

## License
This project is open source and available under the MIT License.

