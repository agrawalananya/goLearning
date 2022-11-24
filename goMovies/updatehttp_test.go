package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUpdateMovie(t *testing.T) {
	testCases := []struct {
		target    string
		body      Movies
		expOutput MoviesData
		expError  error
	}{
		{
			target: "/movie",
			expOutput: MoviesData{
				Error: "No movie found with the id"},
		},
		{
			target: "/movie",
			body: Movies{
				1, "hiiii Valley", "romantic", 7.5, "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.", true},
			expOutput: MoviesData{Code: 200, Status: "SUCCESS",
				Data: &Data{Movie: &Movies{1, "hiiii Valley", "romantic", 7.5, "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.", true}},
			},
		},
	}

	for i, tt := range testCases {
		jsnBdy, _ := json.Marshal(tt.body)
		buff := bytes.NewBuffer(jsnBdy)
		r := httptest.NewRequest(http.MethodGet, tt.target, buff)
		w := httptest.NewRecorder()
		params := map[string]string{
			"id": "1",
		}
		if i == 0 {
			params["id"] = "1234"
		}
		r = mux.SetURLVars(r, params)
		updateMovie(w, r)

		var movieResponse MoviesData
		res := w.Result()
		data, err := io.ReadAll(res.Body)

		if err != tt.expError {
			t.Fatalf("Wrong Output.")
			return
		}

		err = json.Unmarshal(data, &movieResponse)
		if err != nil {
			t.Fatalf("Wrong Output nil.")
			return
		}

		if !reflect.DeepEqual(movieResponse, tt.expOutput) {
			t.Fatalf("Wrong Output, Expected: %v, Got: %v", tt.expOutput, movieResponse)
		}
		fmt.Printf("output, Expected: %v, Got: %v\n", tt.expOutput, movieResponse)
	}
}
