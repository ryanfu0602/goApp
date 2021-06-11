package main

import (
	"goApp/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//Log & Recover
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//public API
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to golang!")
	})

	e.POST("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Succ!")
	})

	//auth API
	r := e.Group("/auth")
	{
		// config := middleware.JWTConfig{
		// 	KeyFunc: getKey,
		// }
		// r.Use(middleware.JWTWithConfig(config))
		r.GET("/articles/:id", controllers.GetArticleById)
		r.POST("/articles", controllers.AddArticle)
		r.PUT("/articles/:id", controllers.UpdateArticleById)
		r.DELETE("/articles/:id", controllers.DeleteArticleById)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
