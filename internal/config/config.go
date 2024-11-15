package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Структура для JWT-секрета
type JWT struct {
	Secret string `json:"jwt_secret"`
}

// Структура для конфигурации базы данных
type Database struct {
	Host             string `json:"database_host"`
	Port             int    `json:"database_port"`
	User             string `json:"database_user"`
	Password         string `json:"database_password"`
	Name             string `json:"database_name"`
	SSLMode          string `yaml:"sslmode"`           // Добавьте sslmode, если нужно
	ConnectionString string `yaml:"connection_string"` // или строка подключения целиком
}

// Структура для CORS-конфигурации
type CORS struct {
	Origins []string `json:"cors_origins"`
	Methods []string `json:"cors_methods"`
	Headers []string `json:"cors_headers"`
}

// Структура для общей конфигурации приложения
type Config struct {
	JWT           JWT      `json:"jwt"`
	Database      Database `json:"database"`
	ServerPort    int      `json:"server_port"`
	ServerAddress string   `json:"server_address"`
	Debug         bool     `json:"debug"`
	CORS          CORS     `json:"cors"`
	DB            Database `yaml:"db"`
}

// Функция для загрузки конфигурации из файла .env или переменных окружения
func LoadConfig() (*Config, error) {
	// Устанавливаем файл конфигурации
	viper.SetConfigFile(".env")

	// Автоматически читаем переменные окружения
	viper.AutomaticEnv()

	// Читаем конфигурацию
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфигурации: %w", err)
	}

	// Создаем структуру конфигурации
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Ошибка разбора конфигурации: %w", err)
	}

	return &cfg, nil
}

// Функция для валидации конфигурации (дополнительная)
func (c *Config) Validate() error {
	// Дополнительная валидация конфигурации, если нужно
	// Например, проверить, что SERVER_PORT > 0
	if c.ServerPort <= 0 {
		return fmt.Errorf("неверный порт сервера: %d", c.ServerPort)
	}

	return nil
}

// Функция для получения значения конфигурации по ключу
func GetConfigValue(key string, v interface{}) error {
	return viper.UnmarshalKey(key, v)
}

// Функция для получения значения конфигурации по ключу в виде строки
func GetConfigString(key string) string {
	var value string
	if err := GetConfigValue(key, &value); err != nil {
		fmt.Println("Ошибка получения значения конфигурации:", err)
	}
	return value
}

// Функция для получения значения конфигурации по ключу в виде целого числа
func GetConfigInt(key string) int {
	var value int
	if err := GetConfigValue(key, &value); err != nil {
		fmt.Println("Ошибка получения значения конфигурации:", err)
		return 0
	}
	return value
}

// Функция для получения значения конфигурации по ключу в виде булевого значения
func GetConfigBool(key string) bool {
	var value bool
	if err := GetConfigValue(key, &value); err != nil {
		fmt.Println("Ошибка получения значения конфигурации:", err)
		return false
	}
	return value
}

// Функция для получения значения конфигурации по ключу в виде времени
func GetConfigDuration(key string) time.Duration {
	var value time.Duration
	if err := GetConfigValue(key, &value); err != nil {
		fmt.Println("Ошибка получения значения конфигурации:", err)
		return 0
	}
	return value
}
