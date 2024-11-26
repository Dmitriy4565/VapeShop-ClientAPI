package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"VapeShop-ClientAPI/internal/config"
)

// Структура для хранения подключения к базе данных
type DB struct {
	*sql.DB
}

// Создание нового подключения к базе данных
func NewDB(cfg config.Database) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.User, cfg.Password, cfg.Name)
	if cfg.Host != "" {
		dsn += fmt.Sprintf(" host=%s", cfg.Host)
	}
	if cfg.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", cfg.Port)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Ошибка ping к базе данных: %w", err)
	}

	return &DB{db}, nil
}

// Закрытие подключения к базе данных
func (db *DB) Close() error {
	return db.DB.Close()
}

// Метод для выполнения запроса
func (db *DB) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.QueryContext(ctx, sql, args...)
}

// Метод для выполнения запроса, возвращающего одну строку
func (db *DB) QueryRowContext(ctx context.Context, sql string, args ...interface{}) *sql.Row {
	return db.DB.QueryRowContext(ctx, sql, args...)
}

// Метод для выполнения команды
func (db *DB) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, sql, args...)
}
