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
    hospitalService := services.NewHospitalService(database)
    
    router := gin.Default()
    routes.SetupAccidenteRoutes(router, accidenteService)
    routes.SetupHospitalRoutes(router, hospitalService)
    
    router.Run("localhost:8080")
}