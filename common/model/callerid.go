package model

import (
	"github.com/uptrace/bun"
)

type CallerId struct {
	bun.BaseModel `bun:"cc_callerid"`
	ID            int64  `bun:"id,pk" json:"id"`              //
	Cid           string `bun:"cid" json:"cid"`               //
	IDCcCard      int64  `bun:"id_cc_card" json:"id_cc_card"` //
	Activated     string `bun:"activated" json:"activated"`   //
	// CardInfo      []CardInfo `bun:"rel:has-many" json:"cc_card"`
}

// TableName sets the insert table name for this struct type
