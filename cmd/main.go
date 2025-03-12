package main

import (
	"log"
	"s0109-23/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	svc := service.New()

	api.Get("/", svc.RootHandler)
	api.Get("/api/user", svc.UserRead)
	api.Get("/api/profile", svc.ProfileRead)

	if err := api.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
