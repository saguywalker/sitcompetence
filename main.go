package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	admin := e.Group("/api/admin")
	{
		admin.POST("givebadge", api.GiveBadge())
		admin.POST("approveactivity", api.ApproveActivity())
	}

	e.Logger.Fatal(e.Start(":1323"))
}
