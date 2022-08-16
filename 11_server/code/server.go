package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Home!")
}

type Todo struct {
	Title, Content string
}

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func getTodos(writer http.ResponseWriter, request *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}
	err = t.Execute(writer, pageVariables)

}

func addTodos(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error: ", err)
	}

	todo := Todo{
		Title:   request.FormValue("title"),
		Content: request.FormValue("content"),
	}
	todos = append(todos, todo)
	log.Println(todos)
	http.Redirect(writer, request, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodos)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
