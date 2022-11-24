package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/agrawalananya/goMovies/internal/models"
	"log"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db}
}
func (h Store) CreateOneMovieData(movie *models.Movies) (models.Movies, error) {
	ctx := context.Background()
	_, err := h.db.ExecContext(ctx, "INSERT INTO movies(movieName,genre,rating,plot,released) values(?,?,?,?,?)", movie.Name, movie.Genre, movie.Rating, movie.Plot, movie.Released)
	if err != nil {
		return models.Movies{}, fmt.Errorf("some error is there")
	}
	return *movie, nil
}
func (h Store) getOneMovie(id int) models.Movies {
	ctx := context.Background()
	val, err := h.db.QueryContext(ctx, "select * from movies where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	var movie models.Movies
	for val.Next() {
		val.Scan(&movie.Id, &movie.Name, &movie.Plot, &movie.Genre, &movie.Released, &movie.Released)
	}
	return movie
}

func (h Store) updateMovie(id int, movies models.Movies) (models.Movies, error) {
	_, err := h.db.ExecContext(context.Background(), "UPDATE movies SET movieName = ?,genre =?,plot=?, rating=?,released=?  WHERE id = ?;", movies.Name, movies.Genre, movies.Plot, movies.Rating, movies.Released, id)
	if err != nil {
		return models.Movies{}, fmt.Errorf("there is an error")
	}
	val, err := h.db.QueryContext(context.Background(), "SELECT * from movies where id=?;", id)
	var movie models.Movies
	for val.Next() {
		val.Scan(&movie.Id, &movie.Name, &movie.Plot, &movie.Genre, &movie.Released, &movie.Released)
	}
	return movie, err
}
