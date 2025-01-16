package db

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
/*     "github.com/mateopolci/AmbulanciaYa/src/models" */
)

func ConnectNeon() *gorm.DB {
    // Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Obtener la URL de conexión
    connStr := os.Getenv("DATABASE_URL")

    // Conectar a la base de datos usando GORM
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    return db
}

