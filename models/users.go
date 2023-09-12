package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id         int64     `bun:"id,pk,autoincrement"`
	Username   string    `bun:"username,notnull"`
	Email      string    `bun:"email,notnull"`
	IsVerified bool      `bun:"is_verified,default:false"`
	Password   string    `bun:"password,notnull"`
	CreatedAt  time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
