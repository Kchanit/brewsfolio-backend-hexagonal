package handlers

import (
	"github.com/Kchanit/brewsfolio-backend/models"
	"github.com/Kchanit/brewsfolio-backend/services"
	"github.com/Kchanit/brewsfolio-backend/utility"
	"github.com/gofiber/fiber/v2"
)

type BeerHandler struct {
	beerService *services.BeerService
}

func NewBeerHandler(beerService *services.BeerService) *BeerHandler {
	return &BeerHandler{beerService: beerService}
}

func (h *BeerHandler) CreateBeer(c *fiber.Ctx) error {
	beer := new(models.Beer)
	if err := c.BodyParser(beer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	beer, err := h.beerService.CreateBeer(beer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(beer)
}

func (h *BeerHandler) GetBeerByID(c *fiber.Ctx) error {
	beerID := c.Params("id")

	beer, err := h.beerService.GetBeerByID(beerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Beer not found"})
	}

	return c.JSON(beer)
}

func (h *BeerHandler) GetBeers(c *fiber.Ctx) error {
	beers, err := h.beerService.GetBeers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.JSON(beers)
}

func (h *BeerHandler) UpdateBeer(c *fiber.Ctx) error {
	beerID := c.Params("id")

	beer := new(models.Beer)
	if err := c.BodyParser(beer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	beer, err := h.beerService.UpdateBeer(beerID, beer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(beer)
}

func (h *BeerHandler) DeleteBeer(c *fiber.Ctx) error {
	beerID := c.Params("id")

	err := h.beerService.DeleteBeer(beerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Beer successfully deleted"})
}

func (h *BeerHandler) GetFavoritesByUserID(c *fiber.Ctx) error {
	userID := c.Params("id")
	favorites, err := h.beerService.GetFavoritesByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(favorites)
}

func (h *BeerHandler) AddFavoriteBeer(c *fiber.Ctx) error {
	beerID := c.Params("id")

	uid, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	beer, err := h.beerService.AddFavoriteBeer(uid, beerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(beer)
}

func (h *BeerHandler) RemoveFavoriteBeer(c *fiber.Ctx) error {

	beerID := c.Params("id")

	uid, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	beer, err := h.beerService.RemoveFavoriteBeer(uid, beerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(beer)
}
