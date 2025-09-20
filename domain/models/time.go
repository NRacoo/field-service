package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Time struct {
	ID             uint      `gorm:primaryKey;autoIncrement`
	UUID           uuid.UUID `gorm:"type:uuid;not null"`
	StartTime      string    `gorm:"type:time without timezone;not null"`
	EndTime        string    `gorm:"type:time without timezone;not null"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *gorm.DeletedAt
	FieldSchedules []FieldSchedules `gorm:"foreignKey:field_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
