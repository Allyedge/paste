package handlers

import (
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func generateRandomUrl() string {
	id := uuid.New()

	return id.String()
}
