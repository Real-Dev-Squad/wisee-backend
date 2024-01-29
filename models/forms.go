package models

import (
	"time"

	"github.com/uptrace/bun"
)

type FORM_TYPE_ENUM string

const (
	TEXT_SHORT    FORM_TYPE_ENUM = "TEXT_SHORT"
	TEXT_LONG     FORM_TYPE_ENUM = "TEXT_LONG"
	TEXT_NUMBER   FORM_TYPE_ENUM = "TEXT_NUMBER"
	TEXT_EMAIL    FORM_TYPE_ENUM = "TEXT_EMAIL"
	SINGLE_SELECT FORM_TYPE_ENUM = "SINGLE_SELECT"
	MULTI_SELECT  FORM_TYPE_ENUM = "MULTI_SELECT"
	FILE_IMAGE    FORM_TYPE_ENUM = "FILE_IMAGE"
	FILE_PDF      FORM_TYPE_ENUM = "FILE_PDF"
	LINEAR_SCALE  FORM_TYPE_ENUM = "LINEAR_SCALE"
)

type Form struct {
	bun.BaseModel `bun:"table:forms"`

	Id                    int64     `bun:"id,pk,autoincrement"`
	CreatedByID           int64     `bun:"created_by_id"`
	CreatedBy             *User     `bun:"rel:belongs-to,join:created_by_id=id"`
	Title                 string    `bun:"title,notnull"`
	SendEmailOnSubmission bool      `bun:"send_email_on_submission,default:false"`
	CreatedAt             time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt             time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Question struct {
	bun.BaseModel `bun:"table:questions"`

	Id                      int64          `bun:"id,pk,autoincrement"`
	FormID                  int64          `bun:"form_id"`
	Form                    *Form          `bun:"rel:belongs-to,join:form_id=id"`
	Title                   string         `bun:"title,notnull"`
	Type                    FORM_TYPE_ENUM `bun:"type,notnull"`
	Points                  int            `bun:"points,notnull"`
	IsRequired              bool           `bun:"is_required,default:false"`
	IsPartialMarkingEnabled bool           `bun:"is_partial_marking_enabled,default:false"`
	CreatedAt               time.Time      `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt               time.Time      `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Option struct {
	bun.BaseModel `bun:"table:options"`

	Id         int64     `bun:"id,pk,autoincrement"`
	QuestionID int64     `bun:"question_id"`
	Question   *Question `bun:"rel:belongs-to,join:question_id=id"`
	Value      string    `bun:"value,notnull"`
	CreatedAt  time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
