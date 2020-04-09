package routers

import (
	"github.com/LiangQinghai/beegoStudyOne/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
