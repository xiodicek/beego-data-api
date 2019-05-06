package models

import (
	"time"
)

var (
	AppsflyerReportList []*AppsflyerReport
)

type AppsflyerReport struct {
	Id       int64
	Param1   string
	Param2   string
	Param3   string
	Datetime string
}

func init() {
	AppsflyerReportList = make([]*AppsflyerReport, 0)
	AppsflyerReportList = append(AppsflyerReportList, &AppsflyerReport{1, "param1", "param2", "param3", time.Now().Format("2006-01-02 15:04:05")})
	AppsflyerReportList = append(AppsflyerReportList, &AppsflyerReport{2, "param1", "param2", "param3", time.Now().Format("2006-01-02 15:04:05")})
}

func GetAppsflyerReportList(from string, to string) (g []*AppsflyerReport, err error) {
	return AppsflyerReportList, nil
}
