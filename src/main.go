// this is go file
package main

import (
	"fmt"
	"log"
	"net/http"
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

func handleRequest(){
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	fmt.Println("hello ")
}
