package view

import (
	"backend-tutorial/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes( router *gin.RouterGroup)  {
	router.GET("/users", controllers.GetUsers)
	router.POST("/addUser", controllers.AddUser)
	router.GET("/user/:id", controllers.GetUser)
	router.PATCH("/user/:id", controllers.EditUser)
}