package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupChoferRoutes(router *gin.Engine, service *services.ChoferService) {
	ChoferController := controllers.NewChoferController(service)

	chofer := router.Group("/choferes")
	{
		chofer.GET("", ChoferController.GetChoferes)
		chofer.GET("/:id", ChoferController.GetChofer)
		chofer.POST("", ChoferController.PostChofer)
		chofer.PUT("/:id", ChoferController.PutChofer)
		chofer.DELETE("/:id", ChoferController.DeleteChofer)
	}
}
