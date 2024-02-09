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
	ID            int64   `json:"id"`
	Content       []Block `json:"content"`
	OwnerId       int64   `json:"ownerId"`
	CreatedById   int64   `json:"createdById"`
	Status        string  `json:"status"`
	PerformedById int64   `json:"performedById"`
}

type GetFormMetaDataResponseDto struct {
	FormId              int64 `json:"formId"`
	AcceptingResponses  bool  `json:"acceptingResponses"`
	AllowGuestResponses bool  `json:"allowGuestResponses"`
}
