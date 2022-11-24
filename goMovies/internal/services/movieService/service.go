package movieService

import (
	"github.com/agrawalananya/goMovies/internal/services"
	"github.com/agrawalananya/goMovies/internal/stores"
)

type Service struct {
	store stores.StoreHandler
}

func (s Service) gorun() {
	//TODO implement me
	panic("implement me")
}

func New(store stores.StoreHandler) services.ServiceHandler {
	return &Service{store: store}
}
