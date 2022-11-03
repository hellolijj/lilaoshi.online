// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cookie-shop-api/controllers"
	"cookie-shop-api/services/user"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/type",
			beego.NSInclude(
				&controllers.TypeController{},
			),
		),

		beego.NSNamespace("/good",
			beego.NSInclude(
				&controllers.GoodController{},
			),
		),
	)

	beego.AddNamespace(ns)

	// auth验证
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		online, _ := user.Online(ctx)
		if !online && ctx.Request.URL.Path != "/api/login/account" {
			ctx.Redirect(302, "/user/login")
		}
	})


	beego.Post("/api/login/account", controllers.UserController{}.Login)
	beego.Post("/api/user/logout", controllers.UserController{}.Logout)
	beego.Get("/api/currentUser", controllers.UserController{}.CurrentUser)  // 路由有问题
	beego.Post("/api/user/register", controllers.UserController{}.Register)

}
