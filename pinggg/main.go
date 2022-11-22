package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var count int = 0

type ping struct {
	Data string `json:"data"`
}

var pinggg = ping{}

func serveMovie(w http.ResponseWriter, r *http.Request) {
	count++
	w.Header().Set("x-count", strconv.Itoa(count))
	// to show struct data in  the body
	json.NewEncoder(w).Encode(pinggg)
	fmt.Println(count)
}
func postMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func otherMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	pinggg.Data = "pingggg"
	r.HandleFunc("/ping", serveMovie).Methods("GET")
	r.HandleFunc("/", postMethod).Methods("POST")
	r.HandleFunc("", otherMethod).Methods("")
	log.Fatal(http.ListenAndServe(":4000", r))
}
