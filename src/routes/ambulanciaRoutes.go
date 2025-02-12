package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupAmbulanciaRoutes(router *gin.Engine, service *services.AmbulanciaService) {
    ambulanciaController := controllers.NewAmbulanciaController(service)

    // Ruta publica 
    ambulancia := router.Group("/ambulancias")
    {
        ambulancia.POST("/solicitar", ambulanciaController.SolicitarAmbulancia) //Ruta de entrada para el pedido de una ambulancia
    }

    // Rutas protegidas para paramedico
    ambulanciaAuth := router.Group("/ambulancias")
    ambulanciaAuth.Use(middleware.AuthMiddleware())
    {
        ambulanciaAuth.GET("", ambulanciaController.GetAmbulancias)
        ambulanciaAuth.GET("/desc", ambulanciaController.GetAmbulanciasDesc)
        ambulanciaAuth.GET("/:id", ambulanciaController.GetAmbulancia)
    }

    // Rutas protegidas para admin
    ambulanciaAdmin := router.Group("/ambulancias")
    ambulanciaAdmin.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())
    {
        ambulanciaAdmin.GET("/disp", ambulanciaController.GetAmbulanciaDisponible)
        ambulanciaAdmin.POST("", ambulanciaController.PostAmbulancia)
        ambulanciaAdmin.PUT("/:id", ambulanciaController.PutAmbulancia)
        ambulanciaAdmin.DELETE("/:id", ambulanciaController.DeleteAmbulancia)
    }
}
