package models

import (
	"time"
)

var (
	GoogleReportList []*GoogleReport
)

type GoogleReport struct {
	Id       int64
	Param1   string
	Param2   string
	Param3   string
	Param4   string
	Param5   string
	Datetime string
}

func init() {
	GoogleReportList = make([]*GoogleReport, 0)
	GoogleReportList = append(GoogleReportList, &GoogleReport{1, "param1", "param2", "param3", "param4", "param5", time.Now().Format("2006-01-02 15:04:05")})
	GoogleReportList = append(GoogleReportList, &GoogleReport{2, "param1", "param2", "param3", "param4", "param5", time.Now().Format("2006-01-02 15:04:05")})
}

func GetGoogleReportList(from string, to string) (g []*GoogleReport, err error) {
	return GoogleReportList, nil
}
