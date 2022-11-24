package stores

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agrawalananya/goMovies/internal/models"
	"testing"
)

func TestGetMovies(t *testing.T) {
	// open database stub
	db, mock, err := sqlmock.New()
	if err == nil {
		fmt.Println("fehvfhbv")
	}
	movies := models.Movies{
		Id:       1,
		Name:     "sdds",
		Genre:    "dscsd",
		Rating:   0,
		Plot:     "sdvsdafv",
		Released: false,
	}
	testCase := []struct {
		mock   interface{}
		output models.Movies
		desc   string
	}{
		{
			mock:   mock.ExpectQuery("Select * from movies where id =?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "plot", "released"})),
			output: movies,
			desc:   "SUCCESS",
		},
	}
	h := New(db)
	for _, tt := range testCase {

	}

}
