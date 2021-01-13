package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id int `json:"id"`
	Body string `json:"body"`
	Completed bool `json:"completed"` 
}

type Todos []Todo

func allTodos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Id:1, Body:"Take out the trash", Completed:false },
		Todo{Id: 2, Body: "Clean the dishes",	Completed: false },
		Todo{Id: 3,	Body: "Walk the dog", Completed: false },
	}

	enableCors(&w)

	fmt.Println("Endpoint Hit: All Todos Endpoint")
	json.NewEncoder(w).Encode((todos))
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/todos", allTodos).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	handleRequests()
}