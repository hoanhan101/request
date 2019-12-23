package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func test(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]string{
			"status": "ok",
		},
	)

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":1323"))
}
