package dtos

import "github.com/Real-Dev-Squad/wisee-backend/src/models"

type CreateFormRequestDto struct {
	Content       models.FormContent `json:"content"`
	PerformedById int64              `json:"performed_by_id"`
}

type UpdateFormRequestDto struct {
	Status        string             `json:"status"`
	Content       models.FormContent `json:"content"`
	PerformedById int64              `json:"performed_by_id"`
}

type CreateFormSubmissionRequestDto struct {
	Content      models.FormContent `json:"content"`
	ResponseById int64              `json:"response_by_id"`
}
