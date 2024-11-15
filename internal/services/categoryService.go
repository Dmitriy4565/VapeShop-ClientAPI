package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetCategoryByID(ctx context.Context, id string) (*Category, error)
	CreateCategory(ctx context.Context, category Category) (*Category, error)
	UpdateCategory(ctx context.Context, category Category) (*Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type CategoryServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewCategoryService(db *sql.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		db: db,
	}
}

func (s *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *CategoryServiceImpl) GetCategoryByID(ctx context.Context, id string) (*Category, error) {
	var category Category
	err := s.db.QueryRowContext(ctx, "SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("категория не найдена")
		}
		return nil, err
	}
	return &category, nil
}

func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, category Category) (*Category, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	category.ID = fmt.Sprintf("%d", lastInsertID)
	return &category, nil
}

func (s *CategoryServiceImpl) UpdateCategory(ctx context.Context, category Category) (*Category, error) {
	res, err := s.db.ExecContext(ctx, "UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("категория не найдена") // Обработка случая, когда категория не найдена
	}

	return &category, nil // Возвращаем обновлённую категорию
}

func (s *CategoryServiceImpl) DeleteCategory(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM categories WHERE id = $1", id)
	return err
}
