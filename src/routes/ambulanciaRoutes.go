package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupAmbulanciaRoutes(router *gin.Engine, service *services.AmbulanciaService) {
	ambulanciaController := controllers.NewAmbulanciaController(service)

	ambulancia := router.Group("/ambulancias")
	{
		ambulancia.GET("", ambulanciaController.GetAmbulancias)
		ambulancia.GET("/desc", ambulanciaController.GetAmbulanciasDesc)
		ambulancia.GET("/:id", ambulanciaController.GetAmbulancia)
		ambulancia.GET("/disp", ambulanciaController.GetAmbulanciaDisponible)
		ambulancia.POST("", ambulanciaController.PostAmbulancia)
		ambulancia.PUT("/:id", ambulanciaController.PutAmbulancia)
		ambulancia.DELETE("/:id", ambulanciaController.DeleteAmbulancia)
	}
}
