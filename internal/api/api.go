package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API hold echo instance
type API struct {
	*echo.Echo
}

// New returns echo
func New() *API {
	e := echo.New()
	a := &API{
		Echo: e,
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5000",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	}))
	e.GET("/", greeting)

	return a
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
