package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Field struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID      `gorm:"type:uuid;not null"`
	Code      string         `gorm:"type:varchar(15)";not null`
	Name      string         `gorm:"type:varchar(100);not null`
	PriceHour int            `gorm:"type:int;not null`
	Images    pq.StringArray `gorm:"type:text[];not null`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt
}
