package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIrisPagesTable(ctx *context.Context) table.Table {

	irisPages := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := irisPages.GetInfo()

	info.AddField("ID", "id", db.Bigint).FieldSortable()
	info.AddField("页面标题", "title", db.Varchar)
	info.AddField("内容", "content", db.Longtext)
	info.AddField("路径", "slug", db.Varchar)

	info.SetTable("iris_pages").SetTitle("IrisPages").SetDescription("IrisPages")

	formList := irisPages.GetForm()
	formList.AddField("页面标题", "title", db.Varchar, form.Text).FieldMust()
	formList.AddField("内容", "content", db.Longtext, form.RichText).FieldMust()
	formList.AddField("路径", "slug", db.Varchar, form.Text).FieldMust()

	formList.SetTable("iris_pages").SetTitle("IrisPages").SetDescription("IrisPages")

	return irisPages
}
