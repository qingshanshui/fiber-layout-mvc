package main

import (
	_ "fiber-layout-mvc/config"
	_ "fiber-layout-mvc/initalize"
	"fiber-layout-mvc/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New(fiber.Config{
		// 采用html模板引擎
		Views: html.New("./views", ".html"),
	})
	// 静态目录
	app.Static("/", "./static")
	// HTTP 请求/响应日志
	app.Use(logger.New())
	// Recover 中间件将可以堆栈链中的任何位置将 panic 恢复，并将处理集中到
	app.Use(recover.New())
	// 设置路由
	routers.SetRoute(app)
	// 监听端口
	_ = app.Listen(viper.GetString("App.Port"))
}
