package handlers

import (
	"net/http"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DashboardHandler struct {
	propertyUseCase *usecases.PropertyUseCase
	contractUseCase *usecases.ContractUseCase
	paymentUseCase  *usecases.PaymentUseCase
}

func NewDashboardHandler(
	propertyUseCase *usecases.PropertyUseCase,
	contractUseCase *usecases.ContractUseCase,
	paymentUseCase *usecases.PaymentUseCase,
) *DashboardHandler {
	return &DashboardHandler{
		propertyUseCase: propertyUseCase,
		contractUseCase: contractUseCase,
		paymentUseCase:  paymentUseCase,
	}
}

type DashboardResponse struct {
	TotalProperties     int                   `json:"total_properties"`
	RentedProperties    int                   `json:"rented_properties"`
	AvailableProperties int                   `json:"available_properties"`
	MonthlyRevenue      float64               `json:"monthly_revenue"`
	PendingPayments     int                   `json:"pending_payments"`
	OverduePayments     int                   `json:"overdue_payments"`
	RecentPayments      []RecentPayment       `json:"recent_payments"`
	MonthlyRevenues     []MonthlyRevenue      `json:"monthly_revenues"`
	PropertyStatus      []PropertyStatusCount `json:"property_status"`
}

type RecentPayment struct {
	ID       uuid.UUID `json:"id"`
	Tenant   string    `json:"tenant"`
	Property string    `json:"property"`
	Amount   float64   `json:"amount"`
	Date     string    `json:"date"`
}

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

type PropertyStatusCount struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	ownerIDStr := c.Param("ownerId")
	ownerID, err := uuid.Parse(ownerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	// Buscar propriedades do owner
	properties, err := h.propertyUseCase.GetPropertiesByOwner(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get properties"})
		return
	}

	totalProperties := len(properties)
	rentedProperties := 0
	availableProperties := 0

	// Contar propriedades por status
	for _, property := range properties {
		switch property.Status {
		case entities.PropertyStatusRented:
			rentedProperties++
		case entities.PropertyStatusAvailable:
			availableProperties++
		}
	}

	// Buscar contratos ativos para calcular receita
	var monthlyRevenue float64
	var propertyIDs []uuid.UUID
	for _, property := range properties {
		propertyIDs = append(propertyIDs, property.ID)
	}

	// Calcular receita mensal dos contratos ativos
	for _, propertyID := range propertyIDs {
		contract, err := h.contractUseCase.GetActiveContractByProperty(c.Request.Context(), propertyID)
		if err == nil && contract != nil {
			monthlyRevenue += contract.MonthlyRent
		}
	}

	// Buscar pagamentos pendentes e em atraso
	pendingPayments := 0
	overduePayments := 0

	// Buscar todos os pagamentos para calcular pendentes e em atraso
	allPayments, _, err := h.paymentUseCase.GetAllPayments(c.Request.Context(), 1, 1000)
	if err == nil {
		for _, payment := range allPayments {
			// Verificar se o pagamento pertence a uma propriedade do owner
			contract, err := h.contractUseCase.GetContract(c.Request.Context(), payment.ContractID)
			if err == nil {
				// Verificar se a propriedade pertence ao owner
				for _, propertyID := range propertyIDs {
					if contract.PropertyID == propertyID {
						if payment.Status == entities.PaymentStatusPending {
							pendingPayments++
							if payment.DueDate.Before(time.Now()) {
								overduePayments++
							}
						}
						break
					}
				}
			}
		}
	}

	// Buscar pagamentos recentes (últimos 5 pagos)
	recentPayments := []RecentPayment{}
	if err == nil {
		count := 0
		for _, payment := range allPayments {
			if count >= 5 {
				break
			}
			if payment.Status == entities.PaymentStatusPaid && payment.PaidDate != nil {
				// Verificar se pertence ao owner
				contract, err := h.contractUseCase.GetContract(c.Request.Context(), payment.ContractID)
				if err == nil {
					for _, propertyID := range propertyIDs {
						if contract.PropertyID == propertyID {
							property, err := h.propertyUseCase.GetProperty(c.Request.Context(), contract.PropertyID)
							if err == nil {
								recentPayment := RecentPayment{
									ID:       payment.ID,
									Tenant:   "Inquilino",
									Property: property.Title,
									Amount:   *payment.PaidAmount,
									Date:     payment.PaidDate.Format("2006-01-02"),
								}
								recentPayments = append(recentPayments, recentPayment)
								count++
							}
							break
						}
					}
				}
			}
		}
	}

	// Calcular receitas mensais dos últimos 4 meses
	now := time.Now()
	monthlyRevenues := []MonthlyRevenue{}
	for i := 3; i >= 0; i-- {
		month := now.AddDate(0, -i, 0)
		monthName := []string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"}[month.Month()-1]
		
		// Por enquanto usar receita atual, depois implementar cálculo real por período
		revenue := monthlyRevenue
		if i > 0 {
			revenue = monthlyRevenue * (1.0 - float64(i)*0.1) // Simulação de crescimento
		}
		
		monthlyRevenues = append(monthlyRevenues, MonthlyRevenue{
			Month:   monthName,
			Revenue: revenue,
		})
	}

	// Status das propriedades
	propertyStatus := []PropertyStatusCount{
		{Status: "Alugadas", Count: rentedProperties},
		{Status: "Disponíveis", Count: availableProperties},
	}

	response := DashboardResponse{
		TotalProperties:     totalProperties,
		RentedProperties:    rentedProperties,
		AvailableProperties: availableProperties,
		MonthlyRevenue:      monthlyRevenue,
		PendingPayments:     pendingPayments,
		OverduePayments:     overduePayments,
		RecentPayments:      recentPayments,
		MonthlyRevenues:     monthlyRevenues,
		PropertyStatus:      propertyStatus,
	}

	c.JSON(http.StatusOK, response)
}
