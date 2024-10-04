package handlers

import (
	"net/http"
	"short_url/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {
	token, err := jwt.CreateJwtToken()
	if err != nil{
        return c.String(http.StatusInternalServerError, "something went wrong")
    }

	return c.JSON(http.StatusOK, map[string]string{
		"token":  token,
	})
}

