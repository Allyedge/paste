package repositories

import "database/sql"

type Repositories struct {
	Repository *Repository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Repository: NewRepository(db),
	}
}
