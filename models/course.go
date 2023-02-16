package models

import (
	"fiber-layout-mvc/initalize"
)

type Course struct {
	Cid   string
	Cname string
	Tid   string
}

func NewCourse() *Course {
	return &Course{}
}

// GetList 列表
func (t *Course) Home() ([]Course, error) {
	var result []Course
	if err := initalize.DB.Raw("select * from Course LIMIT 10").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// Category 详情
func (t *Course) Category(id string) (*Course, error) {
	if err := initalize.DB.Raw("select * from Course WHERE CId = ? LIMIT 10", id).Find(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
