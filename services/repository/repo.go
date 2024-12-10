package repository

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	GetProjectData(ctx echo.Context) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (repo *repository) GetProjectData(ctx echo.Context) error {
	log.Println("Get Project Data from Repository")
	return nil
}

func (repo *repository) InsertPorjectData(ctx echo.Context) error {
	return nil
}
