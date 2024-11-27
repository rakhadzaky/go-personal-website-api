package http

import (
	"log"

	"github.com/labstack/echo/v4"
)

func (handler *httpHandler) GetProjectData(ctx echo.Context) error {
	log.Println("Get Project Data from Handler")
	handler.ServiceUc.GetProjectData(ctx)
	return nil
}
