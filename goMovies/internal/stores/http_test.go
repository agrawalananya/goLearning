package stores

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agrawalananya/goMovies/internal/models"
	"reflect"
	"testing"
)

func TestPostMovies(t *testing.T) {
	// open database stub
	movies := models.Movies{
		Id:       0,
		Name:     "sdds",
		Genre:    "dscsd",
		Rating:   0,
		Plot:     "sdvsdafv",
		Released: false,
	}
	db, mock, err := sqlmock.New()
	if err == nil {
		fmt.Println("fehvfhbv")
	}
	testCase := []struct {
		input         models.Movies
		mock          interface{}
		output        models.Movies
		desc          string
		ExpectedError string
	}{
		{
			input: models.Movies{
				Id:       0,
				Name:     "sdds",
				Genre:    "dscsd",
				Rating:   0,
				Plot:     "sdvsdafv",
				Released: false,
			},
			mock: mock.ExpectExec("INSERT INTO movies(movieName,genre,rating,plot,released) values(?,?,?,?,?)").WithArgs(movies.Name, movies.Genre, movies.Rating, movies.Plot, movies.Released).WillReturnResult(sqlmock.NewResult(1, 1)),
			desc: "SUCCESS",

			output: models.Movies{
				Id:       0,
				Name:     "sdds",
				Genre:    "dscsd",
				Rating:   0,
				Plot:     "sdvsdafv",
				Released: false,
			},
		},
		{
			input: models.Movies{
				Id:       0,
				Name:     "some name ",
				Genre:    "dscsd",
				Rating:   0,
				Plot:     "sdvsdafv",
				Released: false,
			},
			desc:          "FAIL",
			ExpectedError: "an error",
		},
	}

	h := New(db)
	for _, tt := range testCase {
		ans, errorFrom := h.CreateOneMovieData(&tt.input)
		if errorFrom != nil {
			t.Fatalf("ksdkkhcadsvkhfcdavhcd")
		}
		if !reflect.DeepEqual(ans, tt.output) {
			fmt.Println("some issue is there ")
		}

	}

}
