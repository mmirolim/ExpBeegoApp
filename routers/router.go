package routers

import (
	"expbeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/get", &controllers.MainController{}, "get:Get")
	beego.Router("/post", &controllers.MainController{}, "post:Post")
}
