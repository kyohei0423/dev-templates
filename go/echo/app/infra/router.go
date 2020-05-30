package infra

// TODO: Fix import path
import (
	"github.com/kyohei0423/dev-templates/go/echo/app/interface/controller"
	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
	e.Debug = true // tmp true
	userController := controller.NewSampleController()
	e.GET("/", userController.Hello)
	e.Logger.Fatal(e.Start(":8080"))
}
