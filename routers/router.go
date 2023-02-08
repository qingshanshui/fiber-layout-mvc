package routers

import (
	v1 "fiber-layout-mvc/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	group := app.Group("/")
	group.Get("/", main.Home)             // 首页
	group.Get("/category", main.Category) // 详情
}
