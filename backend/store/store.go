package store

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type store struct {
	db *sql.DB
}

func NewStore(conn string, driver string) (*store, error) {
	db, err := sql.Open(driver, conn)

	if err != nil {
		return nil, err
	}

	return &store{db: db}, nil

}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Ping() error {
	return s.db.Ping()
}
