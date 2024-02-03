package models

import (
	"time"

	"github.com/uptrace/bun"
)

type FORM_TYPE_ENUM string

const (
	TEXT  FORM_TYPE_ENUM = "TEXT"
	TITLE FORM_TYPE_ENUM = "TITLE"
	INPUT FORM_TYPE_ENUM = "INPUT"
	RADIO FORM_TYPE_ENUM = "RADIO"
)

type Form struct {
	bun.BaseModel `bun:"table:forms"`

	Id          int64              `bun:"id,pk,autoincrement"`
	CreatedByID int64              `bun:"created_by_id"`
	CreatedBy   *User              `bun:"rel:belongs-to,join:created_by_id=id"`
	CreatedAt   time.Time          `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time          `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	Type        [][]FORM_TYPE_ENUM `bun:"type"`
	Content     [][]string         `bun:"content"`
}

type FormResponse struct {
	bun.BaseModel `bun:"table:form_responses"`

	Id           int64      `bun:"id,pk,autoincrement"`
	ResponseByID int64      `bun:"response_by_id"`
	ResponseBy   *User      `bun:"rel:belongs-to,join:response_by_id=id"`
	Content      [][]string `bun:"content"`
	FormID       int64      `bun:"form_id"`
	Form         *Form      `bun:"rel:belongs-to,join:form_id=id"`
	CreatedAt    time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
