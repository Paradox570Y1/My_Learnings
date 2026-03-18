package dto

import "time"

type CreateOrderRequest struct {
	ID           string `json:"id" binding:"required"`
	FacilityCode string `json:"facility_code" binding:"required"`
}

type OrderResponse struct {
	ID           string `json:"id" binding:"required"`
	FacilityCode string `json:"facility_code" binding:"required"`
	Status       string `json:"status" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
}