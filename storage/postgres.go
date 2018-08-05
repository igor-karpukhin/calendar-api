package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(addr, user, password, dbname string) (*sql.DB, error) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user, password, addr, dbname)

	conn, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
