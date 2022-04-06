package router

import "github.com/labstack/echo/v4"

type Router struct {
	*echo.Echo
}

func NewRouter() Router {
	return Router{
		echo.New(),
	}
}
