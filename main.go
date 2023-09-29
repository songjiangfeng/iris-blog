package main

import (
	"os"
	"path/filepath"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/iris"
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/k3a/html2text"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/schicho/substring"
	"github.com/songjiangfeng/iris-blog/controller"
	"github.com/songjiangfeng/iris-blog/environment"
	"github.com/songjiangfeng/iris-blog/filemanager"
	"github.com/songjiangfeng/iris-blog/pages"
	"github.com/songjiangfeng/iris-blog/service"
	"github.com/songjiangfeng/iris-blog/tables"
)

func main() {
	startHttpsServer()
}

func startHttpsServer() {
	app := iris.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//sslcert :="./ssl/go365.tech.crt"
	//sslkey :="./ssl/go365.tech.key"
	config := "./config.json"
	if err := eng.AddConfigFromJSON(config).
		AddPlugins(filemanager.
			NewFileManager(filepath.Join(dir, "uploads")),
		).
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}

	service := service.MysqlService{}
	service.Init(eng.MysqlConnection())

	app.HandleDir("/uploads", iris.Dir("./uploads"), iris.DirOptions{
		Compress: true,
	})

	app.HandleDir("/assets", iris.Dir("./assets"), iris.DirOptions{
		Compress: true,
	})

	//public page
	tmpl := iris.HTML("./html", ".html").Reload(true)

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

	// http://localhost:8080

	//app.Run(iris.TLS(":443", sslcert, sslkey))
	app.Listen(":80")
}
