package main

import (
	"net/http"
	"short_url/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
 
func main() {	
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shorter.danyatochka.ru", "http://localhost:4200", "https://shorter.danyatochka.ru/swagger/index.html"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true, 
	  }))
	
	e.GET("/urls", handlers.GetAllURLs)
	e.GET("/url", handlers.FindURL)
	e.POST("/short_url", handlers.CreateShortURL)
	e.DELETE("url/:id", handlers.DeleteURL)
	e.Logger.Fatal(e.Start(":1323"))


}
