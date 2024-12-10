package repository

import (
	"go-personal-website-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	GetProjectData(ctx echo.Context) (records []models.MstProjects, err error)
	InsertPorjectData(ctx echo.Context, request models.MstProjects) (rowAffected int, err error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (repo *repository) GetProjectData(ctx echo.Context) (records []models.MstProjects, err error) {
	result := repo.DB.Find(&records)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (repo *repository) InsertPorjectData(ctx echo.Context, request models.MstProjects) (rowAffected int, err error) {
	result := repo.DB.Create(&request)
	if result.Error != nil {
		err = result.Error
		return
	}

	rowAffected = int(result.RowsAffected)
	return
}
