// @APIVersion 0.0.1
// @Title Data Export API
// @Description データ抽出API
// @Contact xio.dicek.i11@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-data-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/export",
			beego.NSInclude(
				&controllers.ExportController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
