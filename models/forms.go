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

	Id          int64            `bun:"id,pk,autoincrement"`
	Content     FormContent      `bun:"content"`
	CreatedById int64            `bun:"created_by_id"`
	CreatedBy   *User            `bun:"rel:belongs-to,join:created_by_id=id"`
	OwnerId     int64            `bun:"owner_id"`
	Owner       *User            `bun:"rel:belongs-to,join:owner_id=id"`
	Status      FORM_STATUS_TYPE `bun:"status,default:'DRAFT'"`
	CreatedAt   time.Time        `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time        `bun:"updated_at,default:null"`
}

type FormMetaData struct {
	bun.BaseModel `bun:"table:forms.metadata"`

	Id                               int64     `bun:"id,pk,autoincrement"`
	FormId                           int64     `bun:"form_id"`
	Form                             *Form     `bun:"rel:belongs-to,join:form_id=id"`
	IsDeleted                        bool      `bun:"is_deleted,default:false"`
	AccepctingResponses              bool      `bun:"accepting_responses,default:false"`
	AllowGuestResponses              bool      `bun:"allow_guest_responses,default:true"`
	AllowMultipleRepsonses           bool      `bun:"allow_multiple_responses,default:false"`
	SendConfirmationEmailToRespondee bool      `bun:"send_confirmation_email_to_respondee,default:false"`
	SendSubmissionEmailToOwner       bool      `bun:"send_submission_email_to_owner,default:false"`
	ValidTill                        time.Time `bun:"valid_till"`
	// invite code
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:null"`
}

type FormResponse struct {
	bun.BaseModel `bun:"table:form.responses"`

	Id           int64       `bun:"id,pk,autoincrement"`
	ResponseByID int64       `bun:"response_by_id"`
	ResponseBy   *User       `bun:"rel:belongs-to,join:response_by_id=id"`
	Content      FormContent `bun:"content"`
	FormID       int64       `bun:"form_id"`
	Form         *Form       `bun:"rel:belongs-to,join:form_id=id"`
	CreatedAt    time.Time   `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time   `bun:"updated_at,default:null"`
}

type Block struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Content string      `json:"content"`
	GroupId string      `json:"groupID"`
	Order   int         `json:"order"`
	Meta    interface{} `json:"meta"`
}
