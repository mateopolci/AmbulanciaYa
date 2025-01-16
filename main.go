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
    paramedicoService := services.NewParamedicoService(database)
    choferService := services.NewChoferService(database)
    reporteService := services.NewReporteService(database)
    ambulanciaService := services.NewAmbulanciaService(database)
    
    router := gin.Default()

    routes.SetupAccidenteRoutes(router, accidenteService)
    routes.SetupHospitalRoutes(router, hospitalService)
    routes.SetupPacienteRoutes(router, pacienteService)
    routes.SetupParamedicoRoutes(router, paramedicoService)
    routes.SetupChoferRoutes(router, choferService)
    routes.SetupReporteRoutes(router, reporteService)
    routes.SetupAmbulanciaRoutes(router, ambulanciaService)
    
    router.Run("localhost:8080")
}