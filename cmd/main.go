package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Start(":4000")
}
