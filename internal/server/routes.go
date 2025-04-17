package server

import (
	"dotcom-2025/cmd/web"
	"dotcom-2025/cmd/web/home"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/", adaptor.HTTPHandler(templ.Handler(home.Home())))

	s.App.Get("/health", s.healthHandler)

	s.App.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Files),
		PathPrefix: "assets",
		Browse:     false,
	}))

	s.App.Get("/web", adaptor.HTTPHandler(templ.Handler(web.HelloForm())))

	s.App.Post("/hello", func(c *fiber.Ctx) error {
		return web.HelloWebHandler(c)
	})

	s.App.Get("/api/projects", s.ApiProjects)
	s.App.Get("/api/educations", s.ApiEducation)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) ApiProjects(c *fiber.Ctx) error {
	// return c.SendFile("cmd/web/assets/projects.json")
	data, err := os.ReadFile("cmd/web/assets/projects.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read projects.json")
	}
	return c.Type("json").Send(data)
}

func (s *FiberServer) ApiEducation(c *fiber.Ctx) error {
	data, err := os.ReadFile("cmd/web/assets/education.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read education.json")
	}
	return c.Type("json").Send(data)
}
