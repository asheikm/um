package models

import "time"

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
}
