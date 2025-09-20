package dto

import "github.com/google/uuid"

type TimeRequest struct {
	StartTime string `json:"startTime" validate:"required"`
	EndTime   string `json:"endTime" validate:"required"`
}

type TimeRespons struct {
	UUID      uuid.UUID `json:"uuid"`
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	CreatedAt *string   `json:"createdAt"`
	UpdateAt  *string   `json:"updatedAt"`
}
