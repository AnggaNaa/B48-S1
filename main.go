package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// fmt.println("Hello Word!")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Word!")
	})

	e.Logger.Fatal(e.Start("localhost:8000"))
}