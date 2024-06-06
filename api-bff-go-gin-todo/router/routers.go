package router

import (
	"api-bff-go-gin-todo/handlers"

	"github.com/gin-gonic/gin"
)

func initializerRouters(router *gin.Engine) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/task/list", handlers.GetTaskList)
		v1.GET("/task/search/:id", handlers.GetTaskSearch)
		v1.POST("/task/create", handlers.PostTaskCreate)
		v1.PATCH("/task/edit/:id", handlers.PatchTaskEdit)
		v1.DELETE("/task/delete/:id", handlers.DeleteTaskDelete)
	}

}
