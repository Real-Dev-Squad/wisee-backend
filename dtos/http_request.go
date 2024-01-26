package dtos

type FormCreationRequest struct {
	Title  string `json:"title"`
	UserId int64  `json:"userId"`
}

type FormUpdateRequest struct {
	Title                 string `json:"title"`
	SendEmailOnSubmission bool   `json:"sendEmailOnSubmission"`
	UserId                int64  `json:"userId"`
}

// use for creating/updating question
type QuestionUpdateRequest struct {
	FormId                  int64  `json:"formId"`
	Title                   string `json:"title"`
	Type                    string `json:"type"`
	Points                  int    `json:"points"`
	IsReq                   bool   `json:"isRequired"`
	IsPartialMarkingEnabled bool   `json:"isPartialMarkingEnabled"`
}

// use for creating/updating option
type OptionUpdateRequest struct {
	QuestionId int64  `json:"questionId"`
	Value      string `json:"value"`
}
