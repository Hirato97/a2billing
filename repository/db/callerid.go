package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"errors"

	"database/sql"
)

type CallerIdRepository struct {
}

var CallerIdRepo CallerIdRepository

func NewCallerIdRepository() CallerIdRepository {
	repo := CallerIdRepository{}
	return repo
}

func (repo *CallerIdRepository) GetCallerIdByCid(ctx context.Context, agentId, cid string) (*model.CallerId, error) {
	callerId := new(model.CallerId)
	err := repository.BillingSqlClient.GetDB().NewSelect().Model(callerId).
		TableExpr("cc_callerid AS ccid").
		// TableExpr("cc_card AS c").
		ColumnExpr("ccid.id, ccid.cid, ccid.id_cc_card, ccid.activated").
		Join("INNER JOIN cc_card c ON c.id = ccid.id_cc_card").
		Join("INNER JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId).
		Where("ccid.cid = ?", cid).
		Limit(1).Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return callerId, nil
}

func (repo *CallerIdRepository) CreateCallerId(ctx context.Context, callerId *model.CallerId) error {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(callerId).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return errors.New("create callerId failed")
	}
	return nil
}
func (repo *CallerIdRepository) CreateCallerIdTransaction(ctx context.Context, callerId *model.CallerId) error {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(callerId).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return errors.New("create callerId failed")
	}
	return nil
}
func (repo *CallerIdRepository) UpdateCallerIdToCard(ctx context.Context, id int, cardId int) (bool, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewUpdate().Model(&model.CallerId{}).
		Set("id_cc_card = ?", cardId).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return false, err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return false, errors.New("update failed")
	}
	return true, nil
}
