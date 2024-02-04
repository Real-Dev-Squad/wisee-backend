package dtos

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
	ID          int64   `json:"id"`
	Content     []Block `json:"content"`
	OwnerId     int64   `json:"ownerId"`
	CreatedByID int64   `json:"createdById"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetFormMetaDataResponseDto struct {
	FormID              int64  `json:"formId"`
	AcceptingResponses  bool   `json:"acceptingResponses"`
	AllowGuestResponses bool   `json:"allowGuestResponses"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
}

type CreateFormResponseDto struct {
	Content      []Block `json:"content"`
	ResponseById int64   `json:"responseById"`
	FormId       int64   `json:"formId"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}
