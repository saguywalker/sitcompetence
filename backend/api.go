package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func main(){
	fmt.Println("main")
	e := echo.New()
	
	e.GET("/competence", func(c echo.Context) error{
		return c.JSON(200, "GET competencies")
	})
	e.PUT("/competence", func(c echo.Context) error{
		return c.JSON(200, "PUT competencies")
	})
	e.DELETE("/competence/:id", func(c echo.Context) error{
		return c.JSON(200, "DELETE competence: " + c.Param("id"))
	})

	e.Run(standard.New(":8000"))
}