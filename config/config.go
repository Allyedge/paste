package config

import (
	"database/sql"

	"github.com/Allyedge/paste/db/repositories"
	"github.com/Allyedge/paste/handlers"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) InitializeHandlers(r *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(r.Repository)
}

func (c *Config) InitializeRepositories(db *sql.DB) *repositories.Repositories {
	return repositories.NewRepositories(db)
}
