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
        ambulancia.POST("/solicitar", ambulanciaController.SolicitarAmbulancia) //ESTA ES LA RUTA DE ENTRADA PARA PEDIR UNA AMBULANCIA
    }
    
    // Rutas protegidas para admin
    ambulanciaAuth := router.Group("/ambulancias")
    ambulanciaAuth.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())
    {
        ambulanciaAuth.GET("", ambulanciaController.GetAmbulancias)
        ambulanciaAuth.GET("/desc", ambulanciaController.GetAmbulanciasDesc)
        ambulanciaAuth.GET("/:id", ambulanciaController.GetAmbulancia)
        ambulanciaAuth.GET("/disp", ambulanciaController.GetAmbulanciaDisponible)
        ambulanciaAuth.POST("", ambulanciaController.PostAmbulancia)
        ambulanciaAuth.PUT("/:id", ambulanciaController.PutAmbulancia)
        ambulanciaAuth.DELETE("/:id", ambulanciaController.DeleteAmbulancia)
    }
}
