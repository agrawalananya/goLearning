package stores

import "github.com/agrawalananya/goMovies/internal/models"

type StoreHandler interface {
	CreateOneMovieData(movie *models.Movies) (models.Movies, error)
	getOneMovie(id int) models.Movies
	updateMovie(id int, movies models.Movies) (models.Movies, error)
}
