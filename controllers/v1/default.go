package v1

import (
	"fiber-layout-mvc/controllers"
	"fiber-layout-mvc/service"
	"fiber-layout-mvc/validator"
	"fiber-layout-mvc/validator/form"
	"github.com/gofiber/fiber/v2"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// Home 首页
func (t *DefaultController) Home(c *fiber.Ctx) error {
	home, err := service.NewDefaultService().Home()
	if err != nil {
		return err
	}
	return c.Render("index", fiber.Map{
		"Title": home,
	}, "layouts/main")
}

// Category 详情
func (t *DefaultController) Category(c *fiber.Ctx) error {
	// 初始化参数结构体
	categoryForm := form.Category{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &categoryForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Category(categoryForm)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api)) // => ✋ register
}
