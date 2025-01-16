package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupParamedicoRoutes(router *gin.Engine, service *services.ParamedicoService) {
	ParamedicoController := controllers.NewParamedicoController(service)

	paramedico := router.Group("/paramedicos")
	{
		paramedico.GET("", ParamedicoController.GetParamedicos)
		paramedico.GET("/:id", ParamedicoController.GetParamedico)
		paramedico.POST("", ParamedicoController.PostParamedico)
		paramedico.PUT("/:id", ParamedicoController.PutParamedico)
		paramedico.DELETE("/:id", ParamedicoController.DeleteParamedico)
	}
}
