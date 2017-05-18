// @APIVersion 1.0.0
// @Title bussiness-dao-service API
// @Description bussiness-dao-service only serve the PROJECT_T/JOB_T/MODULE_T/POD_T
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dao-service/bussiness-dao-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/job",
			beego.NSInclude(
				&controllers.JobController{},
			),
		),
		beego.NSNamespace("/module",
			beego.NSInclude(
				&controllers.ModuleController{},
			),
		),
		beego.NSNamespace("/pod",
			beego.NSInclude(
				&controllers.PodController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
