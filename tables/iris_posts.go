package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIrisPostsTable(ctx *context.Context) table.Table {

	irisPosts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := irisPosts.GetInfo()

	info.AddField("ID", "id", db.Bigint).FieldSortable()
	info.AddField("标题", "title", db.Varchar)
	info.AddField("发布时间", "created_at", db.Timestamp)
	info.AddField("阅读", "views", db.Int)

	info.SetTable("iris_posts").SetTitle("IrisPosts").SetDescription("IrisPosts")

	formList := irisPosts.GetForm()
	formList.AddField("标题", "title", db.Varchar, form.Text).FieldMust()
	formList.AddField("内容", "content", db.Longtext, form.RichText).FieldMust()
	formList.AddField("阅读", "views", db.Longtext, form.Text).FieldMust()

	formList.SetTable("iris_posts").SetTitle("IrisPosts").SetDescription("IrisPosts")

	return irisPosts
}
