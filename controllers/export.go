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
// @Description get data
// @Param	service	query 	string	true	"The key for staticblock"
// @Param	token	query 	string	true	"The key for staticblock"
// @Param	from	query 	string	false	"The key for staticblock"
// @Param	to		query 	string	false	"The key for staticblock"
// @Param	output	query 	string	false	"The key for staticblock"
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
