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
	pages.Get("/", controllers.RenderFrontend)

	pages.Get("/api-coins", controllers.ApiNfb_Coin) 
	pages.Get("/api-blogs", controllers.ApiNfb_Blogs) 
	
	pages.Get("/view", controllers.RenderSingle)

	pages.Get("/about", controllers.RenderAbout)
	pages.Get("/design", controllers.RenderDesign)
    pages.Get("/svc-login", controllers.RenderLogin)
	pages.Get("/svc-logout", controllers.RenderLogout)
	
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
