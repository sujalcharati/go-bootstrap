package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var (
	Task   []task
	taskID int = 1
)

func updateTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, " invalid request method", http.StatusBadRequest)
		return
	}

	id := r.URL.Path[len("/tasks/"):]

	for i, t := range Task {
		if t.Id == id {

			var updated_task task

			err := json.NewDecoder(r.Body).Decode(&updated_task)

			if err != nil {
				http.Error(w, " invalid body", http.StatusBadRequest)
				return
			}
			Task[i].Title = updated_task.Title
			Task[i].Description = updated_task.Description
			Task[i].Status = updated_task.Status

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Task[i])
			return

		}
	}
	http.Error(w, "task not found", http.StatusNotFound)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, " invalid request", http.StatusBadRequest)
		return
	}

	id := r.URL.Path[len("/gettasks/"):]

	for _, task := range Task {
		if task.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "task not found", http.StatusNotFound)
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "invalid req method", http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Task)

}

func createNewTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusNotAcceptable)
		return

	}
	var newTask task
	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if newTask.Title == "" || newTask.Description == "" {
		http.Error(w, "write down the title and desc", http.StatusBadRequest)
		return
	}

	taskID++
	newTask.Id = fmt.Sprintf("%d", taskID)
	Task = append(Task, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)

}

func main() {

	http.HandleFunc("/tasks", createNewTask)
	http.HandleFunc("/gettasks", getTasks)
	http.HandleFunc("/gettasks/", getTaskByID)
	http.HandleFunc("/tasks/", updateTask)
	fmt.Print(" server is running on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
