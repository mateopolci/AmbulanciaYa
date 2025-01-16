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
    pacienteService := services.NewPacienteService(database)
    
    router := gin.Default()

    routes.SetupAccidenteRoutes(router, accidenteService)
    routes.SetupHospitalRoutes(router, hospitalService)
    routes.SetupPacienteRoutes(router, pacienteService)
    
    router.Run("localhost:8080")
}