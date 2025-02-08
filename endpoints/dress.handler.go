package endpoints

import (
	"core-service/models"
	"core-service/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DressRequest struct {
	Url         string                 `json:"url"`
	Description string                 `json:"description"`
	Labels      []string               `json:"labels"`
	Weather     float64                `json:"weather"`
	Location    map[string]interface{} `json:"location" gorm:"type:jsonb"`
}

func CreateDressHandler(c *gin.Context) {
	dressService := service.DressServiceInstance()
	var req DressRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	locationJSON, _ := json.Marshal(req.Location)
	dress := models.Dress{
		Url:         req.Url,
		Description: req.Description,
		Labels:      req.Labels,
		Weather:     req.Weather,
		Location:    locationJSON,
	}

	err := dressService.AddDress(dress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
