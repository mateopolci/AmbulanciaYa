package main

import (
    "github.com/mateopolci/AmbulanciaYa/src/db"
    "github.com/mateopolci/AmbulanciaYa/src/services"
    "github.com/mateopolci/AmbulanciaYa/src/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    database := db.ConnectNeon()
    accidenteService := services.NewAccidenteService(database)
    
    router := gin.Default()
    routes.SetupAccidenteRoutes(router, accidenteService)
    
    router.Run("localhost:8080")
}