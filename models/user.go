package models

import (
	// "time"

	// "gorm.io/driver/sqlite"
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
	Experience           string
	AreasOfExpertise     []string `gorm:"type:text"`
	Certifications       []string `gorm:"type:text"`
	TwoFactorAuthEnabled bool
	ProjectName          string
	Role                 string
}

/*
type User struct {
	// User specific
	ID        int64     `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Avatar    string    `json:"avatar"`

	// Pentesting-specific
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	Organization         string   `json:"organization"`
	Experience           string   `json:"experience"`
	AreasOfExpertise     []string `json:"areasOfExpertise"`
	Certifications       []string `json:"certifications"`
	TwoFactorAuthEnabled bool     `json:"twoFactorAuthEnabled"`

	// Project specfic
	ProjectName string `json:"projectName"`

	// Role
	Role string `json:"role"`
}*/

const create string = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- SERIAL is auto-incrementing in PostgreSQL
    username VARCHAR(32) NOT NULL CHECK (length(username) >= 3), -- VARCHAR with length check
    email VARCHAR(255) NOT NULL CHECK (email LIKE '%_@__%.__%'), -- Basic email format check
    password VARCHAR(255) NOT NULL CHECK (length(password) >= 8), -- Password length check
    is_active BOOLEAN DEFAULT FALSE, -- BOOLEAN for boolean values
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- TIMESTAMP with default current timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- TIMESTAMP with default current timestamp
    avatar VARCHAR(255), -- VARCHAR for avatar URL
    first_name VARCHAR(255), -- VARCHAR for first name
    last_name VARCHAR(255), -- VARCHAR for last name
    organization VARCHAR(255), -- VARCHAR for organization
    experience TEXT, -- TEXT for longer text
    areas_of_expertise TEXT[], -- TEXT[] for array of strings
    certifications TEXT[], -- TEXT[] for array of strings
    two_factor_auth_enabled BOOLEAN DEFAULT FALSE, -- BOOLEAN for boolean values
    project_name VARCHAR(255), -- VARCHAR for project name
    role VARCHAR(255) -- VARCHAR for role
);`
