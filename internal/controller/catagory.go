package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/internal/model"
)

func (s *Server) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	category := &model.Category{}
	category.GetAllCategory(s.DB, w, r)
}

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := &model.Category{}
	category.CreateCategory(s.DB, w, r)
}

func (s *Server) GetCategory(w http.ResponseWriter, r *http.Request) {
	category := &model.Category{}
	category.GetCategory(s.DB, w, r)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := &model.Category{}
	category.UpdateCategory(s.DB, w, r)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	category := &model.Category{}
	category.DeleteCategory(s.DB, w, r)
}
