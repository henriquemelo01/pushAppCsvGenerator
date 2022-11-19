package di

import (
	"pushAppCsvGenerator/controller"
	"pushAppCsvGenerator/services"
	useCase2 "pushAppCsvGenerator/useCase"
)

func CreateReportCsvGeneratorController() *controller.ReportCsvGeneratorController {
	exportReportToCsvUseCase := createExportReportToCsvUseCase()
	return controller.CreateReportCsvGeneratorController(exportReportToCsvUseCase)
}

func createExportReportToCsvUseCase() *useCase2.ExportReportToCsvUseCase {

	// Criando as Dependencias do Controller
	repository := services.CreateFirebaseReportsRepository()
	csvGeneratorService := services.CreateReportCsvGeneratorService()

	return useCase2.CreateExportReportToCsvUseCase(repository, csvGeneratorService)
}
