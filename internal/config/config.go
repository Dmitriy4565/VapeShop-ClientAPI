package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	JWT           JWT           `mapstructure:",squash"`
	Database      Database      `mapstructure:",squash"`
	ServerPort    int           `mapstructure:"SERVER_PORT"`
	ServerAddress string        `mapstructure:"SERVER_ADDRESS"`
	Debug         bool          `mapstructure:"DEBUG"`
	CORS          CORS          `mapstructure:",squash"`
	Timeout       time.Duration `mapstructure:"SERVER_TIMEOUT"`
}

type JWT struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

type Database struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
}

type CORS struct {
	Origins []string `mapstructure:"CORS_ORIGINS"`
	Methods []string `mapstructure:"CORS_METHODS"`
	Headers []string `mapstructure:"CORS_HEADERS"`
}

// Функция для загрузки конфигурации из файла .env или переменных окружения
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("DATABASE_SSLMODE", "disable") // Установите значение по умолчанию
	viper.SetDefault("SERVER_TIMEOUT", "15s")       // Добавлено значение по умолчанию

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфигурации: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Ошибка разбора конфигурации: %w", err)
	}
	fmt.Printf("Загруженная конфигурация: %+v\n", cfg)
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
