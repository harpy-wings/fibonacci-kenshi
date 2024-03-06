package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	Register(*echo.Echo) error
}
