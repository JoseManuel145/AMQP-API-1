package application

import (
	"encoding/json"
	"report/src/report/domain/entities"
	"report/src/report/domain/repositories"
)

type SendMessageUseCase struct {
	rabbit repositories.IRabbit
}

func NewSendMessageUseCase(r repositories.IRabbit) *SendMessageUseCase {
	return &SendMessageUseCase{
		rabbit: r,
	}
}
func (uc *SendMessageUseCase) Execute(report entities.Report) error {
	reportJSON, err := json.Marshal(report)
	if err != nil {
		return err
	}
	errRabit := uc.rabbit.SendMessage(string(reportJSON))
	if errRabit != nil {
		return errRabit
	}
	return nil
}
