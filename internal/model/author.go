package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string `json:"firstname" gorm:"not null"`
	LastName  string `json:"lastname" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	Username  string `gorm:"unique" json:"username" `
	Password  string `json:"-" gorm:"not null"`
	ImageURL  string `json:"image"`
	Is_active bool   `json:"is_active"`
}

func (a *Author) GetAllAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	author := []Author{}
	db.Find(&author)
	respondJSON(w, http.StatusOK, author)
}

func (a *Author) CreateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	author := Author{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	author.Password, _ = hashPassword(author.Password)
	if err := db.Save(&author).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, author)
}

func (a *Author) GetAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	author := a.getAuthorOr404(db, id, w, r)
	if author == nil {
		return
	}
	respondJSON(w, http.StatusOK, author)
}

func (a *Author) UpdateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	author := a.getAuthorOr404(db, id, w, r)
	if author == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	author.Password, _ = hashPassword(author.Password)

	if err := db.Save(&author).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, author)
}

func (a *Author) DeleteAuthor(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	author := a.getAuthorOr404(db, id, response, request)
	if author == nil {
		return
	}
	if err := db.Delete(&author).Error; err != nil {
		respondError(response, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(response, http.StatusNoContent, nil)
}

func (a *Author) getAuthorOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Author {
	author := Author{}
	if err := db.First(&author, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &author
}
