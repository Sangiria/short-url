package handlers

import (
	"net/http"
	"short_url/jwt"
	"time"

	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {
	token, err := jwt.CreateJwtToken()
	if err != nil{
        return c.String(http.StatusInternalServerError, "something went wrong")
    }

	cookie := &http.Cookie{}
	cookie.Name = "URLCookie"
	cookie.Value = token
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"token":  token,
	})
}

