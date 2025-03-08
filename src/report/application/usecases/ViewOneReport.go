package usecases

import (
	"report/src/report/domain"
	"report/src/report/domain/entities"
)

type ViewOneReportUseCase struct {
	repo domain.IReport
}

func NewViewOneReportUseCase(r domain.IReport) *ViewOneReportUseCase {
	return &ViewOneReportUseCase{repo: r}
}

func (uc *ViewOneReportUseCase) Execute(id int) (*entities.Report, error) {
	return uc.repo.ViewOne(id)
}
