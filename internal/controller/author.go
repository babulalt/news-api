package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/internal/model"
)

func (s *Server) GetAllAuthor(w http.ResponseWriter, r *http.Request) {
	author := &model.Author{}
	author.GetAllAuthor(s.DB, w, r)
}

func (s *Server) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := &model.Author{}
	author.CreateAuthor(s.DB, w, r)
}

func (s *Server) GetAuthor(w http.ResponseWriter, r *http.Request) {
	author := model.Author{}
	author.GetAuthor(s.DB, w, r)
}

func (s *Server) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	author := &model.Author{}
	author.UpdateAuthor(s.DB, w, r)
}

func (s *Server) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	author := model.Author{}
	author.DeleteAuthor(s.DB, w, r)
}
