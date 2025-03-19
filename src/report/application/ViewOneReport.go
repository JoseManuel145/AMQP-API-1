package application

import (
	"report/src/report/domain/entities"
	"report/src/report/domain/repositories"
)

type ViewOneReportUseCase struct {
	repo repositories.IReport
}

func NewViewOneReportUseCase(r repositories.IReport) *ViewOneReportUseCase {
	return &ViewOneReportUseCase{repo: r}
}

func (uc *ViewOneReportUseCase) Execute(id int) (*entities.Report, error) {
	return uc.repo.ViewOne(id)
}
