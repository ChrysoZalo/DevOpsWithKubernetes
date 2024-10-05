package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("layout.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ListTodo(tmpl, w, r)
	}).Methods("GET")

	port := os.Getenv("PORT")
	// port := "80" For testing
	fmt.Printf("Server started in port %s\n", port)
	http.ListenAndServe(":"+port, r)
}

func ListTodo(t *template.Template, w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
