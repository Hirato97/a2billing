package model

import (
	"github.com/uptrace/bun"
)

type CallerId struct {
	bun.BaseModel `bun:"cc_callerid"`
	ID            int64  `bun:"id,pk" json:"-"`             //
	Cid           string `bun:"cid" json:"cid"`             //
	IDCcCard      int64  `bun:"id_cc_card" json:"-"`        //
	Activated     string `bun:"activated" json:"activated"` //
}

// TableName sets the insert table name for this struct type
