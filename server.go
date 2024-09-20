package main

import (
	"short_url/handlers"
	"github.com/labstack/echo/v4"
)


func main() {	
	e := echo.New()
	e.GET("/urls", handlers.GetAllURLs)
	e.POST("/short_url", handlers.CreateShortURL)
	e.DELETE("url/:id", handlers.DeleteURL)
	e.Logger.Fatal(e.Start(":1323"))


}
