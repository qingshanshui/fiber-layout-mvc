package service

import (
	"fiber-layout-mvc/models"
	"fiber-layout-mvc/validator/form"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

// Home 首页
func (t *Default) Home() ([]models.Course, error) {
	list, err := models.NewCourse().GetList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Category 详情
func (t *Default) Category(c form.Category) (*models.Course, error) {
	list, err := models.NewCourse().Category(c.Id)
	if err != nil {
		return nil, err
	}
	return list, nil
}
