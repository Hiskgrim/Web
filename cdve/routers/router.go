// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cdve/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/carga_academica",
			beego.NSInclude(
				&controllers.CargaAcademicaController{},
			),
		),

		beego.NSNamespace("/tipo_proyecto_curricular",
			beego.NSInclude(
				&controllers.TipoProyectoCurricularController{},
			),
		),

		beego.NSNamespace("/tipo_contratacion",
			beego.NSInclude(
				&controllers.TipoContratacionController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/facultad",
			beego.NSInclude(
				&controllers.FacultadController{},
			),
		),

		beego.NSNamespace("/docente",
			beego.NSInclude(
				&controllers.DocenteController{},
			),
		),

		beego.NSNamespace("/categoria",
			beego.NSInclude(
				&controllers.CategoriaController{},
			),
		),

		beego.NSNamespace("/cancelacion_contrato",
			beego.NSInclude(
				&controllers.CancelacionContratoController{},
			),
		),

		beego.NSNamespace("/funcionario",
			beego.NSInclude(
				&controllers.FuncionarioController{},
			),
		),

		beego.NSNamespace("/estado_contrato",
			beego.NSInclude(
				&controllers.EstadoContratoController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/contrato_especial",
			beego.NSInclude(
				&controllers.ContratoEspecialController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),

		beego.NSNamespace("/clasificacion",
			beego.NSInclude(
				&controllers.ClasificacionController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
