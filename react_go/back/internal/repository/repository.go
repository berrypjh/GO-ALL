package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseReop interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}