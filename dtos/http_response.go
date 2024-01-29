package dtos

type ResponseDto struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type CreateUpdateFormResponseDto struct {
	FormId                int64  `json:"formId"`
	SendEmailOnSubmission bool   `json:"sendEmailOnSubmission"`
	Title                 string `json:"title"`
}

type CreateUpdateQuestionResponseDto struct {
	QuestionId int64  `json:"questionId"`
	FormId     int64  `json:"formId"`
	Type       string `json:"type"`
	Title      string `json:"title"`
}

// use for creating/updating option
type CreateUpdateOptionResponseDto struct {
	OptionId   int64  `json:"optionId"`
	QuestionId int64  `json:"questionId"`
	Value      string `json:"value"`
}
