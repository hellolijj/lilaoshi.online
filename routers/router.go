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
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),

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
	beego.Post("/api/login/account", controllers.UserController{}.Login)
	beego.Get("/api/currentUser", controllers.UserController{}.Current)

}
