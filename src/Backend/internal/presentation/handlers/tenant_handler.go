package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TenantHandler struct {
	tenantUseCase *usecases.TenantUseCase
}

func NewTenantHandler(tenantUseCase *usecases.TenantUseCase) *TenantHandler {
	return &TenantHandler{
		tenantUseCase: tenantUseCase,
	}
}

type CreateTenantRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	CPF       string `json:"cpf" binding:"required"`
	OwnerID   string `json:"owner_id" binding:"required"`
	BirthDate string `json:"birth_date,omitempty"`
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var req CreateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, err := uuid.Parse(req.OwnerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner_id"})
		return
	}

	var birthDate *time.Time
	if req.BirthDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			birthDate = &parsed
		}
	}

	tenant, err := h.tenantUseCase.CreateTenant(
		c.Request.Context(),
		req.Name,
		req.Email,
		req.Phone,
		req.CPF,
		ownerID,
		birthDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tenant)
}

func (h *TenantHandler) GetTenant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	tenant, err := h.tenantUseCase.GetTenant(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tenant not found"})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

func (h *TenantHandler) GetTenantsByOwner(c *gin.Context) {
	ownerID, err := uuid.Parse(c.Param("ownerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	tenants, err := h.tenantUseCase.GetTenantsByOwner(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenants)
}

func (h *TenantHandler) GetAllTenants(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	tenants, total, err := h.tenantUseCase.GetAllTenants(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(total) + limit - 1) / limit
	c.JSON(http.StatusOK, gin.H{
		"data": tenants,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": totalPages,
		},
	})
}

func (h *TenantHandler) DeleteTenant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	if err := h.tenantUseCase.DeleteTenant(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}