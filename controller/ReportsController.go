package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pushAppCsvGenerator/constant"
	"pushAppCsvGenerator/services"
	"pushAppCsvGenerator/useCase"
)

type ReportsController struct {
	route string
}

func CreateReportsController() *ReportsController {

	reportsControllerRoute := fmt.Sprintf("%s/reports", constant.ApiRoute)

	return &ReportsController{
		route: reportsControllerRoute,
	}
}

func (reportsController *ReportsController) SetupRoute(mux *chi.Mux) {
	mux.Route(reportsController.route, func(r chi.Router) {
		r.Get("/", reportsController.getAllReports)
	})
}

func (reportsController *ReportsController) getAllReports(w http.ResponseWriter, r *http.Request) {
	repository := services.CreateFirebaseReportsRepository()
	getAllReportsUseCase := useCase.CreateGetAllReportsUseCase(repository)

	reports, getAllReportsError := getAllReportsUseCase.Execute()
	if getAllReportsError != nil {
		http.Error(w, getAllReportsError.Error(), http.StatusNotFound)
		return
	}

	jsonSerializeError := json.NewEncoder(w).Encode(reports)
	if jsonSerializeError != nil {
		http.Error(w, jsonSerializeError.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
