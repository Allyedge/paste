package repositories

import (
	"database/sql"

	"github.com/Allyedge/paste/models"
)

type RepositoryInterface interface {
	Create(*models.Code) error
	FindByID(string) (*models.Code, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Create(code *models.Code) error {
	query := "INSERT INTO paste (id, content, language) VALUES (?, ?, ?)"

	_, err := r.db.Exec(query, code.ID, code.Content, code.Language)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) FindByID(id string) (*models.Code, error) {
	query := "SELECT id, content, language FROM paste WHERE id = $1"

	code := &models.Code{}

	err := r.db.QueryRow(query, id).Scan(&code.ID, &code.Content, &code.Language)

	if err != nil {
		return nil, err
	}

	return code, nil
}
