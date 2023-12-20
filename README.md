# Go Dummy API

This is a simple RESTful API built with Go that provides information about students and their courses. The project demonstrates basic web server functionality, handling HTTP requests, and serving JSON responses.

## Prerequisites

- Make sure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/TarifSadman/go-dummyApi.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-dummyApi
   ```

3. Run the project:

   ```bash
   go run main.go
   ```

   The server will start on http://localhost:8080.

## API Endpoints

- **GET /students**: Returns a list of students with their courses.
- **GET /students/{id}**: Returns a specific student based on the provided ID.

## Example Usage

- Access the list of students: [http://localhost:8080/students](http://localhost:8080/students)
