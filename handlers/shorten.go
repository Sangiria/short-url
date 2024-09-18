package handlers

import (
	"net/http"
	"short_url/decoder"
	"short_url/models"
	"short_url/utils"

	"github.com/labstack/echo/v4"
)

func GetAllURLs(c echo.Context) error {
	var urls []models.Shorten
	db.Find(&urls)
	return c.JSON(http.StatusOK, urls)
}

func CreateURL(c echo.Context) error {
	url := models.Shorten {
		Hash: utils.StringToUint32(c.FormValue("hash")),
		LongURL: c.FormValue("long_url"),
	}		
    db.Create(&url)
    return c.JSON(http.StatusCreated, url)
}

func CreateShortURL(c echo.Context) error {
	input := c.FormValue("input")
	var shorten models.Shorten
	db.Where("long_url = ?", input).Last(&shorten)
	if shorten.ID == 0 {
		ok := false
		var result uint32
		for !ok {
			result = decoder.DecodeURL(input)
			db.Where("hash = ?", result).First(&shorten)
			if shorten.Hash != result {
				ok = true
				new_shorten := models.Shorten{
					Hash: result,
                    LongURL: input,
				}
				db.Create(&new_shorten)
				return c.JSON(http.StatusCreated, new_shorten.Hash)
			} else {
				ok = false
			}
		}
	} else {
		return c.JSON(http.StatusOK, shorten.Hash)
	}
	return c.String(http.StatusNotFound, "error")
}