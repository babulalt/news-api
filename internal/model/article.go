package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	CategoryID    int        `json:"category_id" gorm:"not null"`
	Category      *Category  `json:"category" gorm:"foreignkey:CategoryID"`
	Title         string     `json:"title" gorm:"unique"`
	Image         string     `json:"image"`
	Image_Caption string     `json:"image_caption"`
	Content       string     `json:"content"`
	TagID         int        `json:"tag_id" gorm:"not null"`
	Tag           *Tag       `json:"tag" gorm:"foreignkey:TagID"`
	Comments      int        `json:"comments" gorm:"default 0"`
	Comment       []*Comment `json:"-" gorm:"many2many:article_comments"`
	AuthorID      uint       `json:"author_id" gorm:"not null"`
	Author        *Author    `gorm:"foreignkey:AuthorID" json:"author"`
	Views         int        `gorm:"default 0"`
}

func (a *Article) GetAllArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	articles := []Article{}
	result := db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article)
	length := int(result.RowsAffected)
	for i := 0; i >= length; i++ {
		article[i].Comments = (len(article[i].Comment))
	}
	for i := length - 1; i >= 0; i-- {
		articles = append(articles, article[i])
	}
	respondJSON(w, http.StatusOK, articles[:5])
}
func (a *Article) GetArticleByPage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	articles := []Article{}
	vars := mux.Vars(r)
	pageno := vars["page"]
	page, _ := strconv.Atoi(pageno)
	result := db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article)
	length := int(result.RowsAffected)
	for i := 0; i >= length; i++ {
		article[i].Comments = (len(article[i].Comment))
	}
	for i := length - 1; i >= 0; i-- {
		articles = append(articles, article[i])
	}
	if page == 1 {
		respondJSON(w, http.StatusOK, articles[:4])
	} else if page == 2 {
		respondJSON(w, http.StatusOK, articles[4:length])
	} else {
		respondError(w, http.StatusNotFound, "page not available")
	}

}

func (a *Article) GetAllArticleByCatagory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	articles := []Article{}
	vars := mux.Vars(r)
	cname := vars["categories"]
	result := db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article)
	length := int(result.RowsAffected)
	for i := 0; i < length; i++ {
		if article[i].Category.Name == cname {
			articles = append(articles, article[i])
		}
	}
	respondJSON(w, http.StatusOK, articles)
}

func (a *Article) CreateArticle(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	article := Article{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&article); err != nil {
		respondError(response, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	err := db.Save(&article).Error
	if err != nil {
		respondError(response, http.StatusInternalServerError, err.Error())
		return
	}
	db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article, article.ID)
	article.Comments = len(article.Comment)
	respondJSON(response, http.StatusCreated, article)

}

func (a *Article) GetArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	article := a.getArticleOr404(db, id, w, r)
	if article == nil {
		return
	}
	article.Views++
	if err := db.Save(&article).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, article)
}

func (a *Article) GetArticleByTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := Article{}
	vars := mux.Vars(r)
	cname := vars["title"]
	fmt.Println(cname)
	db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article, Article{Title: cname})
	fmt.Println(article.Category.Name)
	respondJSON(w, http.StatusOK, article)
}
func (a *Article) GetArticleByTag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	vars := mux.Vars(r)
	tid := vars["id"]
	id, _ := strconv.Atoi(tid)
	db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article, Article{TagID: id})
	respondJSON(w, http.StatusOK, article)
}
func (a *Article) GetArticleByCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	vars := mux.Vars(r)
	tid := vars["id"]
	id, _ := strconv.Atoi(tid)
	db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article, Article{CategoryID: id})
	respondJSON(w, http.StatusOK, article)
}
func (a *Article) GetTrending(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article := []Article{}
	s := []Article{}
	db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").Find(&article)
	sort.Slice(article, func(i, j int) bool { return article[i].Views > article[j].Views })
	for _, articles := range article {
		s = append(s, articles)
		fmt.Print()
	}

	respondJSON(w, http.StatusOK, s[:5])
}
func (a *Article) UpdateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	article := a.getArticleOr404(db, id, w, r)
	if article == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&article).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	article.Comments = len(article.Comment)
	respondJSON(w, http.StatusOK, article)
}

func (a *Article) DeleteArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.Atoi(aid)
	article := a.getArticleOr404(db, id, w, r)
	if article == nil {
		return
	}
	if err := db.Delete(&article).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func (a *Article) getArticleOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Article {
	article := Article{}
	if err := db.Model(&Article{}).Preload("Author").Preload("Tag").Preload("Category").Preload("Comment").First(&article, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	article.Comments = len(article.Comment)
	return &article
}
