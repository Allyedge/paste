package handlers

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Allyedge/paste/db/repositories"
	"github.com/Allyedge/paste/models"
	"github.com/Allyedge/paste/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	Database   *sql.DB
	Repository *repositories.Repository
}

func NewHomeHandler(r *repositories.Repository) *HomeHandler {
	return &HomeHandler{
		Repository: r,
	}
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	return render(c, home.Show())
}

func (h HomeHandler) HandleHomeCreate(c echo.Context) error {
	formData := c.FormValue("code")
	language := c.FormValue("language")

	id := generateRandomUrl()

	code := &models.Code{
		ID:       id,
		Content:  formData,
		Language: language,
	}

	err := h.Repository.Create(code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	c.Response().Header().Set("HX-Redirect", "/"+id)

	return c.String(http.StatusOK, "")
}
