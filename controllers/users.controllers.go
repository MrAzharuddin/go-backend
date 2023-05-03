package controllers

import (
	"backend-tutorial/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []model.User{
	{UserId: 1, Name: "Sravani", Mobile: 86865144554},
	{UserId: 2, Name: "Azar", Mobile: 8686516714},
	{UserId: 3, Name: "Pratyusha", Mobile: 4215132156},
}



func GetUsers(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message":"success",
		"data":users,
	})
}

func GetUser(ctx *gin.Context){
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		for _, value := range users {
			if id == value.UserId {
				ctx.JSON(http.StatusAccepted, gin.H{
					"data": value,
				})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"data": ctx.Param("id") + "Not found",
		})
}

func EditUser(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for index, value := range users {
		if id == value.UserId {
			users[index].Name = user.Name
			ctx.JSON(http.StatusAccepted, gin.H{
				"data": users,
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"data": ctx.Param("id") + "Not found",
	})
}

func AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	users = append(users, user)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": users,
	})
}