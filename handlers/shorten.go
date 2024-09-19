package handlers

import (
	"net/http"
	"short_url/decoder"
	"short_url/models"
	"short_url/validator"
	"github.com/labstack/echo/v4"
)

func GetAllURLs(c echo.Context) error {
	var urls []models.Shorten
	db.Find(&urls)
	return c.JSON(http.StatusOK, urls)
}

func CreateURL(c echo.Context) error {
	url := models.Shorten {
		Short: c.FormValue("hash"),
		LongURL: c.FormValue("long_url"),
	}		
    db.Create(&url)
    return c.JSON(http.StatusCreated, url)
}

func CreateShortURL(c echo.Context) error {
	input := c.FormValue("input")

	if validator.ValidateURL(input){
		var shorten models.Shorten
		db.Where("long_url = ?", input).Last(&shorten)
		if shorten.ID == 0 {
			i := 0
			for i < 4 {
				result := decoder.DecodeURL()
				db.Where("short = ?", result).First(&shorten)
				if shorten.Short != result {
					i = 4
					new_shorten := models.Shorten{
                        Short: result,
                        LongURL: input,
                    }
					db.Create(&new_shorten)
					return c.JSON(http.StatusCreated, new_shorten.Short)
				} else {
					i++
				}
			}
			c.String(http.StatusNotFound, "error")			
		} else {
			return c.JSON(http.StatusOK, shorten.Short)
		}
	}
	return c.String(http.StatusNotFound, "error")
}