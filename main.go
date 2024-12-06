package main

import (
	"VapeShop-ClientAPI/internal/config"
	"VapeShop-ClientAPI/internal/db"
	"VapeShop-ClientAPI/internal/router" // Импортируем новый пакет router
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	server := router.NewServer(cfg, db) // Передаем конфигурацию и базу данных
	if err := server.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
