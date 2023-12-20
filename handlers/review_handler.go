package handlers

import (
	"fmt"

	"github.com/Kchanit/brewsfolio-backend/services"
	"github.com/Kchanit/brewsfolio-backend/utility"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewService *services.ReviewService
	beerService   *services.BeerService
}

func NewReviewHandler(reviewService *services.ReviewService, beerService *services.BeerService) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService, beerService: beerService}
}

func (h *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	reviews, err := h.reviewService.GetReviews()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reviews)
}

func (h *ReviewHandler) GetReviewsByBeerID(c *fiber.Ctx) error {
	beerID := c.Params("id")
	reviews, err := h.reviewService.GetReviewsByBeerID(beerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reviews)
}

func (h *ReviewHandler) GetReviewByID(c *fiber.Ctx) error {
	beerID := c.Params("id")
	reviews, err := h.reviewService.GetReviewByID(beerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reviews)
}

func (h *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	fmt.Println("CreateReview")
	payload := new(struct {
		Rating      uint   `json:"rating"`
		Description string `json:"description"`
	})

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	beerID := c.Params("id")
	// Check if the beer with the provided beerID exists
	beer, err := h.beerService.GetBeerByID(beerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Beer not found"})
	}

	// Get userID from the token
	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Rename the review variable to avoid reusing the same name
	newReview, err := h.reviewService.CreateReview(userID, &beer, payload.Rating, payload.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(newReview)
}

func (h *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	beerID := c.Params("id")
	err := h.reviewService.DeleteReview(beerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON("Review successfully deleted")
}
