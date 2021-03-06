package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/service"
)

// UserController
type UserController struct {
	Ctx         iris.Context
	UserService service.UserService
}

// Get
//  @receiver c
func (c *UserController) Get() {
	data, err := c.UserService.GetAll()

	if err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.JSON(data)

	}
}

func (c *UserController) GetBy(id int64) {
	data, err := c.UserService.GetUser(id)

	if err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.JSON(data)
	}
}
