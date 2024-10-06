package models

import (
	"github.com/chinmay211997/RiskManagerApp/utils"
	"gopkg.in/go-playground/validator.v9"
)

type CreateRiskRequest struct {
	State       RiskState `json:"state" validate:"required,oneof=open closed accepted investigating"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

func (s *CreateRiskRequest) ValidateCreateRiskRequest() bool {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		utils.Logger.Error(err.Error())
		return false
	}
	return true
}
