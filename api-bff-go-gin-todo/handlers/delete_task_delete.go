package handlers

import (
	loadenv "api-bff-go-gin-todo/load_env"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTaskDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	// var task models.Task

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", loadenv.LoadEnvFile().ApiBackDeleteDelete+id+"/", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})

}
