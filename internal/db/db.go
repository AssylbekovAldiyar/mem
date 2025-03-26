package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aldiyar/memegen/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Ошибка при проверке соединения: %v", err)
	}

	fmt.Println("✅ Подключение к БД установлено!")
}
