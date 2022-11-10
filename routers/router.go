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
	"strings"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
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
		beego.NSNamespace("/shoppingcar",
			beego.NSInclude(
				&controllers.ShoppingcarController{},
			),
		),
		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrdersController{},
			),
		),
	)

	beego.AddNamespace(ns)

	// auth验证
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		var pass = false
		if online, _ := user.Online(ctx); online {
			pass = true
		}
		if strings.Contains(ctx.Request.URL.Path, "/tools") {
			pass = true
		}

		if !pass && ctx.Request.URL.Path != "/api/login/account" {
			ctx.WriteString("请重新登录")
			return
		}
	})


	beego.Post("/api/login/account", controllers.UserController{}.Login)
	beego.Post("/api/user/logout", controllers.UserController{}.Logout)
	beego.Get("/api/currentUser", controllers.UserController{}.CurrentUser)  // 路由有问题
	beego.Post("/api/user/register", controllers.UserController{}.Register)

	beego.Post("/api/shoppingcar/op", controllers.ShoppingcarController{}.Op) // 加入到购物车
	beego.Get("/api/shoppingcar/goods", controllers.ShoppingcarController{}.Goods) // 加入到购物车

	beego.Post("/api/order/create", controllers.OrdersController{}.Create) // 下单
	beego.Post("/api/order/pay", controllers.OrdersController{}.Pay) // 下单

	beego.Get("/tools/craw", controllers.GoodController{}.Crawler)
}
