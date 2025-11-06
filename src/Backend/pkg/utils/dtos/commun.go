package dtos

import (
	"time"
)

type IDResponse struct {
	ID string `json:"id"`
}

type TimestampResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaginationRequest struct {
	Page     int `form:"page" validate:"min=1" default:"1"`
	PageSize int `form:"page_size" validate:"min=1,max=100" default:"20"`
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}
