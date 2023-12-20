package repositories

import (
	"fmt"

	"github.com/Kchanit/brewsfolio-backend/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewsByBeerID(id string) ([]models.Review, error)
	GetReviews() ([]models.Review, error)
	GetReviewByID(id string) (*models.Review, error)
	CreateReview(review *models.Review) (*models.Review, error)
	DeleteReview(id string) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetReviewsByBeerID(id string) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("beer_id = ?", id).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviewByID(id string) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("id = ?", id).First(&review).Error
	return &review, err
}

func (r *reviewRepository) CreateReview(review *models.Review) (*models.Review, error) {
	fmt.Println("review repositories")
	err := r.db.Create(review).Error
	return review, err
}

func (r *reviewRepository) DeleteReview(id string) error {
	err := r.db.Where("id = ?", id).Delete(&models.Review{}).Error
	return err
}
