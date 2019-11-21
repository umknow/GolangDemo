// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	. "fmt"
	"quickapi/controllers"

	"github.com/astaxie/beego"
)


func init() {
	// 自带的无限嵌套
	/*
	支持了如下这样的请求 URL
	GET /v1/notallowed
	GET /v1/version
	GET /v1/changepassword
	POST /v1/changepassword
	GET /v1/shop/123
	GET /v1/cms/ 对应 MainController、CMSController、BlockController 中得注解路由
	而且还支持前置过滤,条件判断,无限嵌套 namespace
	*/
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		// 自己添加
		//beego.NSRouter("/login", &logincontrollers.LoginWebController{}, "get:Login"),

		)
	beego.AddNamespace(ns)

	// 快速路由法
	// 自定义配置
	// 分时行情数据
	//beego.Router("/api/hq/minute/today", &controllers.HqController{},"*:GetDayPrice")
	Println("【router hq today】")
	hq := beego.NewNamespace("/api/hq",
		beego.NSNamespace("/minute",  // 分时

			beego.NSInclude(
				&controllers.HqController{},
			),
		),
		//beego.NSNamespace("/minute/five",  // 五日
		//	beego.NSInclude(
		//		&controllers.HqController{},
		//	),
		//),

		)

	//Println("【router hq today】")
	beego.AddNamespace(hq)
}
