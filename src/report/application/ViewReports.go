package application

import (
	"report/src/report/domain/entities"
	"report/src/report/domain/repositories"
)

type ViewReportsUseCase struct {
	db repositories.IReport
}

func NewViewReports(db repositories.IReport) *ViewReportsUseCase {
	return &ViewReportsUseCase{db: db}
}

func (uc *ViewReportsUseCase) Execute() ([]entities.Report, error) {
	return uc.db.ViewAll()
}
