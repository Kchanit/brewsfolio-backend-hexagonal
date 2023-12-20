package services

import (
	"errors"
	"os"
	"time"

	"github.com/Kchanit/brewsfolio-backend/models"
	"github.com/Kchanit/brewsfolio-backend/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id string) (models.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUsers retrieves all users.
func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user *models.User) (*models.User, string, error) {

	_, err := s.userRepository.GetUserByEmail(user.Email)
	if err == nil {
		return nil, "", errors.New("Email already exists")
	}

	hash := HashPassword(user.Password)

	user.Password = hash
	user.Role = "USER"

	user, err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, "", err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return user, tokenString, nil
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	// Find the user in the database
	existingUser, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Update the user information
	existingUser.Name = user.Name
	existingUser.Email = user.Email

	// Save the changes back to the database
	err = s.userRepository.UpdateUser(&existingUser)
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

// DeleteUser deletes a user.
func (s *UserService) DeleteUser(id string) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

// LoginUser logs in a user.
func (s *UserService) Login(email string, password string) (*models.User, string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("Incorrect password")
	}

	tokenString, err := GenerateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return &user, tokenString, nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
