package db

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectNeon() *gorm.DB {
    // Cargar variables de entorno
    _ = godotenv.Load()


    // Obtener la URL de conexión
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL no está configurada")
    }

    // Conectar a la base de datos usando GORM
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatal("Error conectando a la base de datos: ", err)
    }
    
    return db
}

