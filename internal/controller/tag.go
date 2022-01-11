package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/internal/model"
)

func (s *Server) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags := &model.Tag{}
	tags.GetAllTags(s.DB, w, r)
}

func (s *Server) CreateTag(w http.ResponseWriter, r *http.Request) {
	tag := &model.Tag{}
	tag.CreateTag(s.DB, w, r)
}

func (s *Server) GetTag(w http.ResponseWriter, r *http.Request) {
	tag := &model.Tag{}
	tag.GetTag(s.DB, w, r)
}

func (s *Server) UpdateTag(w http.ResponseWriter, r *http.Request) {
	tag := &model.Tag{}
	tag.UpdateTag(s.DB, w, r)
}

func (s *Server) DeleteTag(w http.ResponseWriter, r *http.Request) {
	tag := &model.Tag{}
	tag.DeleteTag(s.DB, w, r)
}
