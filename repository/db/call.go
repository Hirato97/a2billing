package db

import (
	"a2billing-go-api/repository"
	"context"
	"database/sql"
)

type CallRepository struct {
}

func NewCallRepository() CallRepository {
	repo := CallRepository{}
	return repo
}

var CallRepo CallRepository

func (repo *CallRepository) GetCallLogs(ctx context.Context, agentId, cardId, source string, fromDate, toDate string, limit, offset int) (interface{}, int, error) {
	result := []map[string]interface{}{}
	var count int
	field := "ca.id, ca.uniqueid as uniqueid, ca.starttime as start_time, ca.stoptime as end_time, ca.sessiontime as duration, ca.sessionbill as amount, ca.src as source, ca.calledstation as destination"
	query := repository.BillingSqlClient.GetDB().NewSelect().
		TableExpr("cc_call AS ca").
		Join("JOIN cc_card c ON c.id = ca.card_id").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId).
		ColumnExpr(field)
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
	count, err := query.ScanAndCount(ctx)
	if err == sql.ErrNoRows {
		return result, 0, nil

	} else if err != nil {
		return nil, 0, err

	}
	return result, count, nil
}
