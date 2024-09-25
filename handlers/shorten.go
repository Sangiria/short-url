package handlers

import (
	"net/http"
	"short_url/decoder"
	"short_url/models"
	"short_url/validator"
	"github.com/labstack/echo/v4"
)
// GetAllUrls godoc
// @Description  получение всех записей
// @Success      200  {object}  model.Shorten
// @Router /urls [get]
func GetAllURLs(c echo.Context) error {
	var urls []models.Shorten
	db.Find(&urls)
	return c.JSON(http.StatusOK, urls)
}

func DeleteURL(c echo.Context) error {
	var  shorten models.Shorten
	db.Where("id = ?", c.Param("id")).Take(&shorten)
	if shorten.ID == 0 {
		return c.String(http.StatusNotFound, "no item found")
	}
	db.Delete(&shorten)
	return c.String(http.StatusOK, "deleted")
}

func FindURL(c echo.Context) error {
	var shorten models.Shorten
	db.Where("short = ?", c.QueryParam("code")).Take(&shorten)
	if shorten.ID != 0 {
		return c.JSON(http.StatusOK, shorten.LongURL)
	} else {
		return c.String(http.StatusNotFound, "error")
	}
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