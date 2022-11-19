package useCase

import (
	"log"
	"pushAppCsvGenerator/model"
	"pushAppCsvGenerator/services"
)

type ExportReportToCsvUseCase struct {
	repository                services.ReportsRepository
	reportCsvGeneratorService *services.ReportCsvGeneratorService
}

func CreateExportReportToCsvUseCase(
	repository services.ReportsRepository,
	reportCsvGeneratorService *services.ReportCsvGeneratorService,
) *ExportReportToCsvUseCase {
	return &ExportReportToCsvUseCase{
		repository:                repository,
		reportCsvGeneratorService: reportCsvGeneratorService,
	}
}

func (exportReportToCsvUseCase *ExportReportToCsvUseCase) Execute(id string) (*model.ReportCsvFile, error) {

	reportRepository := exportReportToCsvUseCase.repository

	reportModel, getReportModelByIdError := reportRepository.GetReportById(id)
	if getReportModelByIdError != nil {
		log.Println("getReportModelByIdError: ", getReportModelByIdError.Error())
		return &model.ReportCsvFile{}, getReportModelByIdError
	}

	reportCsvGeneratorService := exportReportToCsvUseCase.reportCsvGeneratorService
	reportCsvFile := reportCsvGeneratorService.Export(reportModel)

	return reportCsvFile, nil
}
