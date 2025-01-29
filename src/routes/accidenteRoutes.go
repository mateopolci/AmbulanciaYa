package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupAccidenteRoutes(router *gin.Engine, service *services.AccidenteService) {
    accidenteController := controllers.NewAccidenteController(service)
    
    // Rutas protegidas para paramédicos
    accidenteAuth := router.Group("/accidentes")
    accidenteAuth.Use(middleware.AuthMiddleware())
    {
        accidenteAuth.GET("", accidenteController.GetAccidentes)
        accidenteAuth.GET("/desc", accidenteController.GetAccidentesDesc)
        accidenteAuth.GET("/:id", accidenteController.GetAccidente)
        accidenteAuth.PUT("/:id", accidenteController.PutAccidente)
    }
    
    // Rutas solo para admin
    accidenteAdmin := router.Group("/accidentes")
    accidenteAdmin.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())
    {
        accidenteAdmin.POST("", accidenteController.PostAccidente)
        accidenteAdmin.DELETE("/:id", accidenteController.DeleteAccidente)
        accidenteAdmin.POST("/enviarambulancia", accidenteController.PostAccidenteAndSendAmbulancia)
    }
}