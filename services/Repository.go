package services

import "pushAppCsvGenerator/model"

type ReportsRepository interface {
	GetReportById(id string) (model.ReportModel, error)
	GetAll() ([]model.ReportModel, error)
}
