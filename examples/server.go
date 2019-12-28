package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func get(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"method":  "GET",
			"path":    c.Path(),
			"payload": c.QueryString(),
		},
	)
}

func post(c echo.Context) error {
	m := map[string]interface{}{}
	if err := c.Bind(&m); err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"method":  "POST",
			"path":    c.Path(),
			"payload": m,
		},
	)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/get", get)
	e.POST("/post", post)

	e.Logger.Fatal(e.Start(":8000"))
}
