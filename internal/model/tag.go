package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string `json:"name"`
	Sub_Name string `json:"sub_name"`
}

func (t *Tag) GetAllTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tags := []Tag{}
	db.Find(&tags)
	if err := db.Find(&tags).Error; err != nil {
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func (t *Tag) CreateTag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tag := Tag{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tag); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&tag).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, tag)
}

func (t *Tag) GetTag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	tag := t.getTagOr404(db, id, w, r)
	if tag == nil {
		return
	}
	respondJSON(w, http.StatusOK, tag)
}

func (t *Tag) UpdateTag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	tag := t.getTagOr404(db, id, w, r)
	if tag == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tag); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&tag).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tag)
}

func (t *Tag) DeleteTag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	tag := t.getTagOr404(db, id, w, r)
	if tag == nil {
		return
	}
	if err := db.Delete(&tag).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func (t *Tag) getTagOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Tag {
	tag := Tag{}
	if err := db.First(&tag, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &tag
}
