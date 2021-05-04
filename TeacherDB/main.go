package main

import (
	"fmt"
	"github.com/labstack/echo"
	"main.go/config"
	controllers "main.go/pkg/db/controller"
	"net/http"
	"os"
)

func main() {

	config.LoadEnvironments()
	fmt.Println("Application started successfully. :)")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Student CRUD APP developed with Golang")
	})
	e.POST("/addteacher", controllers.AddTeacher)
	e.GET("/getteacher", controllers.GetTeacher)
	e.PUT("/updateteacher", controllers.UpdateTeacher)
	e.DELETE("/deleteteacher", controllers.DeleteTeacher)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))

}
