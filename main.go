package main

import (
	"backend-tutorial/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	// r.Group(view.UserRoutes)

	r.GET("/users", controllers.GetUsers)

	r.GET("/user/:id", controllers.GetUser)

	r.PATCH("/user/:id",controllers.EditUser)

	r.POST("/addUser", controllers.AddUser )

	r.Run()
}
