package main

import (
	"cube.evolve.go/cube"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Person struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Age     int     `json:"age"`
	Worth   float64 `json:"worth"`
}

var People = []Person{
	Person{Id: 3, Name: "Iron Maiden", Address: "666 Acacia Avenue, Sydney", Age: 12, Worth: 654321},
	Person{Id: 5, Name: "Harry Potter", Address: "42 Privet Drive, Cambridge", Age: 50, Worth: 7654321},
}

func rootUrl(writer http.ResponseWriter, request *http.Request) {
	Info.Printf("the root (home) url endpoint has been hit - %s", request)
	fmt.Fprintf(writer, "This REST API is all about People %s", request.URL.Path[1:])
}

func listPeople(writer http.ResponseWriter, request *http.Request) {
	Info.Printf("the list people endpoint has been hit by request", request)
	_ = json.NewEncoder(writer).Encode(People)
}

func returnSinglePerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theId := vars["id"]
	intId, _ := strconv.Atoi(theId)
	for _, person := range People {
		if intId == person.Id {
			_ = json.NewEncoder(writer).Encode(person)
		}
	}
}

/*func deletePerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theId := vars["id"]
	intId, _ := strconv.Atoi(theId)
	for _, person := range People {
		if intId == person.Id {
			_ = json.NewEncoder(writer).Encode(person)
		}
	}
}
*/

func createNewPerson(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var person Person
	_ = json.Unmarshal(reqBody, &person)
	People = append(People, person)
	_ = json.NewEncoder(writer).Encode(person)
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Trace.Println("I have something standard to say")
	randomCube := cube.RandomCube()
	Info.Printf("The random cube state is %s", randomCube)
	randomCube.Revolve(cube.Xy)
	Warning.Printf("Warning is that after the %s move the state is %s", cube.Xy, randomCube)
	Info.Println("Person ID is %s and name is %s and address is %s", People[0].Id, People[0].Name, People[0].Address)
	Error.Println("If an error occurs use this method.")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", rootUrl)
	myRouter.HandleFunc("/all", listPeople)
	myRouter.HandleFunc("/person", createNewPerson).Methods("POST")
	myRouter.HandleFunc("/person/{id}", returnSinglePerson)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
