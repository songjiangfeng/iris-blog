package controller

import (
	"errors"

	"github.com/GoAdminGroup/go-admin/template"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/service"
)

type BlogController struct {
	Ctx         iris.Context
	BlogService service.BlogService
	PageService service.PageService
	SiteService service.SiteService
	MenuService service.MenuService
}

const pagesize = 10

func (c *BlogController) Get() {
	page := c.Ctx.URLParamInt64Default("page", 1)
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.BlogService.GetPostByPage(page, pagesize)
	pageNext := page + 1
	pagePrev := page - 1

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)

		c.Ctx.ViewData("posts", data)
		c.Ctx.ViewData("menu", menu)
		c.Ctx.ViewData("PageNo", page)
		c.Ctx.ViewData("PageNext", pageNext)
		c.Ctx.ViewData("PagePrev", pagePrev)
		c.Ctx.View("blog.html")
	}
}

func (c *BlogController) GetPageBy(page int64) {

	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.BlogService.GetPostByPage(page, pagesize)
	pageNext := page + 1
	pagePrev := page - 1

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.StopWithError(iris.StatusBadRequest, iris.ErrNotFound)
	} else {

		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)

		c.Ctx.ViewData("posts", data)
		c.Ctx.ViewData("menu", menu)
		c.Ctx.ViewData("PageNo", page)
		c.Ctx.ViewData("PageNext", pageNext)
		c.Ctx.ViewData("PagePrev", pagePrev)

		c.Ctx.View("blog.html")

	}
}

func (c *BlogController) GetBy(id int64) {

	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.BlogService.GetPost(id)

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(site_err, gorm.ErrRecordNotFound) || errors.Is(menu_err, gorm.ErrRecordNotFound) {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.ViewData("Title", data.Title)
		c.Ctx.ViewData("ID", data.ID)
		c.Ctx.ViewData("Created_at", data.Created_at)
		c.Ctx.ViewData("Content", template.HTML(data.Content))
		c.Ctx.ViewData("SiteName", site.SiteName)
		c.Ctx.ViewData("SiteEmail", site.SiteEmail)
		c.Ctx.ViewData("Slogan", site.Slogan)
		c.Ctx.ViewData("Notice", site.Notice)
		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("blogdetail.html")
	}
}
