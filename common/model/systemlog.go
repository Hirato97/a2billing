package model

import (
	"time"

	"github.com/uptrace/bun"
)

type SystemLog struct {
	bun.BaseModel `bun:"cc_system_log"`
	ID            int       `bun:"id,pk" json:"id"`                         //
	Iduser        int       `bun:"iduser" json:"iduser"`                    //
	Loglevel      int       `bun:"loglevel" json:"loglevel"`                //
	Action        string    `bun:"action" json:"action"`                    //
	Description   string    `bun:"description,nullzero" json:"description"` //
	Data          []byte    `bun:"data" json:"data"`                        //
	Tablename     string    `bun:"tablename,nullzero" json:"tablename"`     //
	Pagename      string    `bun:"pagename,nullzero" json:"pagename"`       //
	Ipaddress     string    `bun:"ipaddress,nullzero" json:"ipaddress"`     //
	Creationdate  time.Time `bun:"creationdate" json:"creationdate"`        //
	Agent         int64     `bun:"agent,nullzero" json:"agent"`             //
}

// TableName sets the insert table name for this struct type
