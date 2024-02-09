package dtos

type CreateUpdateFormRequestDto struct {
	Status  string  `json:"status"`
	Content []Block `json:"content"`
	OwnerId int64   `json:"ownerId"`
}

type CreateFormSubmissionRequestDto struct {
	Content      []Block `json:"content"`
	ResponseById int64   `json:"responseById"`
	FormId       int64   `json:"formId"`
}
