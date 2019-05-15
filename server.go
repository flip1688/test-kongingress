package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/sub-path", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is sub path")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
