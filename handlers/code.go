package handlers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Allyedge/paste/db/repositories"
	"github.com/Allyedge/paste/view/code"
	"github.com/labstack/echo/v4"
)

type CodeHandler struct {
	Database   *sql.DB
	Repository *repositories.Repository
}

func NewCodeHandler(r *repositories.Repository) *CodeHandler {
	return &CodeHandler{
		Repository: r,
	}
}

func (h CodeHandler) HandleCodeShow(c echo.Context) error {
	id := c.Param("id")

	r, err := h.Repository.FindByID(id)
	if err != nil {
		return c.String(404, "Not Found")
	}

	return render(c, code.Show(r.Content, r.Language))
}
