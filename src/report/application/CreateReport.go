package application

import (
	"report/src/report/domain/entities"
	"report/src/report/domain/repositories"
)

type CreateReportUseCase struct {
	repo repositories.IReport
}

func NewCreateReportUseCase(r repositories.IReport) *CreateReportUseCase {
	return &CreateReportUseCase{
		repo: r,
	}
}
func (uc *CreateReportUseCase) Execute(report entities.Report) error {
	err := uc.repo.Create(report.ID, report.Title, report.Content)
	if err != nil {
		return err
	}
	return nil
}
