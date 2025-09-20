package models

import (
	"field-service/constants"
	"time"

	"github.com/google/uuid"
)

type FieldSchedules struct {
	ID        uint      `gorm:primaryKey;autoIncrement`
	UUID      uuid.UUID `gorm:type:uuid;not null`
	FieldID   uint      `gorm:type:int;not null`
	TimeId    uint      `gorm:type:int;not null`
	Date      time.Time `gorm:type:int;not null`
	Status    constants.FieldScheduleStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Field     Field `gorm:"foreignKey:field_id;references:id;constraint:OnUpdate:CASCADE;OnDelete:CASCADE`
	Time      Time  `gorm:"foreignKey:time_id;references:id;constraint:OnUpdate:CASCADE;OnDelete:CASCADE`
}
