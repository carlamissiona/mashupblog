package router

import (
	"vestibulum/blog/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	pages := app.Group("", cors.New(), csrf.New())
	pages.Get("/", controllers.RenderHome)

	pages.Get("/vue", controllers.RenderFrontend)
	pages.Get("/view", controllers.RenderSingle)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
