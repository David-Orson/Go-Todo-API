package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id 					int 			`json:"id"`
	Body 				string 		`json:"body"`
	Completed 	bool 			`json:"completed"` 
}

var todos []Todo

type Todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	fmt.Println("Endpoint Hit: Get Todos Endpoint")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((todos))
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.Id = rand.Intn(10000000)
	todos = append(todos, todo)

	fmt.Println("Endpoint Hit: Add Todo Endpoint")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((todos))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)


	for index, item := range todos {
		if strconv.Itoa(item.Id) == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}

	fmt.Println("Endpoint Hit: Delete Todo Endpoint")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((todos))
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	router := mux.NewRouter()

	todos = append(todos, Todo{Id:1, Body:"Take out the trash", Completed:false })
	todos = append(todos, Todo{Id: 2, Body: "Clean the dishes",	Completed: false })
	todos = append(todos, Todo{Id: 3,	Body: "Walk the dog", Completed: false })
	

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todos", addTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	handleRequests()
}