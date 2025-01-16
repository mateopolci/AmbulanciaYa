package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mateopolci/AmbulanciaYa/src/controllers"
    "github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupAccidenteRoutes(router *gin.Engine, service *services.AccidenteService) {
    accidenteController := controllers.NewAccidenteController(service)
    
    accidente := router.Group("/accidentes")
    {
        accidente.GET("", accidenteController.GetAccidentes)
        accidente.GET("/:id", accidenteController.GetAccidente)
        accidente.POST("", accidenteController.PostAccidente)
        accidente.PUT("/:id", accidenteController.PutAccidente)
        accidente.DELETE("/:id", accidenteController.DeleteAccidente)
    }
}