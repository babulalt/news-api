package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/internal/model"
)

func (s *Server) GetAllComment(w http.ResponseWriter, r *http.Request) {
	comment := &model.Comment{}
	comment.GetAllComment(s.DB, w, r)
}
func (s *Server) CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := &model.Comment{}
	comment.CreateComment(s.DB, w, r)
}

func (s *Server) GetComment(w http.ResponseWriter, r *http.Request) {
	comment := &model.Comment{}
	comment.GetComment(s.DB, w, r)
}

func (s *Server) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment := &model.Comment{}
	comment.UpdateComment(s.DB, w, r)
}

func (s *Server) DeleteComment(w http.ResponseWriter, r *http.Request) {
	comment := &model.Comment{}
	comment.DeleteComment(s.DB, w, r)
}
