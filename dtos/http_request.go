package dtos

type CreateUpdateFormRequestDto struct {
	Title                 string `json:"title"`
	SendEmailOnSubmission bool   `json:"sendEmailOnSubmission"`
	UserId                int64  `json:"userId"`
}

type CreateUpdateQuestionRequestDto struct {
	FormId     int64  `json:"formId"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	IsRequired bool   `json:"isRequired"`
	QuestionId int64  `json:"questionId"`
}

type CreateUpdateOptionRequestDto struct {
	QuestionId int64  `json:"questionId"`
	Value      string `json:"value"`
	OptionId   int64  `json:"optionId"`
}
