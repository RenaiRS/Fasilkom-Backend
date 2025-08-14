package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string     `json:"username" gorm:"not null;unique"`
	Password string     `json:"password" gorm:"not null"`
	Token    string     `json:"token"`
	TokenExp *time.Time `json:"token_exp"` // pointer supaya nullable
}
