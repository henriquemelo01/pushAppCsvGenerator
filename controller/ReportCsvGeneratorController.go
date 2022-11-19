package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"pushAppCsvGenerator/constant"
	"pushAppCsvGenerator/useCase"
)

type ReportCsvGeneratorController struct {
	route                    string
	exportReportToCsvUseCase *useCase.ExportReportToCsvUseCase
}

func CreateReportCsvGeneratorController(
	exportReportToCsvUseCase *useCase.ExportReportToCsvUseCase,
) *ReportCsvGeneratorController {

	reportCsvControllerRoute := fmt.Sprintf("%s/reportCsvGenerator", constant.ApiRoute)

	return &ReportCsvGeneratorController{
		route:                    reportCsvControllerRoute,
		exportReportToCsvUseCase: exportReportToCsvUseCase,
	}
}

func (reportCsvController *ReportCsvGeneratorController) SetupRouter(mux *chi.Mux) {
	mux.Route(reportCsvController.route, func(r chi.Router) {
		r.Get("/{id}", reportCsvController.exportReportToCsv)
	})
}

func (reportCsvController *ReportCsvGeneratorController) exportReportToCsv(w http.ResponseWriter, r *http.Request) {

	reportId := chi.URLParam(r, "id")

	exportReportToCsvUseCase := reportCsvController.exportReportToCsvUseCase

	reportCsvFile, exportReportToCsvError := exportReportToCsvUseCase.Execute(reportId)
	if exportReportToCsvError != nil {
		log.Printf("Erro exportReportToCsvError: %s\n", exportReportToCsvError)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, reportCsvFile.FileName))
	_, _ = w.Write(reportCsvFile.FileBuffer.Bytes())

	return
}
