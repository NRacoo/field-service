package dto

import (
	"field-service/constants"
	"time"

	"github.com/google/uuid"
)

type FieldScheduleRequest struct {
	FieldId string `json:"fieldID" validate:"required"`
	Date    string `json:"date" validate:"required"`
	TimeIds string `json:"timeIDs" validate:"required"`
}

type GenerateFieldScheduleForOneMonthRequest struct {
	FieldID string `json:"fieldID" validate:"required"`
}

type UpdateFieldScheduleRequest struct {
	Date   string `json:"date" validate:"required"`
	TimeID string `json:"timeIDs" validate:"required"`
}

type UpdateStatusFieldScheduleRequest struct {
	FieldScheduleIDS []string `json:"fieldScheduleIDs" validate:"required"`
}

type FieldScheduleRespons struct {
	UUID      uuid.UUID                         `json:"uuid"`
	FieldName string                            `json:"fieldName"`
	PriceHour string                            `json:"priceHour"`
	Date      string                            `json:"date"`
	Time      constants.FieldScheduleStatusName `json:"status"`
	CreatedAt *time.Time                        `json:"createdAt"`
	UpdatedAt *time.Time                        `json:"updatedAt"`
}

type FieldScheduleForBookingRespons struct {
	UUID      uuid.UUID                         `json:"uuid"`
	PriceHour string                            `json:"priceHour"`
	Date      string                            `json:"date"`
	Status    constants.FieldScheduleStatusName `json:"status"`
	Time      string                            `json:"time"`
}

type FieldScheduleRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}

type FieldScheduleByFieldIDAndDateRequestParam struct {
	Date string `form:"date" validate:"required"`
}
