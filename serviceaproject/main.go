package main

import (
	"github.com/labstack/echo"
	"main.go/config"
	"main.go/controller"
	"os"
)

func main()  {
	config.LoadEnvironments()
	e := echo.New()
	e.GET("/get",controller.Get)
	e.POST("/post",controller.Post)
	e.PUT("/put",controller.Put)
	e.DELETE("/delete",controller.Delete)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}