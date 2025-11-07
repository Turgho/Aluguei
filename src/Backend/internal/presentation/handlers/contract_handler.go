package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
)

type ContractHandler struct {
	contractUseCase *usecases.ContractUseCase
}

func NewContractHandler(contractUseCase *usecases.ContractUseCase) *ContractHandler {
	return &ContractHandler{
		contractUseCase: contractUseCase,
	}
}

type CreateContractRequest struct {
	PropertyID     uuid.UUID `json:"property_id" binding:"required"`
	TenantID       uuid.UUID `json:"tenant_id" binding:"required"`
	StartDate      string    `json:"start_date" binding:"required"`
	EndDate        *string   `json:"end_date,omitempty"`
	MonthlyRent    float64   `json:"monthly_rent" binding:"required,gt=0"`
	PaymentDueDay  int       `json:"payment_due_day" binding:"required,min=1,max=28"`
	Status         string    `json:"status,omitempty"`
}

type UpdateContractRequest struct {
	StartDate     *string  `json:"start_date,omitempty"`
	EndDate       *string  `json:"end_date,omitempty"`
	MonthlyRent   *float64 `json:"monthly_rent,omitempty"`
	PaymentDueDay *int     `json:"payment_due_day,omitempty"`
	Status        *string  `json:"status,omitempty"`
}

func (h *ContractHandler) CreateContract(c *gin.Context) {
	var req CreateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	var endDate time.Time
	if req.EndDate != nil {
		endDate, err = time.Parse("2006-01-02", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}
	}

	status := entities.ContractStatusActive
	if req.Status != "" {
		status = entities.ContractStatus(req.Status)
	}

	contract, err := h.contractUseCase.CreateContract(c.Request.Context(), req.PropertyID, req.TenantID, startDate, endDate, req.MonthlyRent, req.PaymentDueDay, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contract)
}

func (h *ContractHandler) GetContracts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := entities.ContractFilters{
		Page:  page,
		Limit: limit,
	}

	if status := c.Query("status"); status != "" {
		filters.Status = &status
	}

	if propertyID := c.Query("property_id"); propertyID != "" {
		if id, err := uuid.Parse(propertyID); err == nil {
			filters.PropertyID = &id
		}
	}

	if tenantID := c.Query("tenant_id"); tenantID != "" {
		if id, err := uuid.Parse(tenantID); err == nil {
			filters.TenantID = &id
		}
	}

	contracts, total, err := h.contractUseCase.GetAllContracts(c.Request.Context(), filters.Page, filters.Limit, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(total) + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data": contracts,
		"pagination": gin.H{
			"page":   page,
			"limit":  limit,
			"total":  total,
			"pages":  totalPages,
		},
	})
}

func (h *ContractHandler) GetContractByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contract ID"})
		return
	}

	contract, err := h.contractUseCase.GetContract(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (h *ContractHandler) UpdateContract(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contract ID"})
		return
	}

	var req UpdateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.contractUseCase.UpdateContract(c.Request.Context(), id, time.Time{}, time.Time{}, 0, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contract updated successfully"})
}

func (h *ContractHandler) DeleteContract(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contract ID"})
		return
	}

	err = h.contractUseCase.DeleteContract(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ContractHandler) GetContractsByProperty(c *gin.Context) {
	propertyIDStr := c.Param("propertyId")
	propertyID, err := uuid.Parse(propertyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	contracts, err := h.contractUseCase.GetContractsByProperty(c.Request.Context(), propertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (h *ContractHandler) GetContractsByTenant(c *gin.Context) {
	tenantIDStr := c.Param("tenantId")
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	contracts, err := h.contractUseCase.GetContractsByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (h *ContractHandler) GetActiveContractByProperty(c *gin.Context) {
	propertyIDStr := c.Param("propertyId")
	propertyID, err := uuid.Parse(propertyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	contract, err := h.contractUseCase.GetActiveContractByProperty(c.Request.Context(), propertyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active contract found for this property"})
		return
	}

	c.JSON(http.StatusOK, contract)
}