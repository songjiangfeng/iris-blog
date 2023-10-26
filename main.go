package main

import (
	"os"
	"path/filepath"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/iris"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/k3a/html2text"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/rewrite"
	"github.com/kataras/iris/v12/mvc"

	"github.com/schicho/substring"
	"github.com/songjiangfeng/iris-blog/assets"
	"github.com/songjiangfeng/iris-blog/controller"
	"github.com/songjiangfeng/iris-blog/environment"
	"github.com/songjiangfeng/iris-blog/filemanager"
	"github.com/songjiangfeng/iris-blog/pages"
	"github.com/songjiangfeng/iris-blog/service"
	"github.com/songjiangfeng/iris-blog/tables"
)

//go:generate go-bindata -fs  -prefix "static" -o=./assets/assets.go -pkg=assets   ./static/...

var (
	configPath = "app/config.json"
	redirectPath = "app/redirects.yml"
)

func main() {

	f := config.ReadFromJson(configPath)
	app := newApp(f)
	env := f.Extra["env"]
	domain := f.Extra["domain"].(string)

	switch env {
	case config.EnvLocal:
		app.Logger().SetLevel("debug")
		app.Listen(":8888")
	case config.EnvProd:
		//let's encrypt
		app.Run(iris.AutoTLS(":443", domain, "admin@admin.com"))
	}
}

func newApp(f config.Config) *iris.Application {
	app := iris.New()
	redirects := rewrite.Load(redirectPath)
	app.WrapRouter(redirects)
	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	uploadsDir := f.Store.Path
	siteTheme := f.Extra["site_theme"]
	dir, _ := os.Getwd()

	if err := eng.AddConfigFromJSON(configPath).
		AddPlugins(filemanager.
			NewFileManager(filepath.Join(dir, uploadsDir)),
		).
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}
	service := service.MysqlService{}
	service.Init(eng.MysqlConnection())

	// asset & files
	app.HandleDir("/static", assets.AssetFile())
	app.HandleDir("/assets", iris.Dir(siteTheme.(string) + "/assets"))
	app.HandleDir("/files", iris.Dir(uploadsDir))

	//default theme

	tmpl := iris.HTML(siteTheme, ".html").Reload(true)

	app.RegisterView(tmpl.Layout("layout.html"))

	tmpl.AddFunc("html", func(s string) string {
		s = html2text.HTML2Text(s)
		s = substring.SubstringEnd(s, 50) + "..."

		return s
	})

	tmpl.AddFunc("timeformat", func(t time.Time) string {
		nt := t.Format(time.RFC3339)
		return nt[0:10]
	})

	mvc.Configure(app.Party("/"), func(app *mvc.Application) {
		// Register Dependencies.
		app.Register(
			environment.DEV, // DEV, PROD
		)
		// Register Controllers.
		app.Handle(new(controller.IndexController))
	})

	mvc.Configure(app.Party("/blog"), func(app *mvc.Application) {
		// Register Dependencies.

		app.Register(
			environment.DEV, // DEV, PROD
		)
		// Register Controllers.
		app.Handle(new(controller.BlogController))
	})

	mvc.Configure(app.Party("/tag"), func(app *mvc.Application) {
		// Register Dependencies.

		app.Register(
			environment.DEV, // DEV, PROD
		)
		// Register Controllers.
		app.Handle(new(controller.TagController))
	})

	mvc.Configure(app.Party("/user"), func(app *mvc.Application) {
		// Register Dependencies.
		app.Register(
			environment.DEV, // DEV, PROD
		)
		// Register Controllers.
		app.Handle(new(controller.UserController))
	})

	eng.HTML("GET", "/admin", pages.DashboardPage)
	// eng.HTML("GET", "/admin/form", pages.GetFormContent)
	// eng.HTML("GET", "/admin/table", pages.GetTableContent)
	return app
}
