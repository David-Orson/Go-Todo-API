package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Todo struct {
	Id 					int 			`json:"id"`
	Body 				string 		`json:"body"`
	Completed 	bool 			`json:"completed"` 
}

var todos []Todo

type Todos []Todo

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://david:OrsonDC@localhost/todos?sslmode=disable")

	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		fmt.Println("y")
		panic(err)
	}
	fmt.Println("connected to database.")

	rows, err := db.Query("SELECT * FROM todos;")
	if err != nil {
		fmt.Println("x")
		panic(err)
	}
	// close to free the connection to the pool for other use.
	defer rows.Close()
}
	

func getTodos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	rows, err := db.Query("SELECT * FROM todos;")
  if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Body, &todo.Completed)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err !=nil {
		http.Error(w, http.StatusText(500), 500)
	}
	

	fmt.Println("Endpoint Hit: Get Todos Endpoint")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((todos))
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	/* todo.Id = rand.Intn(10000000) */ // postgres generates serial key
	
	_, err := db.Exec("INSERT INTO todos (Body, completed) VALUES ($1, $2);", todo.Body, todo.Completed)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Println("Fetched Todos")
	w.Header().Set("Content-Type", "application/json")
	getTodos(w, r)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	/* for index, item := range todos {
		if strconv.Itoa(item.Id) == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	} */

	_, err := db.Exec("DELETE FROM todos WHERE id=$1;", params["id"])
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Println("Endpoint Hit: Delete Todo Endpoint")
	w.Header().Set("Content-Type", "application/json")
	getTodos(w, r)
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	router := mux.NewRouter()

	/* todos = append(todos, Todo{Id: 1, Body:"Take out the trash", Completed:false })
	todos = append(todos, Todo{Id: 2, Body: "Clean the dishes",	Completed: false })
	todos = append(todos, Todo{Id: 3,	Body: "Walk the dog", Completed: false }) */
	

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todos", addTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/api/todos", handleOptions).Methods("OPTIONS")
	router.HandleFunc("/api/todos/{id}", handleOptions).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
}



	

func main() {
	handleRequests()
}

// postgresql://postgres:OrsonDC@localhost:5432/todos