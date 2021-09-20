package controller

import (
	"errors"

	"github.com/GoAdminGroup/go-admin/template"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/service"
)

type IndexController struct {
	Ctx         iris.Context
	BlogService service.BlogService
	PageService service.PageService
	SiteService service.SiteService
	MenuService service.MenuService
}

func (c *IndexController) Get() {

	site, site_err := c.SiteService.GetSite()

	data, err := c.BlogService.GetLatest()

	menu, menu_err := c.MenuService.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Content", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)

		c.Ctx.ViewData("posts", data)
		c.Ctx.ViewData("menu", menu)
		c.Ctx.View("home.html")
	}

}

func (c *IndexController) GetLogin() {

	const content = `login page`

	c.Ctx.ViewData("title", "Home Page")
	c.Ctx.ViewData("content", template.HTML(content))
	c.Ctx.View("home.html")
}

func (c *IndexController) GetByWildcard(slug string) {

	page, err := c.PageService.GetPage(slug)
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.StopWithError(iris.StatusBadRequest, iris.ErrNotFound)
	} else {
		c.Ctx.ViewData("Title", page.Title)
		c.Ctx.ViewData("Content", template.HTML(page.Content))

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)
		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("page.html")
	}
}
