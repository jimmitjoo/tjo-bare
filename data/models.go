package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

type Models struct {
}

// New creates a new Models instance with the given database pool.
// Returns the Models and any error that occurred during setup.
func New(databasePool *sql.DB) (Models, error) {
	db = databasePool

	var err error
	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, err = mysql.New(db)
		if err != nil {
			return Models{}, fmt.Errorf("failed to initialize MySQL adapter: %w", err)
		}
	case "postgres", "postgresql":
		upper, err = postgresql.New(db)
		if err != nil {
			return Models{}, fmt.Errorf("failed to initialize PostgreSQL adapter: %w", err)
		}
	default:
		log.Println("Warning: No DATABASE_TYPE set, database operations may not work")
	}

	return Models{}, nil
}

func getInsertID(i db2.ID) int {

	if i == nil {
		return 0
	}

	idType := fmt.Sprintf("%T", i)

	if idType == "int64" {
		return int(i.(int64))
	} else if idType == "func() db.ID" {
		return getInsertID(i.(func() db2.ID)())
	}

	return int(i.(int))
}
