package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Age     int     `json:"age"`
	Worth   float64 `json:"worth"`
}

var People = []Person{
	Person{Id: 3, Name: "Aaliyah", Address: "40 Inca Drive, London", Age: 12, Worth: 654321},
	Person{Id: 5, Name: "Apollo", Address: "40 Inca Drive, London", Age: 50, Worth: 7654321},
}

func rootUrl(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("the root (home) url endpoint has been hit")
	fmt.Fprintf(writer, "This REST API is all about People %s", request.URL.Path[1:])
}

func listPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the list people endpoint has been hit")
	json.NewEncoder(w).Encode(People)
}

func returnSinglePerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theId := vars["id"]
	intId, _ := strconv.Atoi(theId)
	for _, person := range People {
		if intId == person.Id {
			json.NewEncoder(writer).Encode(person)
		}
	}
}

func deletePerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theId := vars["id"]
	intId, _ := strconv.Atoi(theId)
	for _, person := range People {
		if intId == person.Id {
			json.NewEncoder(writer).Encode(person)
		}
	}
}

func createNewPerson(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var person Person
	json.Unmarshal(reqBody, &person)
	People = append(People, person)
	json.NewEncoder(writer).Encode(person)
}

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", rootUrl)
	myRouter.HandleFunc("/all", listPeople)
	myRouter.HandleFunc("/person", createNewPerson).Methods("POST")
	myRouter.HandleFunc("/person/{id}", returnSinglePerson)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
