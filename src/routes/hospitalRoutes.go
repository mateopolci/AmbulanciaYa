package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupHospitalRoutes(router *gin.Engine, service *services.HospitalService) {
    hospitalController := controllers.NewHospitalController(service)
    
    // Rutas protegidas para paramedico
    hospitalAuth := router.Group("/hospitales")
    hospitalAuth.Use(middleware.AuthMiddleware())

    {
        hospitalAuth.GET("", hospitalController.GetHospitales)
        hospitalAuth.GET("/:id", hospitalController.GetHospital)
    }

    // Rutas protegidas para admin
    hospitalAdmin := router.Group("/hospitales")
    hospitalAdmin.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())

    {
        hospitalAdmin.POST("", hospitalController.PostHospital)
        hospitalAdmin.PUT("/:id", hospitalController.PutHospital)
        hospitalAdmin.DELETE("/:id", hospitalController.DeleteHospital)
    }
}