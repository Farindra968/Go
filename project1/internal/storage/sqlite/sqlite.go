package sqlite

import (
	"database/sql"

	"github.com/Farindra968/go_project1/internal/config"
	_ "modernc.org/sqlite"
)

type SqLite struct {
	Db *sql.DB
}

func NewSqLite(cfg config.Config) (*SqLite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
	name STRING NOT NULL,
	email STRING NOT NULL UNIQUE,
	age INTEGER NOT NULL,
	password STRING NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &SqLite{
		Db: db,
	}, nil
}


func (sql *SqLite) CreateStudent(name string, email string, age int, password string) (int64, error) {

	stmt, err :=sql.Db.Prepare(`INSERT INTO users (name, email, age, password) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age, password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}