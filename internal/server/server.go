package server

import (
	"github.com/gofiber/fiber/v2"

	"dotcom-2025/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "dotcom-2025",
			AppName:      "dotcom-2025",
		}),

		db: database.New(),
	}

	return server
}
