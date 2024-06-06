package router

import (
	loadenv "api-bff-go-gin-todo/load_env"

	"github.com/gin-gonic/gin"
)

func InitializerRounter() {
	router := gin.Default()

	initializerRouters(router)

	router.Run(loadenv.LoadEnvFile().ApiBffPort)

}
