package handlers

import (
	loadenv "api-bff-go-gin-todo/load_env"
	"api-bff-go-gin-todo/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PatchTaskEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.TaskResponse

	jsonData, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, loadenv.LoadEnvFile().ApiBackPatchEdit+id+"/", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error in request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var taskResonseBody models.Task
	if err := json.NewDecoder(resp.Body).Decode(&taskResonseBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task edited successfully",
		"data":    taskResonseBody,
	})
}
