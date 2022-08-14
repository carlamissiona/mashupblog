package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RenderFrontend(c *fiber.Ctx) error {
	 
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}

func RenderHome(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	},)
}

func RenderSingle(c *fiber.Ctx) error {
	log.Println("Singel PAge!")
	return c.Render("view", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderLogin(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}

func RenderLogout(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}
func RenderIndex(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}
func RenderDesign(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}
func RenderAbout(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("frontend", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}