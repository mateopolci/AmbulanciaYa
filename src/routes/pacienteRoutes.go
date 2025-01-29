package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupPacienteRoutes(router *gin.Engine, service *services.PacienteService) {
	pacienteController := controllers.NewPacienteController(service)

	// Rutas protegidas para paramédicos
	pacienteAuth := router.Group("/pacientes")
	pacienteAuth.Use(middleware.AuthMiddleware())
	{
		pacienteAuth.GET("", pacienteController.GetPacientes)
		pacienteAuth.GET("/:id", pacienteController.GetPaciente)
		pacienteAuth.GET("/telefono/:telefono", pacienteController.GetByTelefono)
		pacienteAuth.POST("", pacienteController.PostPaciente)
		pacienteAuth.PUT("/:id", pacienteController.PutPaciente)
		pacienteAuth.DELETE("/:id", pacienteController.DeletePaciente)
	}
}
