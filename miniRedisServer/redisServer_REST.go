package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Record struct {
	ID   string `json:"id,omitempty`
	Name string `json:"firstname,omitempty`
}

var people []Record

func GetRecordEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item.Name)
			return
		}
	}
	json.NewEncoder(w).Encode(&Record{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	var num = 0
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode(len(people))
		return
	}
	for _, item := range people {
		if strings.HasPrefix(item.ID, params["id"]) {
			num++
		}
	}
	json.NewEncoder(w).Encode(num)
	return

}
func CreateRecordEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var record Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	record.ID = params["id"]
	people = append(people, record)
	json.NewEncoder(w).Encode(people)
}
func DeleteRecordEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Record{ID: "total_records", Name: "3"})
	people = append(people, Record{ID: "total_bytes", Name: "2"})
	people = append(people, Record{ID: "min_records", Name: "1"})
	router.HandleFunc("/peopleCount", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/peopleCount/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetRecordEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreateRecordEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeleteRecordEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
