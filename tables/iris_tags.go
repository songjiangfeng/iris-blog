package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIrisTagsTable(ctx *context.Context) table.Table {

	irisTags := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := irisTags.GetInfo()

	info.AddField("Id", "id", db.Bigint).FieldSortable()
	info.AddField("名称", "name", db.Varchar)

	info.SetTable("iris_tags").SetTitle("IrisTags").SetDescription("IrisTags")

	formList := irisTags.GetForm()
	formList.AddField("名称", "name", db.Varchar, form.Text).FieldMust()

	formList.SetTable("iris_tags").SetTitle("IrisTags").SetDescription("IrisTags")

	return irisTags
}
