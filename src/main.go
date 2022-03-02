package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"module/name"
	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	for _, singleEven := range events {
		if singleEven.ID == eventId {
			json.NewEncoder(w).Encode(singleEven)
			return
		}
	}

}

func getAllEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event/{id}", getEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvent).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	fmt.Print("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
