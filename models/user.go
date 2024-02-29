package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username             string `gorm:"unique;not null;size:32"`
	Email                string `gorm:"unique;not null"`
	Password             string `gorm:"not null"`
	IsActive             bool
	Avatar               string
	FirstName            string
	LastName             string
	Organization         string
	Experience           int32
	AreasOfExpertise     string
	Certifications       string
	TwoFactorAuthEnabled bool
	ProjectName          string
	Role                 string
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
