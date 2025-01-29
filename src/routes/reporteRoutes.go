package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupReporteRoutes(router *gin.Engine, service *services.ReporteService) {
    reporteController := controllers.NewReporteController(service)
    
    // Rutas protegidas para paramédicos
    reporteAuth := router.Group("/reportes")
    reporteAuth.Use(middleware.AuthMiddleware())
    {
        reporteAuth.GET("", reporteController.GetReportes)
        reporteAuth.GET("/:id", reporteController.GetReporte)
        reporteAuth.GET("/accidente/:accidenteId", reporteController.GetReporteByAccidente)
        reporteAuth.POST("", reporteController.PostReporte)
        reporteAuth.POST("/accidente/:accidenteId", reporteController.CreateReporteAndUpdateHospital)
        reporteAuth.PUT("/:id", reporteController.PutReporte)
        reporteAuth.PUT("/accidente/:reporteId", reporteController.UpdateReporteAndHospital)
        reporteAuth.DELETE("/:id", reporteController.DeleteReporte)
        reporteAuth.DELETE("/accidente/:accidenteId", reporteController.DeleteReporteByAccidente)
    }
}