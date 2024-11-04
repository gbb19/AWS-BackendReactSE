package routes

import (
	"onez19/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// auth := app.Group("/auth")
	// auth.Post("/register", controllers.Register)
	// auth.Post("/login", controllers.Login)
	app.Get("/users", controllers.GetUsers)
	app.Get("/workspaces/:workspaceId/sections", controllers.GetAllSectionsByWorkspaceID)
	app.Get("/users/:username/workspaces", controllers.GetAllWorkspacesByUsername)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/workspaces/:workspaceId/sections/:sectionId/activities", controllers.GetActivitiesBySectionAndWorkspace)

	app.Post("/users/:username/workspaces/create", controllers.CreateWorkspace)
	app.Post("/users/:username/workspaces/:workspace_id/join", controllers.JoinWorkspace)
	app.Post("/workspaces/:workspaceId/sections/create", controllers.CreateSection)
	app.Post("/workspaces/:workspaceId/:sectionId/activities/create", controllers.CreateActivity)
	app.Post("/workspaces/:workspaceId/activities/move", controllers.MoveActivity)
	app.Post("/workspaces/:workspaceId/:sectionId/activities/:activityId/edit", controllers.EditActivity)
	app.Post("/workspaces/:workspaceId/:sectionId/edit", controllers.EditSectionName)
	app.Post("/register", controllers.Register)
}
