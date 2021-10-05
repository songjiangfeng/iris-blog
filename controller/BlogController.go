package controller

import (
	"github.com/GoAdminGroup/go-admin/template"
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

	if err != nil || site_err != nil || menu_err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("site", site)

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

	if err != nil || site_err != nil || menu_err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, iris.ErrNotFound)
	} else {

		c.Ctx.ViewData("site", site)

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
	prevpost, _ := c.BlogService.GetPrevPost(id)
	nextpost, _ := c.BlogService.GetNextPost(id)

	if err != nil || site_err != nil || menu_err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("Content", template.HTML(data.Content))
		c.Ctx.ViewData("post", data)

		c.Ctx.ViewData("site", site)

		c.Ctx.ViewData("menu", menu)

		//blog view count plus 1 each visit
		c.BlogService.ViewPlus(id)

		c.Ctx.ViewData("PrevPost", prevpost)

		c.Ctx.ViewData("NextPost", nextpost)

		c.Ctx.View("blogdetail.html")
	}
}

func (c *BlogController) GetHot() {
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.BlogService.GetHot()

	if err != nil || site_err != nil || menu_err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("site", site)

		c.Ctx.ViewData("posts", data)
		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("bloghot.html")
	}
}
