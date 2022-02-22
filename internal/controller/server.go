package controller

import (
	"fmt"
	"log"

	"github.com/berrybytes/sugam/config"
	"github.com/berrybytes/sugam/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) Initialize(config *config.Config) {
	db, err := gorm.Open(sqlite.Open(config.DB.Name), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database %v", err)
	}
	s.DB = db
	fmt.Println("Database Connected ")
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Author{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.Comment{})
}
