package http

import (
	"go-personal-website-api/services/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type httpHandler struct {
	ServiceUc usecase.Usecase
}

type HttpHandler interface {
	GetProjectData(ctx echo.Context) error
}

func NewHttpHandler(router *echo.Echo, serviceUc usecase.Usecase) {
	h := httpHandler{
		ServiceUc: serviceUc,
	}

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.GET("/get-project", h.GetProjectData)
}
