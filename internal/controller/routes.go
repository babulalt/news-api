package controller

import (
	"net/http"

	"github.com/berrybytes/sugam/config"
	"github.com/gorilla/mux"
)

func (s *Server) SetRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/", s.Home).Methods("GET")
	//endpoints for author
	router.HandleFunc("/authors", s.GetAllAuthor).Methods("GET")
	router.HandleFunc("/author", s.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/{id}", s.GetAuthor).Methods("GET")
	router.HandleFunc("/author/{id}", s.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{id}", s.DeleteAuthor).Methods("DELETE")

	//endpoints for articles
	router.HandleFunc("/articles", s.GetAllArticle).Methods("GET")
	router.HandleFunc("/article", s.CreateArticle).Methods("POST")
	router.HandleFunc("/article/{id}", s.GetArticle).Methods("GET")
	router.HandleFunc("/article/{title}", s.GetArticleByTitle).Methods("GET")
	router.HandleFunc("/articles/{categories}", s.GetAllArticleByCatagory).Methods("GET")
	router.HandleFunc("/articles/tag/{id}", s.GetArticleByTag).Methods("GET")
	router.HandleFunc("/articles/category/{id}", s.GetArticleByCategory).Methods("GET")
	router.HandleFunc("/trending", s.GetTrending).Methods("GET")
	router.HandleFunc("/article/{page}/page", s.GetArticleByPage).Methods("GET")
	router.HandleFunc("/article/{id}", s.UpdateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", s.DeleteArticle).Methods("DELETE")

	//endpoints for catagory
	router.HandleFunc("/categories", s.GetAllCategory).Methods("GET")
	router.HandleFunc("/category", s.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", s.GetCategory).Methods("GET")
	router.HandleFunc("article/{category}", s.GetAllArticleByCatagory).Methods("GET")
	router.HandleFunc("/category/{id}", s.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", s.DeleteCategory).Methods("DELETE")

	//endpoints for comment
	router.HandleFunc("/comments", s.GetAllComment).Methods("GET")
	router.HandleFunc("/article/{id}/comment", s.CreateComment).Methods("POST")
	router.HandleFunc("/comment/{id}", s.GetComment).Methods("GET")
	router.HandleFunc("/comment/{id}", s.UpdateComment).Methods("PUT")
	router.HandleFunc("/comment/{id}", s.DeleteComment).Methods("DELETE")

	//endpoints for tags
	router.HandleFunc("/tags", s.GetAllTags).Methods("GET")
	router.HandleFunc("/tag", s.CreateTag).Methods("POST")
	router.HandleFunc("/tag/{id}", s.GetTag).Methods("GET")
	router.HandleFunc("/tag/{id}", s.UpdateTag).Methods("PUT")
	router.HandleFunc("/tag/{id}", s.DeleteTag).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}

func Run() {
	config := config.GetConfig()
	server := &Server{}
	server.Initialize(config)
	server.SetRouters()

}
