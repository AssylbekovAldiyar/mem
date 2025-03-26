package main

import (
	"fmt"

	"github.com/aldiyar/memegen/internal/config"
	"github.com/aldiyar/memegen/internal/db"
)

func main() {
	fmt.Println("🚀 Starting Meme Generator Service...")

	// Загружаем конфиг
	cfg := config.LoadConfig()

	// Инициализируем БД
	db.InitDB(cfg)

	fmt.Println("✅ Сервис успешно запущен!")
}
