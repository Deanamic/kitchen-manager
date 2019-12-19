package recipes

import (
	"github.com/Deanamic/kitchen-helper/pkg/domain"
	"github.com/go-chi/chi"
	"net/http"
)

type Service struct {
	repo   domain.Repository
	Router *chi.Mux
}

func NewService(r domain.Repository) Service {
	service := Service{
		repo: r,
	}
	router := chi.NewRouter()
	router.Get("/recipes", foo)
	router.Get("/recipe/{recipename}", foo)
	service.Router = router
	return service
}

func foo (w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
