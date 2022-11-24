package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type series struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Story string `json:"story"`
}

//var seriesData = []series{{
//	Id:    1,
//	Name:  "TMKOC",
//	Story: "hehe lol"},
//}

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

func (handler *handler)updateSeries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		fmt.Println("atleast one")
		return
	}

	params := mux.Vars(r)
	var series series
	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		fmt.Println("issue with what ")
	}
	for index, tt := range  {
		if tt.Name == params["name"] {
			seriesData = append(seriesData[:index], seriesData[index+1:]...)
			seriesData = append(seriesData, series)
			json.NewEncoder(w).Encode(seriesData)
			return
		}
	}

}
func (handler *handler)createSeries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hy welcome to create one series")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		fmt.Println("atleast one")
	}
	var series series
	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		fmt.Println("issue with what ")
	}
	//seriesData = append(seriesData, series)
	seriesDta,err:=handler.db.ExecContext(context.Background(),"INSERT INTO series values(?,?,?)",series.Id,series.Name,series.Story)
	json.NewEncoder(w).Encode(seriesDta)
}
func main() {
	r := mux.NewRouter()
	db, _ := sql.Open("mysql", "root:my-secret-pw@/series")
	h:=New(db)
	r.HandleFunc("/serve", h.createSeries).Methods("POST")
	r.HandleFunc("/serve/{name}", h.updateSeries).Methods("PUT")
	http.ListenAndServe(":4500", r)
}
