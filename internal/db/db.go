package db

import (
 "database/sql"
 "fmt"

 _ "github.com/lib/pq" // Драйвер PostgreSQL
)

type DB struct {
 *sql.DB
}

func NewDB(dbURL string) (*DB, error) {
 db, err := sql.Open("postgres", dbURL)
 if err != nil {
  return nil, fmt.Errorf("error opening database connection: %w", err)
 }

 if err := db.Ping(); err != nil {
  return nil, fmt.Errorf("error pinging database: %w", err)
 }

 return &DB{db}, nil
}

func (db *DB) Close() error {
 return db.DB.Close()
}
