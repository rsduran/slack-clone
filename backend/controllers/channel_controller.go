package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Channel
func CreateChannel(c *gin.Context) {
	var input struct {
		Name        string `json:"name"`
		WorkspaceID uint   `json:"workspace_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel := models.Channel{Name: input.Name, WorkspaceID: input.WorkspaceID}
	models.DB.Create(&channel) // Reference models.DB

	c.JSON(http.StatusOK, channel)
}

// Get Channels
func GetChannels(c *gin.Context) {
	var channels []models.Channel
	models.DB.Find(&channels) // Reference models.DB

	c.JSON(http.StatusOK, channels)
}
