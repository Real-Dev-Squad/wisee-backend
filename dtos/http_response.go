package dtos

type FormCreationResponse struct {
	Mesaage string `json:"message"`
	FormId  int64  `json:"formId"`
	Title   string `json:"title"`
}

type FormUpdateResponse struct {
	Mesaage               string `json:"message"`
	FormId                int64  `json:"formId"`
	SendEmailOnSubmission bool   `json:"sendEmailOnSubmission"`
	Title                 string `json:"title"`
}

// use for creating/updating question
type QuestionCreationResponse struct {
	Mesaage    string `json:"message"`
	QuestionId int64  `json:"questionId"`
	FormId     int64  `json:"formId"`
	Type       string `json:"type"`
	Title      string `json:"title"`
}

// use for creating/updating option
type OptionCreationResponse struct {
	Mesaage    string `json:"message"`
	OptionId   int64  `json:"optionId"`
	QuestionId int64  `json:"questionId"`
	Value      string `json:"value"`
}
