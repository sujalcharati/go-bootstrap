package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type task struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
}


var (
	Task []task
	taskID int = 1
)

func createNewTask(w http.ResponseWriter,r *http.Request){
      
	if r.Method != http.MethodPost {
		http.Error(w,"invalid request method",http.StatusNotAcceptable)
		return

	} 
      var newTask task
	  err := json.NewDecoder(r.Body).Decode(&newTask)

	  if err !=nil {
          http.Error(w,"invalid body",http.StatusBadRequest)
		  return
	  }

	  if newTask.Title == "" || newTask.Description == ""{
		http.Error(w, "write down the title and desc",http.StatusBadRequest)
		return
	  } 


	  taskID++
	  newTask.Id = fmt.Sprintf("%d", taskID)
	  taskID++
	  Task = append(Task, newTask)
	  
	  w.Header().Set("Content-Type","application/json")
	  json.NewEncoder(w).Encode(newTask)

}

func main(){

	http.HandleFunc("/tasks",createNewTask)
	fmt.Print(" server is running on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
