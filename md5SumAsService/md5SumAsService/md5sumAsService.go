package main

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"net/http"
)

type dataToPass struct {
	Submissions map[string][]string
}

func handler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	sum := md5.Sum(body)
	w.Header().Set("content-type", "application/json")
	w.Write(sum[:])
}

func main() {

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
