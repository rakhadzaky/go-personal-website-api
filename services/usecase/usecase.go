package usecase

import (
	"go-personal-website-api/services/repository"
	"log"

	"github.com/labstack/echo/v4"
)

type usecase struct {
	Repo repository.Repository
}

type Usecase interface {
	GetProjectData(ctx echo.Context) error
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (uc *usecase) GetProjectData(ctx echo.Context) error {
	log.Println("Get Project Data from Usecase")
	uc.Repo.GetProjectData(ctx)
	return nil
}
