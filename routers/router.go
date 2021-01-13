// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/test-crud",
			beego.NSRouter("/read", &controllers.DogsController{}, "get:ReadDogs"),
			beego.NSRouter("/create", &controllers.DogsController{}, "post:CreateDog"),
			beego.NSRouter("/update", &controllers.DogsController{}, "put:UpdateDog"),
			beego.NSRouter("/delete", &controllers.DogsController{}, "delete:DeleteDog"),
		),
		beego.NSRouter("/breed-list", &controllers.DogsController{}, "get:BreedList"),
		beego.NSRouter("/breed-detail/:id", &controllers.DogsController{}, "get:BreedDetail"),
	)
	beego.AddNamespace(ns)
}
