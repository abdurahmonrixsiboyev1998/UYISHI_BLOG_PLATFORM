package handler

import (
	"blog/internal/repository"
	"blog/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterHandlers(r *mux.Router, db *gorm.DB) {
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var UserData struct {
			Username string `json:"username"`
			Email string `json:"email"`
			Password string `json:"password"`
		}
		json.NewDecoder(r.Body).Decode(&UserData)

		if err := userService.Register(UserData.Username, UserData.Email, UserData.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var loginData struct{
			Email string `json:"email"`
			Password string `json:"password"`
		}
		json.NewDecoder(r.Body).Decode(&loginData)

		token, err := userService.Login(loginData.Email, loginData.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}).Methods("POST")
}
