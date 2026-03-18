package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"MyProject/internal/services"
	"MyProject/internal/dto"
)

type FacilityHandler struct {
	facilityService services.FacilityService
}

func NewFacilityHandler(service services.FacilityService) *FacilityHandler {
	return &FacilityHandler{
		facilityService: service,
	}
}

func (h *FacilityHandler) GetAll(c *gin.Context) {
	facilities, err := h.facilityService.GetAll(c.Request.Context())
	if err != nil {
		if services.IsOrderNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})	
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, facilities)
}

func (h *FacilityHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")
	facility, err := h.facilityService.GetByCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, facility)
}

func (h *FacilityHandler) Create(c *gin.Context) {
	var facility dto.CreateFacilityRequest
	err := c.ShouldBindJSON(&facility)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = h.facilityService.Create(c.Request.Context(), facility)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "facility created successfully"})
}