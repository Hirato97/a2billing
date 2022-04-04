package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Agent struct {
	bun.BaseModel `bun:"cc_agent"`
	ID            int64     `bun:"id,pk" json:"id"`
	DateCreation  time.Time `bun:"datecreation" json:"datecreation"`
	Active        string    `bun:"active" json:"active"`
	Login         string    `bun:"login" json:"login"`
	Passwd        string    `bun:"passwd" json:"passwd"`
	ApiKey        string    `json:"api_key" bun:"api_key"`
}
