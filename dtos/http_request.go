package dtos

type CreateUpdateFormRequestDto struct {
	Title                 string `json:"title"`
	SendEmailOnSubmission bool   `json:"sendEmailOnSubmission"`
	UserId                int64  `json:"userId"`
}

// use for creating/updating question
type CreateUpdateQuestionRequestDto struct {
	FormId int64  `json:"formId"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	// Points                  int    `json:"points"`
	isRequired bool  `json:"isRequired"`
	QuestionId int64 `json:"questionId"`
	// IsPartialMarkingEnabled bool   `json:"isPartialMarkingEnabled"`
}

// use for creating/updating option
type CreateUpdateOptionRequestDto struct {
	QuestionId int64  `json:"questionId"`
	Value      string `json:"value"`
	OptionId   int64  `json:"optionId"`
}
