package controller

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/service"
)

type TagController struct {
	Ctx         iris.Context
	SiteService service.SiteService
	MenuService service.MenuService
	TagService  service.TagService
}

func (c *TagController) Get() {

	data, err := c.TagService.GetAll()
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)

		c.Ctx.ViewData("tags", data)
		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("tag.html")
	}
}

func (c *TagController) GetBy(id int64) {

	tag, tag_err := c.TagService.GetTag(id)
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.TagService.GetTagPost(id)

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(tag_err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.ViewData("ID", tag.ID)
		c.Ctx.ViewData("Name", tag.Name)
		c.Ctx.ViewData("tagposts", data)

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)

		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("tagdetail.html")
	}
}
