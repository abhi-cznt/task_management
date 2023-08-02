// this is go file
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "encoding/json"
	"github.com/gorilla/mux"
)

type Tasks struct {
	ID          string `json:"id"`
	TaskName    string `json:"task_name"`
	TaskDetails string `json:"taks_details"`
	Date        string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:          "1",
		TaskName:    "New Projects",
		TaskDetails: "You must lead the project and finish it",
		Date:        "2023-07-26",
	}
	tasks = append(tasks, task)

	task1 := Tasks{
		ID:          "2",
		TaskName:    "Power Projects",
		TaskDetails: "You have to hire more stuffs before deadline",
		Date:        "2023-07-26",
	}
	tasks = append(tasks, task1)

	fmt.Println(tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("I am home page")

}

func getTasks(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("I am tasks  page")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("I am task page")

	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}

	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("I am home page")

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("I am home page")

}

func updateTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("I am home page")

}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
	// log.Fatal(http.ListenAndServe(":55336", router))
}

func main() {
	allTasks()
	fmt.Println("hello ")
	handleRequest()
}
