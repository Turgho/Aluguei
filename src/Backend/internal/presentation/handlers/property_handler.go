package handlers

import (
	"net/http"
	"strconv"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PropertyHandler struct {
	propertyUseCase *usecases.PropertyUseCase
}

func NewPropertyHandler(propertyUseCase *usecases.PropertyUseCase) *PropertyHandler {
	return &PropertyHandler{
		propertyUseCase: propertyUseCase,
	}
}

type CreatePropertyRequest struct {
	OwnerID     string                  `json:"owner_id" binding:"required"`
	Title       string                  `json:"title" binding:"required"`
	Description string                  `json:"description"`
	Address     string                  `json:"address" binding:"required"`
	City        string                  `json:"city" binding:"required"`
	State       string                  `json:"state" binding:"required"`
	ZipCode     string                  `json:"zip_code"`
	Bedrooms    int                     `json:"bedrooms"`
	Bathrooms   int                     `json:"bathrooms"`
	Area        int                     `json:"area"`
	RentAmount  float64                 `json:"rent_amount" binding:"required,min=0"`
	Status      entities.PropertyStatus `json:"status"`
}

type UpdatePropertyRequest struct {
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Address     string                  `json:"address"`
	City        string                  `json:"city"`
	State       string                  `json:"state"`
	ZipCode     string                  `json:"zip_code"`
	Bedrooms    int                     `json:"bedrooms"`
	Bathrooms   int                     `json:"bathrooms"`
	Area        int                     `json:"area"`
	RentAmount  float64                 `json:"rent_amount"`
	Status      entities.PropertyStatus `json:"status"`
}

func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	var req CreatePropertyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, err := uuid.Parse(req.OwnerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner_id"})
		return
	}

	property, err := h.propertyUseCase.CreateProperty(
		c.Request.Context(),
		ownerID,
		req.Title,
		req.Description,
		req.Address,
		req.City,
		req.State,
		req.ZipCode,
		req.Bedrooms,
		req.Bathrooms,
		req.Area,
		req.RentAmount,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, property)
}

func (h *PropertyHandler) GetProperty(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
		return
	}

	property, err := h.propertyUseCase.GetProperty(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "property not found"})
		return
	}

	c.JSON(http.StatusOK, property)
}

func (h *PropertyHandler) GetPropertiesByOwner(c *gin.Context) {
	ownerID, err := uuid.Parse(c.Param("ownerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	properties, err := h.propertyUseCase.GetPropertiesByOwner(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}

func (h *PropertyHandler) GetAllProperties(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	status := c.Query("status")

	properties, total, err := h.propertyUseCase.GetAllProperties(c.Request.Context(), page, limit, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(total) + limit - 1) / limit
	c.JSON(http.StatusOK, gin.H{
		"data": properties,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": totalPages,
		},
	})
}

func (h *PropertyHandler) UpdateProperty(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
		return
	}

	var req UpdatePropertyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.propertyUseCase.UpdateProperty(
		c.Request.Context(),
		id,
		req.Title,
		req.Description,
		req.Address,
		req.City,
		req.State,
		req.ZipCode,
		req.Bedrooms,
		req.Bathrooms,
		req.Area,
		req.RentAmount,
		req.Status,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	property, _ := h.propertyUseCase.GetProperty(c.Request.Context(), id)
	c.JSON(http.StatusOK, property)
}

func (h *PropertyHandler) DeleteProperty(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
		return
	}

	if err := h.propertyUseCase.DeleteProperty(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}