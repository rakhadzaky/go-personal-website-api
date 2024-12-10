package common

import (
	"go-personal-website-api/domains"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ResponseError(ctx echo.Context, err error) error {
	var (
		respCode    int
		respMessage string
	)
	switch {
	case strings.Contains(err.Error(), "invalid date"),
		strings.Contains(err.Error(), "wrong input"):
		respCode = http.StatusBadRequest
		respMessage = err.Error()
	default:
		log.Println("[INFO] [500]", err.Error())
		respCode = http.StatusInternalServerError
		respMessage = "something went wrong, please try again later"
	}

	return ctx.JSON(respCode, domains.ResponseError{
		Code:    respCode,
		Message: respMessage,
	})
}
