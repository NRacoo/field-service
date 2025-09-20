package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type FieldRequest struct {
	Name      string                 `json:"name" validate:"required"`
	Code      string                 `json:"code" validate:"required"`
	PriceHour int                    `json:"priceHour" validate:"required"`
	Images    []multipart.FileHeader `json:"images" validate:"required"`
}

type UpdateFieldRequest struct {
	Name      string                 `json:"name" validate:"required"`
	Code      string                 `json:"code" validate:"required"`
	PriceHour int                    `json:"priceHour" validate:"required"`
	Images    []multipart.FileHeader `json:"images"`
}

type FieldRespons struct {
	UUID      uuid.UUID  `json:"uiid"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceHour int        `json:"priceHour"`
	Images    []string   `json:"images"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type FieldDetailRespons struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceHour int        `json:"priceHour"`
	Images    []string   `json:"images"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type FieldRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"page" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}
