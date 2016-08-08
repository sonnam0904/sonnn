package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sonnn/controllers:MainController"] = append(beego.GlobalControllerRouter["sonnn/controllers:MainController"],
		beego.ControllerComments{
			"Home",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sonnn/controllers:PostController"] = append(beego.GlobalControllerRouter["sonnn/controllers:PostController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sonnn/controllers:PostController"] = append(beego.GlobalControllerRouter["sonnn/controllers:PostController"],
		beego.ControllerComments{
			"Create",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sonnn/controllers:PostController"] = append(beego.GlobalControllerRouter["sonnn/controllers:PostController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sonnn/controllers:PostController"] = append(beego.GlobalControllerRouter["sonnn/controllers:PostController"],
		beego.ControllerComments{
			"Update",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["sonnn/controllers:PostController"] = append(beego.GlobalControllerRouter["sonnn/controllers:PostController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
