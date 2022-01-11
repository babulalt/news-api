package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Email   string `json:"email"`
	Content string `json:"content" gorm:"not null"`
}

func (c *Comment) GetAllComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	comments := []Comment{}
	if err := db.Find(&comments).Error; err != nil {
		return
	}
	respondJSON(w, http.StatusOK, comments)
}

func (c *Comment) CreateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	comment := Comment{}
	a := &Article{}
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&comment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	article := a.getArticleOr404(db, id, w, r)
	article.Comment = append(article.Comment, &comment)
	if article == nil {
		return
	}
	if err := db.Save(&article).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	article.Comments = len(article.Comment)
	respondJSON(w, http.StatusOK, article)
}

func (c *Comment) GetComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	comment := c.getCommentOr404(db, id, w, r)
	if comment == nil {
		return
	}
	respondJSON(w, http.StatusOK, comment)
}

func (c *Comment) UpdateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	comment := c.getCommentOr404(db, id, w, r)
	if comment == nil {
		return
	}
	author := c.getCommentOr404(db, id, w, r)
	if author == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&comment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, comment)
}

func (c *Comment) DeleteComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	comment := c.getCommentOr404(db, id, w, r)
	if comment == nil {
		return
	}
	if err := db.Delete(&comment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
func (c *Comment) getCommentOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Comment {
	comment := Comment{}
	if err := db.First(&comment, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &comment
}
