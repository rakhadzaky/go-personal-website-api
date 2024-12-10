package http

import (
	"errors"
	"go-personal-website-api/domains"
	"go-personal-website-api/services/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *httpHandler) GetProjectData(ctx echo.Context) error {
	fetchData, err := handler.ServiceUc.GetProjectData(ctx)
	if err != nil {
		return common.ResponseError(ctx, err)
	}

	if fetchData == nil {
		return ctx.JSON(http.StatusOK, domains.ResponseFind{
			Data:  []domains.ProjectFetchRes{},
			Page:  0,
			Total: 0,
		})
	}

	return ctx.JSON(http.StatusOK, domains.ResponseFind{
		Data:  fetchData,
		Page:  1,
		Total: 1,
	})
}

func (handler *httpHandler) PostProjectData(ctx echo.Context) error {
	var (
		requestInsert domains.ProjectInsertReq
	)

	err := ctx.Bind(&requestInsert)
	if err != nil {
		return common.ResponseError(ctx, errors.New("wrong input"))
	}

	rowAffected, err := handler.ServiceUc.InsertProjectData(ctx, requestInsert)
	if err != nil {
		return common.ResponseError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, domains.ResponseInsert{
		RowAffected: rowAffected,
	})
}
