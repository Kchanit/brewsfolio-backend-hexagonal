package services

import (
	"fmt"
	"log"

	"github.com/Kchanit/brewsfolio-backend/models"
	"github.com/Kchanit/brewsfolio-backend/repositories"
)

type ReviewService struct {
	reviewRepository repositories.ReviewRepository
}

func NewReviewService(reviewRepository repositories.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepository: reviewRepository}
}

// GetReviews retrieves all reviews.
func (s *ReviewService) GetReviews() ([]models.Review, error) {
	reviews, err := s.reviewRepository.GetReviews()
	if err != nil {
		log.Println(err)
		return []models.Review{}, err
	}
	return reviews, nil
}

// GetReviewsByBeerId retrieves reviews by beer ID.
func (s *ReviewService) GetReviewsByBeerID(beerID string) ([]models.Review, error) {
	reviews, err := s.reviewRepository.GetReviewsByBeerID(beerID)
	if err != nil {
		log.Println(err)
		return []models.Review{}, err
	}
	return reviews, nil
}

// GetReviewByID retrieves a review by ID.
func (s *ReviewService) GetReviewByID(id string) (*models.Review, error) {
	review, err := s.reviewRepository.GetReviewByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return review, nil
}

// CreateReview creates a new review.
func (s *ReviewService) CreateReview(userID uint, beer *models.Beer, rating uint, description string) (*models.Review, error) {
	fmt.Println("Review service")
	newReview := &models.Review{
		UserID:      userID,
		BeerID:      beer.ID,
		Rating:      rating,
		Description: description,
	}

	createdReview, err := s.reviewRepository.CreateReview(newReview)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return createdReview, nil
}

// DeleteReview deletes a review.
func (s *ReviewService) DeleteReview(id string) error {
	err := s.reviewRepository.DeleteReview(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
