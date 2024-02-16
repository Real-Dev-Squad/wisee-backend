package dtos

import "github.com/Real-Dev-Squad/wisee-backend/src/models"

type CreateUpdateFormRequestDto struct {
	Status        string             `json:"status"`
	Content       models.FormContent `json:"content"`
	PerformedById int64              `json:"performed_by_id"`
}

type CreateFormSubmissionRequestDto struct {
	Content      models.FormContent `json:"content"`
	ResponseById int64              `json:"reponse_by_id"`
	FormId       int64              `json:"form_id"`
}
