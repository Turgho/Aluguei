package entities

import "github.com/google/uuid"

type SearchFilters struct {
	MinRooms int            `json:"min_rooms,omitempty"`
	MaxRooms int            `json:"max_rooms,omitempty"`
	MinArea  float64        `json:"min_area,omitempty"`
	MaxArea  float64        `json:"max_area,omitempty"`
	MinRent  float64        `json:"min_rent,omitempty"`
	MaxRent  float64        `json:"max_rent,omitempty"`
	Status   PropertyStatus `json:"status,omitempty"`
	City     string         `json:"city,omitempty"`
	State    string         `json:"state,omitempty"`
}

type OwnerFilters struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type TenantFilters struct {
	Page    int        `json:"page"`
	Limit   int        `json:"limit"`
	OwnerID *uuid.UUID `json:"owner_id,omitempty"`
}

type PropertyFilters struct {
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
	Status  *string         `json:"status,omitempty"`
	OwnerID *uuid.UUID      `json:"owner_id,omitempty"`
}

type ContractFilters struct {
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Status     *string    `json:"status,omitempty"`
	PropertyID *uuid.UUID `json:"property_id,omitempty"`
	TenantID   *uuid.UUID `json:"tenant_id,omitempty"`
}

type PaymentFilters struct {
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Status     *string    `json:"status,omitempty"`
	ContractID *uuid.UUID `json:"contract_id,omitempty"`
}

// HasFilters checks if any filters are applied
func (sf *SearchFilters) HasFilters() bool {
	return sf.MinRooms > 0 ||
		sf.MaxRooms > 0 ||
		sf.MinArea > 0 ||
		sf.MaxArea > 0 ||
		sf.MinRent > 0 ||
		sf.MaxRent > 0 ||
		sf.Status != "" ||
		sf.City != "" ||
		sf.State != ""
}