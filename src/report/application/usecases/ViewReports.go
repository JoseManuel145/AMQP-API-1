package usecases

import (
	"report/src/report/domain"
	"report/src/report/domain/entities"
)

type ViewReportsUseCase struct {
	db domain.IReport
}

func NewViewReports(db domain.IReport) *ViewReportsUseCase {
	return &ViewReportsUseCase{db: db}
}

func (uc *ViewReportsUseCase) Execute() ([]entities.Report, error) {
	return uc.db.ViewAll()
}
