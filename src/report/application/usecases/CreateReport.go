package usecases

import (
	"report/src/report/domain"
)

type CreateReportUseCase struct {
	repo domain.IReport
}

func NewCreateReportUseCase(r domain.IReport) *CreateReportUseCase {
	return &CreateReportUseCase{
		repo: r,
	}
}

func (uc *CreateReportUseCase) Execute(id int, title, content string) error {
	// Llama al repositorio para crear el reporte.
	err := uc.repo.Create(title, content)
	if err != nil {
		return err
	}
	return nil
}
