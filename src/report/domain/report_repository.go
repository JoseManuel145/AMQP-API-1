package domain

import (
	"report/src/report/domain/entities"
)

type IReport interface {
	Create(title, content string) error
	ViewOne(id int) (*entities.Report, error)
	ViewAll() ([]entities.Report, error)
}
