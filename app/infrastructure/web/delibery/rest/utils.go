package rest

import "github.com/labstack/echo"

func ErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}

