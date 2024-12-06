package routes

import (
	"codex-backend/controllers"
	"codex-backend/db"
	"codex-backend/middleware"
	"codex-backend/models"
	"codex-backend/repositories"
	"codex-backend/services"
	"log"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes registers API routes
func RegisterRoutes(app *fiber.App) {
	log.Println("🏗️  Initializing API dependencies and routes...")
	
	// Initialize dependencies
	platformRepo := repositories.NewMongoPlatformRepository(db.GetCollection("platforms"))
	platformService := services.NewPlatformService(platformRepo)
	platformController := controllers.NewPlatformController(platformService)

	gameRepo := repositories.NewMongoGameRepository(db.GetCollection("games"))
	gameService := services.NewGameService(gameRepo)
	gameController := controllers.NewGameController(gameService)

	genreRepo := repositories.NewMongoGenreRepository(db.GetCollection("genres"))
	genreService := services.NewGenreService(genreRepo)
	genreController := controllers.NewGenreController(genreService)

	developerRepo := repositories.NewMongoDeveloperRepository(db.GetCollection("developers"))
	developerService := services.NewDeveloperService(developerRepo)
	developerController := controllers.NewDeveloperController(developerService)

	publisherRepo := repositories.NewMongoPublisherRepository(db.GetCollection("publishers"))
	publisherService := services.NewPublisherService(publisherRepo)
	publisherController := controllers.NewPublisherController(publisherService)

	api := app.Group("/api")

	// Apply API key middleware to all routes
	api.Use(middleware.ValidateAPIKey)

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("It runs!")
	})

	// Register all routes
	registerPlatformRoutes(api, platformController)
	registerGameRoutes(api, gameController)
	registerGenreRoutes(api, genreController)
	registerDeveloperRoutes(api, developerController)
	registerPublisherRoutes(api, publisherController)

	log.Println("✅ API routes initialized successfully")
}

func registerPlatformRoutes(api fiber.Router, c *controllers.PlatformController) {
	api.Get("/platforms", c.GetPlatforms)
	api.Get("/platforms/:id", c.GetPlatformByID)
	api.Post("/platforms", middleware.ValidateRequestBody(&models.Platform{}), c.CreatePlatform)
	api.Put("/platforms/:id", middleware.ValidateRequestBody(&models.Platform{}), c.UpdatePlatform)
	api.Delete("/platforms/:id", c.DeletePlatform)
}

func registerGameRoutes(api fiber.Router, c *controllers.GameController) {
	api.Get("/games", c.GetGames)
	api.Get("/games/:id", c.GetGameByID)
	api.Post("/games", middleware.ValidateRequestBody(&models.Game{}), c.CreateGame)
	api.Put("/games/:id", middleware.ValidateRequestBody(&models.Game{}), c.UpdateGame)
	api.Delete("/games/:id", c.DeleteGame)
}

func registerGenreRoutes(api fiber.Router, c *controllers.GenreController) {
	api.Get("/genres", c.GetGenres)
	api.Get("/genres/:id", c.GetGenreByID)
	api.Post("/genres", middleware.ValidateRequestBody(&models.Genre{}), c.CreateGenre)
	api.Put("/genres/:id", middleware.ValidateRequestBody(&models.GenreUpdate{}), c.UpdateGenre)
	api.Delete("/genres/:id", c.DeleteGenre)
}

func registerDeveloperRoutes(api fiber.Router, c *controllers.DeveloperController) {
	api.Get("/developers", c.GetDevelopers)
	api.Get("/developers/:id", c.GetDeveloperByID)
	api.Post("/developers", middleware.ValidateRequestBody(&models.Developer{}), c.CreateDeveloper)
	api.Put("/developers/:id", middleware.ValidateRequestBody(&models.DeveloperUpdate{}), c.UpdateDeveloper)
	api.Delete("/developers/:id", c.DeleteDeveloper)
}

func registerPublisherRoutes(api fiber.Router, c *controllers.PublisherController) {
	api.Get("/publishers", c.GetPublishers)
	api.Get("/publishers/:id", c.GetPublisherByID)
	api.Post("/publishers", middleware.ValidateRequestBody(&models.Publisher{}), c.CreatePublisher)
	api.Put("/publishers/:id", middleware.ValidateRequestBody(&models.PublisherUpdate{}), c.UpdatePublisher)
	api.Delete("/publishers/:id", c.DeletePublisher)
}
