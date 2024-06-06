package handlers

import (
	loadenv "api-bff-go-gin-todo/load_env"
	"api-bff-go-gin-todo/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskList(ctx *gin.Context) {
	resp, err := http.Get(loadenv.LoadEnvFile().ApiBackGetList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var taskList []models.Task
	err = json.Unmarshal(body, &taskList)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, taskList)
}
