package controllers

import (
	"beego-data-api/models"

	"github.com/astaxie/beego"
)

// Operations about Export
type ExportController struct {
	beego.Controller
}

// @Title Get
// @Description データ抽出API
// @Param	service	query 	string	true	"抽出対象のサービスを指定"
// @Param	token	query 	string	true	"抽出対象のアカウントのトークンを指定"
// @Param	from	query 	string	false	"抽出対象期間の開始日時"
// @Param	to		query 	string	false	"抽出対象期間の終了日時"
// @Param	output	query 	string	false	"レスポンス形式を指定 (json|xml|etc...)"
// @Success 200 {object} models.AppsflyerReport
// @Success 200 {object} models.GoogleReport
// @Failure 400 no enough input
// @router / [get]
func (e *ExportController) Get() {
	var (
		data interface{}
		err  error
	)

	service := e.GetString("service")
	// token := e.GetString("token")
	from := e.GetString("from")
	to := e.GetString("to")
	output := e.GetString("output")

	// TODO: validation

	switch service {
	case "appsflyer":
		data, err = models.GetAppsflyerReportList(from, to)
	default:
		data, err = models.GetGoogleReportList(from, to)
	}

	if err != nil {
		e.Data[output] = err.Error()
	} else {
		e.Data[output] = data
	}

	switch output {
	case "xml":
		e.ServeXML()
	default:
		e.ServeJSON()
	}
}
