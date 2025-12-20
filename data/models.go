package data

import (
	"database/sql"
)

// Models holds all application data models.
// Add your models here after generating them with: gq make model <name>
type Models struct {
	// Add model fields here, e.g.:
	// Users UserModel
}

// New creates a new Models instance with the given database pool.
// Returns the Models and any error that occurred during setup.
func New(databasePool *sql.DB) (Models, error) {
	return Models{}, nil
}
