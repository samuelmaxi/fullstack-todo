package router

import (
	loadenv "api-bff-go-gin-todo/load_env"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializerRounter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	initializerRouters(router)

	router.Run(loadenv.LoadEnvFile().ApiBffPort)

}
