package http

import (
	"github.com/Deanamic/kitchen-helper/pkg/domain"
	"github.com/Deanamic/kitchen-helper/pkg/recipes"
	"github.com/Deanamic/kitchen-helper/pkg/users"
	"net/http"
)

type Service struct {
	Router *http.ServeMux
}

func New(r domain.Repository) Service {
	var service Service

	//delegate to user and recipe service and merge here!!!
	recipesService := recipes.NewService(r)
	usersService := users.NewService(r)

	router := http.NewServeMux()
	router.Handle("/users/", http.StripPrefix("/users", usersService.Router))
	router.Handle("/recipes/", http.StripPrefix("/recipes", recipesService.Router))
	service.Router = router

	return service
}
