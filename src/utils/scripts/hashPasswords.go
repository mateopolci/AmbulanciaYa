package main

import (
    "log"
    "os"
    "path/filepath"
    "golang.org/x/crypto/bcrypt"
    "github.com/joho/godotenv"
    "github.com/mateopolci/AmbulanciaYa/src/db"
    "github.com/mateopolci/AmbulanciaYa/src/models"
)

func main() {
    // Get current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatal("Error getting working directory:", err)
    }

    // Navigate up to project root (3 levels up from /src/utils/scripts)
    projectRoot := filepath.Join(currentDir, "..", "..", "..")
    
    // Load .env from project root
    err = godotenv.Load(filepath.Join(projectRoot, ".env"))
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    // Verify DATABASE_URL exists
    if os.Getenv("DATABASE_URL") == "" {
        log.Fatal("DATABASE_URL must be set in .env file")
    }

    database := db.ConnectNeon()
    var paramedicos []models.Paramedico
    
    if err := database.Find(&paramedicos).Error; err != nil {
        log.Fatal("Error fetching paramedicos:", err)
    }
    
    for _, p := range paramedicos {
        hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
        if err != nil {
            log.Printf("Error hashing password for paramedico %s: %v", p.Id, err)
            continue
        }
        
        if err := database.Model(&p).Update("password", string(hash)).Error; err != nil {
            log.Printf("Error updating paramedico %s: %v", p.Id, err)
        }
    }
    
    log.Println("Password hashing completed successfully")
}