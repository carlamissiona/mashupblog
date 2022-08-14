package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RenderFrontend(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/fend")
}

func RenderHome(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderSingle(c *fiber.Ctx) error {
	log.Println("Singel PAge!")
	return c.Render("view", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
