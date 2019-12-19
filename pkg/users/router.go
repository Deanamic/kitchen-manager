package users

import (
	"github.com/Deanamic/kitchen-helper/pkg/domain"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/go-chi/chi"
	"fmt"
	// "context"
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
	router.Post("/login", foo)
	router.Put("/register", service.registerHandler)
	router.Get("/user/{username}", service.userHandler)
	service.Router = router
	return service
}

func foo (w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

// userHandler ...
func (s Service) userHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("HERE")
	// user, err := s.repo.FindUser(r.Context().Value("username").(string))
	username := chi.URLParam(r, "username")
	user, err := s.repo.FindUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

// registerHandler ...
func (s Service) registerHandler(w http.ResponseWriter, r *http.Request)  {
	payload, _ := ioutil.ReadAll(r.Body)

	fmt.Println(payload)
	newUser := domain.User{}
	_ = json.Unmarshal(payload, &newUser)
	fmt.Println(newUser)

	err := s.repo.CreateUser(newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// func userCtx(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		userId := chi.URLParam(r, "username")
//		ctx := context.WithValue(r.Context(), "username", userId)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
// }
