package main

import (
	"log"
	"os"

	"github.com/HubPavKul1/vetstore2025/internal/db"
	"github.com/HubPavKul1/vetstore2025/internal/ui"
	// "fyne.io/fyne/v2/app"
)

func main() {
    // Инициализируем базу данных
    err := db.InitDB("vetstore2025.db") // Передаем путь к файлу базы данных
    if err != nil {
        log.Fatalf("database initialization failed: %v", err)
    }

    // Закрываем базу данных после завершения работы приложения
    defer db.Close()

    ui.RunUI()

    os.Exit(0)
}