package handlers

import (
	"github.com/Kchanit/brewsfolio-backend/services"
	"github.com/Kchanit/brewsfolio-backend/utility"
	"github.com/gofiber/fiber/v2"
)

type CollectionHandler struct {
	collectionService *services.CollectionService
}

func NewCollectionHandler(collectionService *services.CollectionService) *CollectionHandler {
	return &CollectionHandler{collectionService: collectionService}
}

func (h *CollectionHandler) GetCollections(c *fiber.Ctx) error {

	collections, err := h.collectionService.GetCollections()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(collections)
}

func (h *CollectionHandler) GetCollectionByID(c *fiber.Ctx) error {
	collectionID := c.Params("id")
	collection, err := h.collectionService.GetCollectionByID(collectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(collection)
}

func (h *CollectionHandler) GetCollectionsByUserID(c *fiber.Ctx) error {
	userID := c.Params("id")
	collections, err := h.collectionService.GetCollectionsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(collections)
}

func (h *CollectionHandler) GetPublicCollections(c *fiber.Ctx) error {
	collections, err := h.collectionService.GetPublicCollections()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(collections)
}

func (h *CollectionHandler) CreateCollection(c *fiber.Ctx) error {
	payload := new(struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		IsPrivate   bool   `json:"is_private"`
	})

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	collection, err := h.collectionService.CreateCollection(userID, payload.Title, payload.Description, payload.IsPrivate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(collection)
}

func (h *CollectionHandler) DeleteCollection(c *fiber.Ctx) error {
	collectionID := c.Params("id")

	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.collectionService.DeleteCollection(userID, collectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Collection deleted successfully"})
}

func (h *CollectionHandler) AddBeerToCollection(c *fiber.Ctx) error {
	collectionID := c.Params("id")
	beerID := c.Params("beerId")

	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	collection, err := h.collectionService.AddBeerToCollection(userID, collectionID, beerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Beer added to collection successfully", "collection": collection})
}

func (h *CollectionHandler) AddBeersToCollection(c *fiber.Ctx) error {
	collectionID := c.Params("id")
	beerIDs := new(struct {
		BeerIDs []uint `json:"beer_ids"`
	})

	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	collection, err := h.collectionService.GetCollectionByID(collectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if collection.UserID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if err := c.BodyParser(beerIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	collection, err = h.collectionService.AddBeersToCollection(userID, collectionID, beerIDs.BeerIDs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Beers added to collection successfully", "collection": collection})
}

func (h *CollectionHandler) RemoveBeerFromCollection(c *fiber.Ctx) error {
	collectionID := c.Params("id")
	beerID := c.Params("beerId")

	userID, err := utility.UserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	collection, err := h.collectionService.RemoveBeerFromCollection(userID, collectionID, beerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Beer removed from collection successfully", "collection": collection})
}
