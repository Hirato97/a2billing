package model

import (
	"github.com/uptrace/bun"
)

type TariffGroup struct {
	bun.BaseModel `bun:"cc_tariffgroup"`
	ID            int    `bun:"id,pk" json:"id"`                        //
	Iduser        int    `bun:"iduser" json:"iduser"`                   //
	Loglevel      string `bun:"tariffgroupname" json:"tariffgroupname"` //
}
