package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var tasks = []*Task{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", health)
	http.HandleFunc("/tasks", handleTasks)

	log.Print("starting server ...")
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			nil))
}

func respondOK(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(body)
	w.Write(js)
}

func health(w http.ResponseWriter, r *http.Request) {
	respondOK(w, map[string]any{"message": "it's alive!"})
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		postTask(w, r)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	respondOK(w, tasks)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = len(tasks) + 1
	tasks = append(tasks, &t)
	respondOK(w, t)
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
