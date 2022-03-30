package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"simaopoc/database"
)

type Service struct {
	router *mux.Router
	db     *gorm.DB
}

func New(r *mux.Router, db *gorm.DB) *Service {
	err := database.Seed(db)
	if err != nil {
		panic(err)
	}

	return &Service{
		router: r,
		db:     db,
	}
}

func (s Service) FetchData(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("endpoint unimplemented!"))
	w.WriteHeader(http.StatusOK)
	return
}

func (s Service) CreateRandomUser(w http.ResponseWriter, req *http.Request) {
	err := s.db.Create(&database.User{
		Username: "jorge.valadas",
		Password: "gostodegelado",
		Email:    "jorge.valadas@fct.pt",
	}).Error
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("user created successfully!"))
	w.WriteHeader(http.StatusOK)
	return
}
