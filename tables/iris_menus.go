package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIrisMenusTable(ctx *context.Context) table.Table {

	irisMenus := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := irisMenus.GetInfo()

	info.AddField("Id", "id", db.Bigint).FieldSortable()
	info.AddField("名称", "name", db.Varchar)
	info.AddField("路径", "path", db.Varchar)
	info.AddField("权重", "weight", db.Int)

	info.SetTable("iris_menus").SetTitle("IrisMenus").SetDescription("IrisMenus")

	formList := irisMenus.GetForm()
	formList.AddField("名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("路径", "path", db.Varchar, form.Text).FieldMust()
	formList.AddField("权重", "weight", db.Varchar, form.Text).FieldMust()

	formList.SetTable("iris_menus").SetTitle("IrisMenus").SetDescription("IrisMenus")

	return irisMenus
}
