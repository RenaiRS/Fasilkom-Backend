package models

import (
	"time"

	"gorm.io/gorm"
)

type KemahasiswaanKerjasama struct {
    gorm.Model
    Judul     string    `json:"judul" gorm:"not null"`
    Tanggal   time.Time `json:"tanggal"`
    Cover     string    `json:"cover"`
    Content   string    `json:"content"`
    PostedBy  string    `json:"posted_by"`
}
