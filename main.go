package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Body string `json:"Body"`
	Completed bool `json:"Completed"` 
}

type Todos []Todo

func allTodos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Body:"Take out the trash", Completed:true },
	}

	fmt.Println("Endpoint Hit: All Todos Endpoint")
	json.NewEncoder(w).Encode((todos))
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/todos", allTodos)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}