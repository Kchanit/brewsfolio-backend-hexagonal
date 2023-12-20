package router

import (
	"github.com/Kchanit/brewsfolio-backend/database"
	"github.com/Kchanit/brewsfolio-backend/handlers"
	"github.com/Kchanit/brewsfolio-backend/middlewares"
	"github.com/Kchanit/brewsfolio-backend/repositories"
	"github.com/Kchanit/brewsfolio-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.DBConn

	userRepo := repositories.NewUserRepository(db)
	beerRepo := repositories.NewBeerRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	collectionRepo := repositories.NewCollectionRepository(db)

	userService := services.NewUserService(userRepo)
	beerService := services.NewBeerService(beerRepo)
	reviewService := services.NewReviewService(reviewRepo)
	collectionService := services.NewCollectionService(collectionRepo, userRepo, beerRepo)

	userHandler := handlers.NewUserHandler(userService)
	beerHandler := handlers.NewBeerHandler(beerService)
	reviewHandler := handlers.NewReviewHandler(reviewService, beerService)
	collectionHandler := handlers.NewCollectionHandler(collectionService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yo, World 👋!")
	})

	api := app.Group("/api", middlewares.AuthMiddleware)
	api.Post("/register", userHandler.CreateUser)
	api.Post("/login", userHandler.Login)

	users := api.Group("/users")
	users.Get("/:id", userHandler.GetUserByID)
	users.Get("", userHandler.GetUsers)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
	users.Get("/:id/favorites", beerHandler.GetFavoritesByUserID)
	users.Get("/:id/collections", collectionHandler.GetCollectionsByUserID)

	beers := api.Group("/beers")
	beers.Get("", beerHandler.GetBeers)
	beers.Get("/:id", beerHandler.GetBeerByID)
	beers.Post("", middlewares.AdminMiddleware, beerHandler.CreateBeer)
	beers.Put("/:id", middlewares.AdminMiddleware, beerHandler.UpdateBeer)
	beers.Delete("/:id", middlewares.AdminMiddleware, beerHandler.DeleteBeer)

	beers.Post("/:id/favorite", beerHandler.AddFavoriteBeer)
	beers.Delete("/:id/favorite", beerHandler.RemoveFavoriteBeer)
	beers.Post("/:id/reviews", reviewHandler.CreateReview)
	beers.Get("/:id/reviews", reviewHandler.GetReviewsByBeerID)

	reviews := api.Group("/reviews")
	reviews.Get("", reviewHandler.GetReviews)
	reviews.Delete("/:id", middlewares.AuthMiddleware, reviewHandler.DeleteReview)
	reviews.Get("/:id", reviewHandler.GetReviewByID)

	collections := api.Group("/collections")
	collections.Post("", collectionHandler.CreateCollection)
	collections.Delete("/:id", collectionHandler.DeleteCollection)
	collections.Post("/:id/beers/:beerId", collectionHandler.AddBeerToCollection)
	collections.Put("/:id/beers", collectionHandler.AddBeersToCollection)
	collections.Get("/:id", collectionHandler.GetCollectionByID)
	collections.Get("", collectionHandler.GetCollections)

	api.Get("/public/collections", collectionHandler.GetPublicCollections)
}
