package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/songjiangfeng/iris-blog/ip"
	"github.com/songjiangfeng/iris-blog/service"
)

// TagController
type TagController struct {
	Ctx         iris.Context
	SiteService service.SiteService
	MenuService service.MenuService
	TagService  service.TagService
	BlogService service.BlogService
}

// Get
//  @receiver c
func (c *TagController) Get() {

	data, err := c.TagService.GetAll()
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()

	if err != nil || site_err != nil || menu_err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {

		c.Ctx.ViewData("site", site)

		c.Ctx.ViewData("tags", data)
		c.Ctx.ViewData("menu", menu)
		c.Ctx.ViewData("ip",  ip.IP())

		c.Ctx.View("tag.html")
	}
}

// GetBy
//  @receiver c
//  @param id
func (c *TagController) GetBy(id int64) {

	tag, tag_err := c.TagService.GetTag(id)
	site, site_err := c.SiteService.GetSite()
	menu, menu_err := c.MenuService.GetAll()
	data, err := c.BlogService.GetPostByTagID(id)

	if err != nil || site_err != nil || menu_err != nil || tag_err != nil {
		c.Ctx.JSON(iris.Map{"message": "记录不存在", "status": iris.StatusNotFound})
	} else {
		c.Ctx.ViewData("ID", tag.ID)
		c.Ctx.ViewData("Name", tag.Name)
		c.Ctx.ViewData("ip",  ip.IP())
		c.Ctx.ViewData("tagposts", data)

		c.Ctx.ViewData("site", site)

		c.Ctx.ViewData("menu", menu)

		c.Ctx.View("tagdetail.html")
	}
}
