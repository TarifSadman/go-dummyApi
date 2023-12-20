package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Name     string   `json:"name"`
	Contents []string `json:"contents"`
}

type Student struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}

var students = []Student{
	{
		ID:   1,
		Name: "John Doe",
		Courses: []Course{
			{"C", []string{"Introduction to C", "Advanced C Programming"}},
			{"C++", []string{"Object-Oriented Programming in C++", "STL"}},
			{"C#", []string{"C# Fundamentals", "ASP.NET Core"}},
			{"Java", []string{"Java Basics", "Java EE"}},
			{"Python", []string{"Python Fundamentals", "Django"}},
			{"JavaScript", []string{"JavaScript Basics", "Node.js"}},
			{"Go", []string{"Introduction to Go", "Web Development in Go"}},
		},
	},
	{
		ID:   2,
		Name: "Jane Smith",
		Courses: []Course{
			{"C++", []string{"C++ Programming Basics", "Advanced C++"}},
			{"Java", []string{"Java for Beginners", "JavaFX"}},
			{"Python", []string{"Python for Data Science", "Flask"}},
			{"JavaScript", []string{"Modern JavaScript", "React"}},
			{"Go", []string{"Go for Beginners", "Web Development in Go"}},
		},
	}
}

func main() {
	http.HandleFunc("/students", getStudentsHandler)
	http.HandleFunc("/students/", getStudentHandler)

	port := 8080
	fmt.Printf("API Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/students/"):]
	var foundStudent *Student
	for _, student := range students {
		if fmt.Sprintf("%d", student.ID) == id {
			foundStudent = &student
			break
		}
	}
	if foundStudent != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundStudent)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Student with ID %s not found", id)
	}
}
