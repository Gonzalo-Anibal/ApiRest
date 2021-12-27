package main

import (
	"github.com/Timming/app/infrastructure/timming"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(echoMiddleware.RequestID())
	e.HideBanner = true

	httpClient := timming.NewHttpRepository(&http.Client{})

	userRepository := userRe

	server := &http.Server{
		Addr:         ":8090",
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))

}