package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mateopolci/AmbulanciaYa/src/controllers"
    "github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupReporteRoutes(router *gin.Engine, service *services.ReporteService) {
    reporteController := controllers.NewReporteController(service)
    
    accidente := router.Group("/reportes")
    {
        accidente.GET("", reporteController.GetReportes)
        accidente.GET("/:id", reporteController.GetReporte)
        accidente.GET("/accidente/:accidenteId", reporteController.GetReporteByAccidente)
        accidente.POST("", reporteController.PostReporte)
        accidente.POST("/accidente/:accidenteId", reporteController.CreateReporteAndUpdateHospital)
        accidente.PUT("/:id", reporteController.PutReporte)
        accidente.PUT("/accidente/:reporteId", reporteController.UpdateReporteAndHospital)
        accidente.DELETE("/:id", reporteController.DeleteReporte)
    }
}