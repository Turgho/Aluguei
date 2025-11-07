package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OwnerHandler struct {
	ownerUseCase *usecases.OwnerUseCase
}

func NewOwnerHandler(ownerUseCase *usecases.OwnerUseCase) *OwnerHandler {
	return &OwnerHandler{
		ownerUseCase: ownerUseCase,
	}
}

type CreateOwnerRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Phone     string `json:"phone" binding:"required"`
	CPF       string `json:"cpf" binding:"required"`
	BirthDate string `json:"birth_date,omitempty"`
}

type UpdateOwnerRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date,omitempty"`
}

func (h *OwnerHandler) CreateOwner(c *gin.Context) {
	var req CreateOwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var birthDate *time.Time
	if req.BirthDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			birthDate = &parsed
		}
	}

	owner, err := h.ownerUseCase.CreateOwner(
		c.Request.Context(),
		req.Name,
		req.Email,
		req.Password,
		req.Phone,
		req.CPF,
		birthDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, owner)
}

func (h *OwnerHandler) GetOwner(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	owner, err := h.ownerUseCase.GetOwner(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "owner not found"})
		return
	}

	c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) GetOwnerByEmail(c *gin.Context) {
	email := c.Param("email")

	owner, err := h.ownerUseCase.GetOwnerByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "owner not found"})
		return
	}

	c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) GetAllOwners(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	owners, total, err := h.ownerUseCase.GetAllOwners(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(total) + limit - 1) / limit
	c.JSON(http.StatusOK, gin.H{
		"data": owners,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": totalPages,
		},
	})
}

func (h *OwnerHandler) UpdateOwner(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	var req UpdateOwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var birthDate *time.Time
	if req.BirthDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			birthDate = &parsed
		}
	}

	err = h.ownerUseCase.UpdateOwner(
		c.Request.Context(),
		id,
		req.Name,
		req.Email,
		req.Phone,
		req.CPF,
		birthDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	owner, _ := h.ownerUseCase.GetOwner(c.Request.Context(), id)
	c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) DeleteOwner(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	if err := h.ownerUseCase.DeleteOwner(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}