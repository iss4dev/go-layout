package app

import (
	"fmt"
	"main/internal/app/endpoint"
	"main/internal/app/mw"
	"main/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	s    *service.Service
	e    *endpoint.Endpoint
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server starting")
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
	return nil
}
