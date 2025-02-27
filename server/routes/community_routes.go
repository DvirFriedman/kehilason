package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/bokoness/lashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommunitiesRoutes(app fiber.Router) {
	route := app.Group("/communities")

	route.Get("/:id", handlers.GetCommunity)
	route.Get("/", handlers.GetCommunities)
	route.Post("/", middleware.AuthMiddleware, handlers.CreateCommunity)
	route.Put("/:id", middleware.AuthMiddleware, handlers.UpdateCommunity)
	route.Delete("/:id", middleware.AuthMiddleware, handlers.DeleteCommunity)
	route.Post("/:id/register-user", middleware.AuthMiddleware, handlers.RegisterUserToCommunity)
}
