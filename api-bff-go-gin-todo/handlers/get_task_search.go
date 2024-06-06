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

func GetTaskSearch(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := http.Get(loadenv.LoadEnvFile().ApiBackGetSearch + id + "/")
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

	var task models.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, task)
}
