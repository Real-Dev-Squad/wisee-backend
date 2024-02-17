package dtos

import (
	"time"

	"github.com/Real-Dev-Squad/wisee-backend/src/models"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

type ResponseDto struct {
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Error   ErrorResponse `json:"error"`
}

type CreateUpdateGetFormResponseDto struct {
	Id          int64              `json:"id"`
	Content     models.FormContent `json:"content"`
	OwnerId     int64              `json:"owner_id"`
	CreatedById int64              `json:"created_by_id"`
	Status      string             `json:"status"`
	UpdatedById *int64             `json:"updated_by_id"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}

type GetFormMetaDataResponseDto struct {
	Id                               int64     `json:"id"`
	FormId                           int64     `json:"form_id"`
	IsDeleted                        bool      `json:"is_deleted"`
	AccepctingResponses              bool      `json:"accepting_responses"`
	AllowGuestResponses              bool      `json:"allow_guest_responses"`
	AllowMultipleRepsonses           bool      `json:"allow_multiple_responses"`
	SendConfirmationEmailToRespondee bool      `json:"send_confirmation_email_to_respondee"`
	SendSubmissionEmailToOwner       bool      `json:"send_submission_email_to_owner"`
	ValidTill                        time.Time `json:"valid_till"`
	UpdatedById                      *int64    `json:"updated_by_id"`
	// TODO invite code
	UpdatedAt time.Time `json:"updated_at"`
}

type GetFormsResponseDto []CreateUpdateGetFormResponseDto

type GetFormDetailResponseDto struct {
	Id          int64                      `json:"id"`
	OwnerId     int64                      `json:"owner_id"`
	Status      string                     `json:"status"`
	CreatedById int64                      `json:"created_by_id"`
	UpdatedById *int64                     `json:"updated_by_id"`
	CreatedAt   string                     `json:"created_at"`
	UpdatedAt   string                     `json:"updated_at"`
	Content     models.FormContent         `json:"content"`
	Meta        GetFormMetaDataResponseDto `json:"meta"`
}
