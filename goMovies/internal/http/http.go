package http

import (
	"encoding/json"
	"fmt"
	"github.com/agrawalananya/goMovies/internal/models"
	"github.com/agrawalananya/goMovies/internal/stores"
	"net/http"
)

type Handler struct {
	store *stores.StoreHandler
}

func New(store *stores.StoreHandler) *Handler {
	return &Handler{store: store}
}

func (h Handler) CreateOneMovieData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one movie")
	w.Header().Set("Content-Type", "application/json")

	//body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plesae enter atleast one movie data")
	}

	// if response is {}
	var movie models.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if main.isEmpty(movie) {
		json.NewEncoder(w).Encode("no data inside")
		return
	}

	// generate random id
	//rand.Seed(time.Now().UnixNano())

	//movie.Id = 1
	//movies = append(movies, movie)

	var responseOutput models.MoviesData
	if movie.Id == 2 {

		responseOutput = models.MoviesData{
			Code:   http.StatusForbidden,
			Status: "FAILURE",
			Data:   nil,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}

	responseOutput = models.MoviesData{
		Code:   200,
		Status: "SUCCESS",
		Data:   &models.Data{Movie: &movie},
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&responseOutput)
	return
}
