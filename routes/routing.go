package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}

// Creo un routing unico per accedere a tutti i tipi di dominio
func DomainRoute(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/:domain", nil)
		v1.GET("/:domain/:domainId", nil)
		v1.POST("/:domain", nil)
		v1.PUT("/:domain", nil)
		v1.DELETE("/:domain", nil)
	}
}
