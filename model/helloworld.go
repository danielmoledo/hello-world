package model

import (
	"time"
)

type HelloWorld struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Timestamp time.Time `json:"timestamp"`
}
