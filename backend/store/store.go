package store

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose"
	"log"
	"os"
	"path/filepath"
)

type store struct {
	db *sql.DB
}

func NewStore(conn string, driver string) (*store, error) {
	db, err := sql.Open(driver, conn)

	if err != nil {
		return nil, err
	}

	migrations(db)
	return &store{db: db}, nil

}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Ping() error {
	return s.db.Ping()
}

func migrations(db *sql.DB) {
	migrationsDir := "migrations"
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	migrationsDirPath := filepath.Join(currentDir, migrationsDir)
	goose.SetDialect("postgres")
	err = goose.Up(db, migrationsDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
