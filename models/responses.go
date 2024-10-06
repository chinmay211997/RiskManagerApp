package models

import "github.com/google/uuid"

type CreateRiskResponse struct {
	RiskID  uuid.UUID `json:"risk_id"`
	Message string    `json:"message"`
	Code    int       `json:"statusCode"`
}

type GetRisksListResponse struct {
	RiskListData []Risk `json:"data"`
	TotalLength  int    `json:"total"`
	Code         int    `json:"statusCode"`
}

const RISK_CREATE_SUCCESS = "Risk Created Successfully!"
