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

type PaymentHandler struct {
	paymentUseCase *usecases.PaymentUseCase
}

func NewPaymentHandler(paymentUseCase *usecases.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

type CreatePaymentRequest struct {
	ContractID    uuid.UUID `json:"contract_id" binding:"required"`
	DueDate       string    `json:"due_date" binding:"required"`
	Amount        float64   `json:"amount" binding:"required,gt=0"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	Notes         string    `json:"notes,omitempty"`
}

type UpdatePaymentRequest struct {
	PaidDate      *string `json:"paid_date,omitempty"`
	AmountPaid    *float64 `json:"amount_paid,omitempty"`
	Status        *string `json:"status,omitempty"`
	PaymentMethod *string `json:"payment_method,omitempty"`
	Notes         *string `json:"notes,omitempty"`
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse due date
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}

	payment, err := h.paymentUseCase.CreatePayment(c.Request.Context(), req.ContractID, req.Amount, dueDate, entities.PaymentStatusPending)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPayments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := entities.PaymentFilters{
		Page:  page,
		Limit: limit,
	}

	if status := c.Query("status"); status != "" {
		filters.Status = &status
	}

	if contractID := c.Query("contract_id"); contractID != "" {
		if id, err := uuid.Parse(contractID); err == nil {
			filters.ContractID = &id
		}
	}

	payments, total, err := h.paymentUseCase.GetAllPayments(c.Request.Context(), filters.Page, filters.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(total) + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
		"pagination": gin.H{
			"page":   page,
			"limit":  limit,
			"total":  total,
			"pages":  totalPages,
		},
	})
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := h.paymentUseCase.GetPayment(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var req UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.paymentUseCase.UpdatePayment(c.Request.Context(), id, 0, time.Time{}, nil, nil, entities.PaymentStatusPending)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	err = h.paymentUseCase.DeletePayment(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *PaymentHandler) GetPaymentsByContract(c *gin.Context) {
	contractIDStr := c.Param("contractId")
	contractID, err := uuid.Parse(contractIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contract ID"})
		return
	}

	payments, err := h.paymentUseCase.GetPaymentsByContract(c.Request.Context(), contractID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetOverduePayments(c *gin.Context) {
	payments, err := h.paymentUseCase.GetOverduePayments(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetPaymentsByPeriod(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}

	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	payments, err := h.paymentUseCase.GetPaymentsByPeriod(c.Request.Context(), startTime, endTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}