package handler

import (
	"net/http"

	"github.com/hysem/charts/templates"
	"github.com/labstack/echo/v4"
)

// ShowDashboard shows the dashboard
func ShowDashboard(c echo.Context) error {
	response := templates.DefaultParams(c)
	return c.Render(http.StatusOK, "dashboard", response)
}
