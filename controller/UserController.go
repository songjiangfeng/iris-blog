package controller

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/service"
)

type UserController struct {
	Ctx         iris.Context
	UserService service.UserService
}

func (c *UserController) Get() {
	data, err := c.UserService.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.JSON(data)

	}
}

func (c *UserController) GetBy(id int64) {
	data, err := c.UserService.GetUser(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.JSON(data)
	}
}
