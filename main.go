package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/db"
	"github.com/mateopolci/AmbulanciaYa/src/routes"
	"github.com/mateopolci/AmbulanciaYa/src/services"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
)

func initAuth() {
    _ = godotenv.Load()
    
    secretKey := os.Getenv("SECRET_KEY")
    if secretKey == "" {
        log.Fatal("ERROR: SECRET_KEY is not set")
    } 
    middleware.SetSecretKey(secretKey)
}
func main() {

	if os.Getenv("GIN_MODE") == "release" {
        gin.SetMode(gin.ReleaseMode)
    }

    initAuth()
    
	database := db.ConnectNeon()
	accidenteService := services.NewAccidenteService(database)
	hospitalService := services.NewHospitalService(database)
	pacienteService := services.NewPacienteService(database)
	paramedicoService := services.NewParamedicoService(database)
	choferService := services.NewChoferService(database)
	reporteService := services.NewReporteService(database)
	ambulanciaService := services.NewAmbulanciaService(database)

	router := gin.Default()
	router.Use(middleware.SetupCORS())

	routes.SetupAccidenteRoutes(router, accidenteService)
	routes.SetupHospitalRoutes(router, hospitalService)
	routes.SetupPacienteRoutes(router, pacienteService)
	routes.SetupParamedicoRoutes(router, paramedicoService)
	routes.SetupChoferRoutes(router, choferService)
	routes.SetupReporteRoutes(router, reporteService)
	routes.SetupAmbulanciaRoutes(router, ambulanciaService)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router.Run("0.0.0.0:" + port)
}