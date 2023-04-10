package main

import (
	"net/http"

	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db.ConnectMongo()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, Shadow!",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
