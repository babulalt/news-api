package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique" `
	Description string `json:"description"`
}

func (c *Category) GetAllCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	categories := []Category{}
	db.Find(&categories)
	if err := db.Find(&categories).Error; err != nil {
		return
	}
	respondJSON(w, http.StatusOK, categories)
}

func (c *Category) CreateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	category := Category{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, category)
}

func (c *Category) GetCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	category := c.getCategoryOr404(db, id, w, r)
	if category == nil {
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func (c *Category) UpdateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	category := c.getCategoryOr404(db, id, w, r)
	if category == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func (c *Category) DeleteCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	catagory := c.getCategoryOr404(db, id, w, r)
	if catagory == nil {
		return
	}
	if err := db.Delete(&catagory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func (c *Category) getCategoryOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Category {
	category := Category{}
	if err := db.First(&category, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &category
}
