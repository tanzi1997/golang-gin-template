package model

import (
	"time"

	"gorm.io/gorm"
)

type PrimarykeyIdModel struct {
	Id int `gorm:"primarykey"`
}

type IdModel struct {
	Id int
}

type TimeMode struct {
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp"`
}
