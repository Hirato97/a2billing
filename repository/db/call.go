package db

import (
	"a2billing-go-api/repository"
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type CallRepository struct {
}

func NewCallRepository() CallRepository {
	repo := CallRepository{}
	return repo
}

var CallRepo CallRepository

type CallLog struct {
	bun.BaseModel `bun:"cc_call,alias:ca"`
	ID            int64     `bun:"id,pk" json:"id"`
	Starttime     time.Time `bun:"starttime" json:"start_time"`
	Stoptime      time.Time `bun:"stoptime" json:"end_time"`
	Sessiontime   int64     `bun:"sessiontime,nullzero" json:"duration"`
	Sessionbill   float64   `bun:"sessionbill" json:"amount"`
	Uniqueid      string    `bun:"uniqueid" json:"uniqueid"`
	Calledstation string    `bun:"calledstation" json:"destination"`
	Src           string    `bun:"src" json:"source"`
}

func (repo *CallRepository) GetCallLogs(ctx context.Context, agentId, cardId, source string, fromDate, toDate string, limit, offset int) (*[]CallLog, int, error) {
	result := new([]CallLog)
	var count int
	field := "ca.id, ca.starttime, ca.stoptime, ca.sessiontime, ca.sessionbill, ca.uniqueid, ca.src, ca.calledstation"
	query := repository.BillingSqlClient.GetDB().NewSelect().Model(result).
		ColumnExpr(field).
		Join("JOIN cc_card c ON c.id = ca.card_id").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId)

	if len(cardId) > 0 {
		query = query.Where("ca.card_id = ?", cardId)
	}
	if len(source) > 0 {
		query = query.Where("ca.src = ?", source)
	}
	if fromDate != "" {
		query = query.Where("ca.starttime >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("ca.starttime <= ?", toDate)
	}
	count, err := query.Limit(limit).Offset(offset).ScanAndCount(ctx)
	if err == sql.ErrNoRows {
		return result, 0, nil

	} else if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}
