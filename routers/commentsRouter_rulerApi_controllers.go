package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["rulerApi/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:DominioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:DominioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:DominioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:DominioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:DominioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:PredicadoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerApi/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
