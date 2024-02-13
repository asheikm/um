package repositories

import (
	"errors"

	"um/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(*models.User) error
	GetUserByID(int64) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int64) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *gormUserRepository) UpdateUser(user *models.User) error {
	if user.ID == 0 {
		return errors.New("user ID cannot be zero for update")
	}
	return r.db.Model(&models.User{ID: user.ID}).Updates(user).Error
}

func (r *gormUserRepository) DeleteUser(id int64) error {
	return r.db.Delete(&models.User{ID: id}).Error
}
