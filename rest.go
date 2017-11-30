package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"./lib"
)

type Person struct {
	ID        string
	Firstname string
	Lastname  string
	Address   *Address
}
type Address struct {
	City  string
	State string
}

type Result struct {
	Status bool
	Result []Person
	Message string
}

var result = Result{}
var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	lib.Block{
		Try: func() {
			result.Result = people
			result.Status = true
			result.Message = "Success"
		},
		Catch: func(e lib.Exception) {
			result.Status = false
			result.Message = "Error"
		},
		Finally: func() {
			json.NewEncoder(w).Encode(result)
		},
	}.Do()
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var temp []Person
	lib.Block{
		Try: func() {
			for _, item := range people {
				if item.ID == params["id"] {
					temp = append(temp, item)
					result.Result = temp
					return
				}
			}
			result.Status = true
			result.Message = "Success"
		},
		Catch: func(e lib.Exception) {
			result.Status = false
			result.Message = "Error"
		},
		Finally: func() {
			json.NewEncoder(w).Encode(result)
		},
	}.Do()

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lib.Block{
		Try: func() {
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			people = append(people, person)

			result.Result = people
			result.Status = true
			result.Message = "Success"
		},
		Catch: func(e lib.Exception) {
			result.Status = false
			result.Message = "Error"
		},
		Finally: func() {
			json.NewEncoder(w).Encode(result)
		},
	}.Do()
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lib.Block{
		Try: func() {
			for index, item := range people {
				if item.ID == params["id"] {
					people = append(people[:index], people[index+1:]...)
					break
				}
			}
			result.Result = people
			result.Status = true
			result.Message = fmt.Sprintf("Success deleted id %s", params["id"])
		},
		Catch: func(e lib.Exception) {
			result.Status = false
			result.Message = "Error"
		},
		Finally: func() {
			json.NewEncoder(w).Encode(result)
		},
	}.Do()
}

// our main function
func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8000", router))
}