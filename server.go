package main

import (
	"short_url/handlers"
	"github.com/labstack/echo/v4"
)


func main() {	
	e := echo.New()
	e.GET("/urls", handlers.GetAllURLs)
	e.POST("/url", handlers.CreateURL)
	e.Logger.Fatal(e.Start(":1323"))
}
