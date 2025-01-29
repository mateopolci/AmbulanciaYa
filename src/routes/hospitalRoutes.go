package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupHospitalRoutes(router *gin.Engine, service *services.HospitalService) {
    hospitalController := controllers.NewHospitalController(service)
    
    // Todas las rutas de hospital requieren ser admin
    hospital := router.Group("/hospitales")
    hospital.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())

    {
        hospital.GET("", hospitalController.GetHospitales)
        hospital.GET("/:id", hospitalController.GetHospital)
        hospital.POST("", hospitalController.PostHospital)
        hospital.PUT("/:id", hospitalController.PutHospital)
        hospital.DELETE("/:id", hospitalController.DeleteHospital)
    }
}