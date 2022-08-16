package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func (person *Person) getHomeworld() {
	resp, err := http.Get(person.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld:", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Print("Error reading response body:", err)
	}
	json.Unmarshal(bytes, &person.Homeworld)
}

// create a route for /people
// write a function that uses fmt.Fprint that indicates the request was successful
func getPeople(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Getting people")
	resp, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request Star Wars people: ", err)
	}
	fmt.Println(resp)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body: ", err)
	}

	fmt.Println(string(bytes))

	var people AllPeople
	if err := json.Unmarshal(bytes, &people); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse JSON: ", err)
	}
	fmt.Println(people)

	for _, person := range people.People {
		person.getHomeworld()
		fmt.Println(person)
	}
}

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
