package models

import (
	"time"

	"github.com/uptrace/bun"
)

type FORM_STATUS_TYPE string

const (
	DRAFT     FORM_STATUS_TYPE = "DRAFT"
	PUBLISHED FORM_STATUS_TYPE = "PUBLISHED"
)

type FormContentKeyType string

const FORM_CONTENT_KEY FormContentKeyType = "blocks"

type FormContent map[FormContentKeyType][]Block

type Form struct {
	bun.BaseModel `bun:"table:forms.form"`

	Id          int64            `bun:"id,pk,autoincrement" json:"id"`
	Content     FormContent      `bun:"content" json:"content"`
	CreatedById int64            `bun:"created_by_id" json:"created_by_id"`
	CreatedBy   *User            `bun:"rel:belongs-to,join:created_by_id=id" json:"created_by"`
	UpdatedById *int64           `bun:"updated_by_id,default:null" json:"updated_by_id"`
	UpdatedBy   *User            `bun:"rel:belongs-to,join:updated_by_id=id" json:"updated_by"`
	OwnerId     int64            `bun:"owner_id" json:"owner_id"`
	Owner       *User            `bun:"rel:belongs-to,join:owner_id=id" json:"owner"`
	Status      FORM_STATUS_TYPE `bun:"status,default:'DRAFT'" json:"status"`
	CreatedAt   time.Time        `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time        `bun:"updated_at,default:null" json:"updated_at"`
}

type FormMetaData struct {
	bun.BaseModel `bun:"table:forms.metadata"`

	Id                               int64     `bun:"id,pk,autoincrement" json:"id"`
	FormId                           int64     `bun:"form_id" json:"form_id"`
	Form                             *Form     `bun:"rel:belongs-to,join:form_id=id" json:"form"`
	IsDeleted                        bool      `bun:"is_deleted,default:false" json:"is_deleted"`
	AccepctingResponses              bool      `bun:"accepting_responses,default:false" json:"accepting_responses"`
	AllowGuestResponses              bool      `bun:"allow_guest_responses,default:true" json:"allow_guest_responses"`
	AllowMultipleRepsonses           bool      `bun:"allow_multiple_responses,default:false" json:"allow_multiple_responses"`
	SendConfirmationEmailToRespondee bool      `bun:"send_confirmation_email_to_respondee,default:false" json:"send_confirmation_email_to_respondee"`
	SendSubmissionEmailToOwner       bool      `bun:"send_submission_email_to_owner,default:false" json:"send_submission_email_to_owner"`
	ValidTill                        time.Time `bun:"valid_till" json:"valid_till"`
	UpdatedById                      *int64    `bun:"updated_by_id,default:null" json:"updated_by_id"`
	UpdatedBy                        *User     `bun:"rel:belongs-to,join:updated_by_id=id" json:"updated_by"`
	// TODO invite code
	// TODO remove created by
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,default:null" json:"updated_at"`
}

type FormResponse struct {
	bun.BaseModel `bun:"table:form.responses"`

	Id           int64       `bun:"id,pk,autoincrement" json:"id" `
	ResponseByID int64       `bun:"response_by_id" json:"response_by_id"`
	ResponseBy   *User       `bun:"rel:belongs-to,join:response_by_id=id" json:"response_by"`
	Content      FormContent `bun:"content" json:"content"`
	FormID       int64       `bun:"form_id" json:"form_id"`
	Form         *Form       `bun:"rel:belongs-to,join:form_id=id" json:"form"`
	CreatedAt    time.Time   `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time   `bun:"updated_at,default:null" json:"updated_at"`
}

type Block struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Content string      `json:"content"`
	GroupId string      `json:"groupID"`
	Order   int         `json:"order"`
	Meta    interface{} `json:"meta"`
}
