package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateopolci/AmbulanciaYa/src/db"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/routes"
	"github.com/mateopolci/AmbulanciaYa/src/services"
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
		gin.SetMode(gin.ReleaseMode) // Mejora de rendimiento y menos logs de debug
	}

	// Carga de la secret key
	initAuth()

	// Conexión a bd
	database := db.ConnectNeon()

	// Construcción de servicios
	accidenteService := services.NewAccidenteService(database)
	hospitalService := services.NewHospitalService(database)
	pacienteService := services.NewPacienteService(database)
	paramedicoService := services.NewParamedicoService(database)
	choferService := services.NewChoferService(database)
	reporteService := services.NewReporteService(database)
	ambulanciaService := services.NewAmbulanciaService(database, pacienteService, accidenteService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:5175",
			"http://localhost:5176",
			"http://localhost:5177",
			"http://localhost:5178",
			"http://localhost:5179",
			"http://localhost:5190",
			"https://ambulanciaya.onrender.com",
			"http://localhost:8080",
			"https://ambulancia-ya.vercel.app",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"credentials",
			"Access-Control-Allow-Credentials",
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup de grupos de rutas
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
