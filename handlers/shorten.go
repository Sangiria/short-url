package handlers

import (
	"net/http"
	"short_url/models"

	"github.com/labstack/echo/v4"
)

func GetAllURLs(c echo.Context) error {
	var urls []models.Shorten
	db.Find(&urls)
	return c.JSON(http.StatusOK, urls)
}

func CreateURL(c echo.Context) error {
	url := models.Shorten {
		Token: c.FormValue("token"),
		LongURL: c.FormValue("long_url"),
	}		
    db.Create(&url)
    return c.JSON(http.StatusCreated, url)
}
