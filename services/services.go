package services

import (
	"errors"
	"fmt"

	"um/models"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUserByID(int64) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int64) error
}

type gormUserService struct {
	db *gorm.DB
}

func NewGormUserService(db *gorm.DB) UserService {
	return &gormUserService{db: db}
}

func (r *gormUserService) CreateUser(user *models.User) error {
	fmt.Println("Creating user for signup...")
	return r.db.Create(&user).Error
}

func (r *gormUserService) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *gormUserService) UpdateUser(user *models.User) error {
	if user.ID == 0 {
		return errors.New("user ID cannot be zero for update")
	}
	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *gormUserService) DeleteUser(id int64) error {
	return r.db.Delete(&models.User{}, id).Error
}
