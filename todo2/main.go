package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	opendb()

	http.HandleFunc("/health", health)
	http.HandleFunc("/tasks", handleTasks)

	log.Print("starting server ...")
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			nil))
}

func opendb() {

	var err error
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var server = os.Getenv("SERVER")
	db, err = sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@(%s)/tasksdb", user, password, server))
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Print("Connected!")
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
	tasks := []*Task{}
	err := db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		fmt.Println(err)
	}

	respondOK(w, tasks)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.NamedExec(`INSERT INTO tasks (description) VALUES (:description)`, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ID, _ = res.LastInsertId()
	respondOK(w, t)
}

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}
