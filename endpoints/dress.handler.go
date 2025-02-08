package endpoints

import (
	"core-service/models"
	"core-service/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DressRequest struct {
	Url         string
	Description string                 `json:"description"`
	Labels      []string               `json:"labels"`
	Weather     float64                `json:"weather"`
	Location    map[string]interface{} `json:"location" gorm:"type:jsonb"`
	id          uint
}

type DressHandler struct {
	DressService *service.DressService
}

func NewDressHandler(service *service.DressService) *DressHandler {
	return &DressHandler{DressService: service}
}

func (h *DressHandler) CreateDressHandler(c *gin.Context) {
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

	err := h.DressService.AddDress(dress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (h *DressHandler) GetAllDresses(c *gin.Context) {
	dresses, err := h.DressService.FindDress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch dresses"})
		return
	}
	c.JSON(http.StatusOK, dresses)
}

func (h *DressHandler) UpdatePartialDressById(c *gin.Context) {
	var req map[string]interface{}
	id, _ := strconv.Atoi(c.Param("id"))
	var entity = &models.Dress{}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DressService.UpdatesPartial(entity, uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entity updated successfully"})
}

func (h *DressHandler) DeletePartialDressById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var entity = &models.Dress{}

	if err := h.DressService.DeletePartialDress(entity, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entity Deleted successfully"})
}
