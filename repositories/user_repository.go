package repositories

import (
	"github.com/Kchanit/brewsfolio-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	err := r.db.Save(&user).Error
	return err
}

func (r *userRepository) DeleteUser(id string) error {
	err := r.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}


