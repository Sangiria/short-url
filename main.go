package main

import (
	"net/http"
	"short_url/enviroment"
	"short_url/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
 
func main() {	
	e := echo.New()
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shorter.danyatochka.ru", "http://localhost:4200"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true, 
	  }))

	url := e.Group("/jwt")

	url.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey: []byte(enviroment.GoDotEnvVariable("Secretkey")),
	}))
		
	e.GET("/", handlers.Auth)
	url.GET("/urls", handlers.GetAllURLs)
	url.GET("/url", handlers.FindURL)
	url.POST("/url", handlers.CreateShortURL)
	url.DELETE("/url/:id", handlers.DeleteURL)
	e.Logger.Fatal(e.Start(":1323"))


}
