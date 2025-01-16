package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupPacienteRoutes(router *gin.Engine, service *services.PacienteService) {
	pacienteController := controllers.NewPacienteController(service)

	paciente := router.Group("/pacientes")
	{
		paciente.GET("", pacienteController.GetPacientes)
		paciente.GET("/:id", pacienteController.GetPaciente)
		paciente.POST("", pacienteController.PostPaciente)
		paciente.PUT("/:id", pacienteController.PutPaciente)
		paciente.DELETE("/:id", pacienteController.DeletePaciente)
	}
}
