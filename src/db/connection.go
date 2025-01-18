package db

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectNeon() *gorm.DB {
    // Intentar cargar .env en desarrollo
    _ = godotenv.Load()

    // Obtener la URL de conexión
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL no está configurada")
    }

    // Configuración para desarrollo y producción
    config := postgres.Config{
        DSN: connStr,
        PreferSimpleProtocol: true,
    }

    // Conectar a la base de datos con GORM
    db, err := gorm.Open(postgres.New(config), &gorm.Config{})
    if err != nil {
        log.Fatal("Error conectando a la base de datos: ", err)
    }

    fmt.Println("Conexión a la base de datos establecida")
    return db
}