package models

import (
	"time"

	"gorm.io/gorm"
)

type Berita struct {
    gorm.Model
    Judul      string    `json:"judul" gorm:"not null"`
    Tanggal    time.Time `json:"tanggal"`
    Cover      string    `json:"cover"`
    Content    string    `json:"content"`
    IsPriority bool      `json:"is_priority"`
    PostedBy   string    `json:"posted_by"`
}