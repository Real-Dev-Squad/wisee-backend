package dtos

import "github.com/Real-Dev-Squad/wisee-backend/src/models"

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

type ResponseDto struct {
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Error   ErrorResponse `json:"error"`
}

type CreateUpdateFormResponseDto struct {
	ID          int64              `json:"id"`
	Content     models.FormContent `json:"content"`
	OwnerId     int64              `json:"owner_id"`
	CreatedById int64              `json:"created_by_id"`
	Status      string             `json:"status"`
	UpdatedById int64              `json:"updated_by_id"`
}

type GetFormMetaDataResponseDto struct {
	FormId              int64 `json:"form_id"`
	AcceptingResponses  bool  `json:"accepting_responses"`
	AllowGuestResponses bool  `json:"allow_guest_responses"`
}
