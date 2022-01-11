package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/internal/model"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}
func (s *Server) GetAllArticle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetAllArticle(s.DB, w, r)
}

func (s *Server) CreateArticle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.CreateArticle(s.DB, w, r)
}

func (s *Server) GetArticle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetArticle(s.DB, w, r)
}

func (s *Server) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.UpdateArticle(s.DB, w, r)
}

func (s *Server) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.DeleteArticle(s.DB, w, r)
}
func (s *Server) GetArticleByTitle(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetArticleByTitle(s.DB, w, r)
}
func (s *Server) GetAllArticleByCatagory(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetAllArticleByCatagory(s.DB, w, r)
}

func (s *Server) GetArticleByTag(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetArticleByTag(s.DB, w, r)
}
func (s *Server) GetArticleByCategory(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetArticleByCategory(s.DB, w, r)
}
func (s *Server) GetTrending(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetTrending(s.DB, w, r)
}
func (s *Server) GetArticleByPage(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	article.GetArticleByPage(s.DB, w, r)
}
