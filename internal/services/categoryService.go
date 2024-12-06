package services

import (
	"context"
	"database/sql"
	"errors"
)

type Category struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	StoreID *int64 `json:"store_id"` // Используем sql.NullInt64 для поддержки NULL
}

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetCategoryByID(ctx context.Context, id string) (*Category, error)
	CreateCategory(ctx context.Context, category Category) (*Category, error)
	UpdateCategory(ctx context.Context, category Category) (*Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type CategoryServiceImpl struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		db: db,
	}
}

func (s *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, store_id FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		var storeID sql.NullInt64 // Временная переменная для считывания значения из базы данных

		if err := rows.Scan(&category.ID, &category.Name, &storeID); err != nil {
			return nil, err
		}

		// Присваиваем значение StoreID
		if storeID.Valid {
			category.StoreID = new(int64) // Создаем новый int64 и присваиваем значение
			*category.StoreID = storeID.Int64
		} else {
			category.StoreID = nil // Устанавливаем в nil, если значение NULL в базе данных
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryServiceImpl) GetCategoryByID(ctx context.Context, id string) (*Category, error) {
	var category Category

	// Выполнение запроса с явным указанием столбцов
	err := s.db.QueryRowContext(ctx, "SELECT id, name, store_id FROM categories WHERE id = $1", id).Scan(
		&category.ID,
		&category.Name,
		&category.StoreID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("категория не найдена")
		}
		return nil, err // Возврат ошибки при других проблемах
	}

	return &category, nil // Возврат найденной категории
}

func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, category Category) (*Category, error) {
	// Используем RETURNING для получения ID вставленной записи
	query := `
        INSERT INTO categories (name, store_id) 
        VALUES ($1, $2) 
        RETURNING id`

	// Выполняем запрос и сканируем возвращаемый ID в структуру category
	err := s.db.QueryRowContext(ctx, query, category.Name, category.StoreID).Scan(&category.ID)

	if err != nil {
		return nil, err // Возврат ошибки при выполнении запроса
	}

	return &category, nil // Возврат созданной категории с установленным ID
}

func (s *CategoryServiceImpl) UpdateCategory(ctx context.Context, category Category) (*Category, error) {
	// Проверка существования категории
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM categories WHERE id=$1)", category.ID).Scan(&exists)
	if err != nil {
		return nil, errors.New("ошибка при проверке существования категории")
	}
	if !exists {
		return nil, errors.New("категория не найдена")
	}

	// Обновление категории
	query := `
        UPDATE categories 
        SET name = $1, store_id = $2 
        WHERE id = $3`

	_, err = s.db.ExecContext(ctx, query, category.Name, category.StoreID, category.ID)
	if err != nil {
		return nil, errors.New("ошибка при обновлении категории")
	}

	return &category, nil // Возврат обновленной категории с установленным ID
}

func (s *CategoryServiceImpl) DeleteCategory(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM categories WHERE id = $1", id)
	return err
}
