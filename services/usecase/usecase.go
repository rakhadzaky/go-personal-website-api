package usecase

import (
	"errors"
	"go-personal-website-api/domains"
	"go-personal-website-api/models"
	"go-personal-website-api/services/repository"
	"time"

	"github.com/labstack/echo/v4"
)

type usecase struct {
	Repo repository.Repository
}

type Usecase interface {
	GetProjectData(ctx echo.Context) (fetchProjects *[]domains.ProjectFetchRes, err error)
	InsertProjectData(ctx echo.Context, requestInsert domains.ProjectInsertReq) (rowAffected int, err error)
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (uc *usecase) GetProjectData(ctx echo.Context) (fetchProjects *[]domains.ProjectFetchRes, err error) {
	var (
		projectList []domains.ProjectFetchRes
	)
	dbFetchRes, err := uc.Repo.GetProjectData(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range dbFetchRes {
		projectList = append(projectList, domains.ProjectFetchRes{
			Id:             val.Id,
			Title:          val.Title,
			Client:         val.Client,
			DateCreation:   val.Date.Format("2006-01-02 15:04:05"),
			DateCompletion: val.DateCompletion.Format("2006-01-02 15:04:05"),
			Description:    val.Description,
			Position:       val.Position,
		})
	}

	fetchProjects = &projectList

	return
}

func (uc *usecase) InsertProjectData(ctx echo.Context, requestInsert domains.ProjectInsertReq) (rowAffected int, err error) {
	var (
		dbInsertReq models.MstProjects
	)

	dateCreation, err := time.Parse("2006-01-02 15:04:05", requestInsert.DateCreation)
	if err != nil {
		return 0, errors.New("invalid date: creation")
	}

	dateCompletion, err := time.Parse("2006-01-02 15:04:05", requestInsert.DateCompletion)
	if err != nil {
		return 0, errors.New("invalid date: completion")
	}

	dbInsertReq = models.MstProjects{
		Title:          requestInsert.Title,
		Client:         requestInsert.Client,
		Date:           dateCreation,
		DateCompletion: dateCompletion,
		Description:    requestInsert.Description,
		Position:       requestInsert.Position,
		CreatedAt:      time.Now(),
		CreatedBy:      "SYSTEM",
	}

	rowAffected, err = uc.Repo.InsertPorjectData(ctx, dbInsertReq)
	if err != nil {
		return 0, err
	}

	return
}
