package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

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

var mu sync.Mutex // Mutex to protect concurrent writes to the file

func main() {
	r := mux.NewRouter()

	todoRouter := r.PathPrefix("/todo").Subrouter()

	tmpl := template.Must(template.ParseFiles("layout.html"))

	todoRouter.PathPrefix("/tmp/").Handler(http.StripPrefix("/todo/tmp/", http.FileServer(http.Dir("/tmp/"))))
	todoRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		ListTodo(tmpl, w, r)
	}).Methods("GET")

	// Start the background task to download the image every 60 seconds
	go startImageFetcher()

	// Start the server
	port := "80" // For testing
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

// startImageFetcher fetches a new image every 60 seconds
func startImageFetcher() {
	ticker := time.NewTicker(60 * time.Minute)
	defer ticker.Stop()

	for {
		fetchAndSaveImage()

		<-ticker.C // Wait for the next tick (60 seconds)
	}
}

func fetchAndSaveImage() {
	fmt.Println("Fetching new image...")

	resp, err := http.Get("https://picsum.photos/200")
	if err != nil {
		fmt.Printf("Error downloading image: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Use a mutex to prevent concurrent writes to the file
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Create("/tmp/images/image.jpg")
	if err != nil {
		fmt.Printf("Error creating image file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error saving image: %v\n", err)
		return
	}

	fmt.Println("Image successfully saved.")
}
