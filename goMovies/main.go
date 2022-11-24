package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/agrawalananya/goMovies/internal/stores"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//	func getAllMovies(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("get all the movies")
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(movies)
//	}
//
//	func (handler *handler) getOneMovie(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("get one movie")
//		w.Header().Set("Content-Type", "application/json")
//
//		//id from request
//		params := mux.Vars(r)
//		fmt.Println(params)
//		ctx := context.Background()
//		//val, err := db.QueryContext(ctx, "select * from movies where id = ?", params["id"])
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
//
//		defer val.Close()
//		var movie Movies
//		for val.Next() {
//			err := val.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.Rating, &movie.Plot, &movie.Released)
//			if err != nil {
//				log.Fatal(err)
//			}
//			if val != nil {
//				json.NewEncoder(w).Encode(
//					MoviesData{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &movie}})
//				return
//			}
//		}
//		err = val.Err()
//		if err != nil {
//			log.Fatal(err)
//		}
//		// looping through movies, find matching id and return the response
//		//for _, movie := range movies {
//		//	movieId, err := strconv.Atoi(params["id"])
//		//	if err != nil {
//		//		errors.New("strconv failure")
//		//	}
//		//	if movie.Id == movieId {
//		//		json.NewEncoder(w).Encode(
//		//			MoviesData{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &movie}})
//		//		return
//		//	}
//		//}
//		// if not found anything
//		json.NewEncoder(w).Encode(MoviesData{Error: "No movie found with the id"})
//		//json.NewEncoder(w).Encode("No movie found with the id")
//		return
//	}
//
// //	func (handler *handler) createOneMovieData(w http.ResponseWriter, r *http.Request) {
// //		fmt.Println("create one movie")
// //		w.Header().Set("Content-Type", "application/json")
// //
// //		//body is empty
// //		if r.Body == nil {
// //			json.NewEncoder(w).Encode("Plesae enter atlease one movie data")
// //		}
// //
// //		// if response is {}
// //		var movie Movies
// //		_ = json.NewDecoder(r.Body).Decode(&movie)
// //		if isEmpty(movie) {
// //			json.NewEncoder(w).Encode("no data inside")
// //			return
// //		}
// //		ctx := context.Background()
// //
// //		_, err := handler.db.ExecContext(ctx, "INSERT INTO movies(name,genre,rating,plot,released) values(?,?,?,?,?)", movie.Name, movie.Genre, movie.Rating, movie.Plot, movie.Released)
// //		if err != nil {
// //
// //		}
// //		// generate random id
// //		//rand.Seed(time.Now().UnixNano())
// //
// //		//movie.Id = 1
// //		//movies = append(movies, movie)
// //
// //		var responseOutput MoviesData
// //		if movie.Id == 2 {
// //
// //			responseOutput = MoviesData{
// //				Code:   http.StatusForbidden,
// //				Status: "FAILURE",
// //				Data:   nil,
// //			}
// //			w.WriteHeader(http.StatusForbidden)
// //			json.NewEncoder(w).Encode(&responseOutput)
// //			return
// //		}
// //
// //		responseOutput = MoviesData{
// //			Code:   200,
// //			Status: "SUCCESS",
// //			Data:   &Data{Movie: &movie},
// //		}
// //		w.WriteHeader(http.StatusOK)
// //
// //		json.NewEncoder(w).Encode(&responseOutput)
// //		return
// //	}
// var hand *stores.Store
//
//	func updateMovie(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("update movie")
//		w.Header().Set("Content-Type", "application/json")
//
//		// id from request
//		params := mux.Vars(r)
//		for index, movie := range movies {
//			movieId, err := strconv.Atoi(params["id"])
//			if err != nil {
//				panic(err)
//			}
//			if movie.Id == movieId {
//				movies = append(movies[:index], movies[index+1:]...)
//				var movie Movies
//				_ = json.NewDecoder(r.Body).Decode(&movie)
//				movie.Id = movieId
//				movies = append(movies, movie)
//				json.NewEncoder(w).Encode(MoviesData{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &movie}})
//				return
//			}
//		}
//		hand.CreateOneMovieData()
//		db.ExecContext(context.Background(), "UPDATE movies SET movieName = ? WHERE genre = ?", "DDLJ", "valleyy")
//		// if not found anything
//		json.NewEncoder(w).Encode(MoviesData{Error: "No movie found with the id"})
//		//json.NewEncoder(w).Encode("No movie found with the id")
//		return
//	}
//
//	func deleteOneMovie(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("delete movie")
//		w.Header().Set("Content-Type", "application/json")
//		params := mux.Vars(r)
//		for index, movie := range movies {
//			movieId, err := strconv.Atoi(params["id"])
//			if err != nil {
//				panic(err)
//			}
//			if movie.Id == movieId {
//				movies = append(movies[:index], movies[index+1:]...)
//				break
//			}
//		}
//	}
func getDB() (*sql.DB, error) {
	//connections string user:pass@proto(host:port)/database
	conn := "root:@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to db, err:", err)
		return nil, errors.New(fmt.Sprintln("Failed to connect to db, err:", err))
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to db, err:", err)
		return nil, errors.New(fmt.Sprintln("Failed to connect to db, err:", err))
	}

	return db, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}
	h := stores.New(db)
	defer db.Close()
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":4000", r))
}
