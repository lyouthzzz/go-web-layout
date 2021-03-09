package gormx

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli;not null"`
}
