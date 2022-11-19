package useCase

import (
	"pushAppCsvGenerator/model"
	"pushAppCsvGenerator/services"
)

type GetAllReportsUseCase struct {
	repository services.ReportsRepository
}

func CreateGetAllReportsUseCase(repository services.ReportsRepository) *GetAllReportsUseCase {
	return &GetAllReportsUseCase{
		repository: repository,
	}
}

func (getAllReportsUseCase *GetAllReportsUseCase) Execute() ([]model.ReportModel, error) {
	repository := getAllReportsUseCase.repository

	reports, getAllReportsError := repository.GetAll()
	if getAllReportsError != nil {
		return nil, getAllReportsError
	}

	return reports, nil
}
