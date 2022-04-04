package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"errors"

	"database/sql"
)

type CardRepository struct {
}

func NewCardRepository() CardRepository {
	repo := CardRepository{}
	return repo
}

var CardRepo CardRepository

func (repo *CardRepository) GetCardsOfAgent(ctx context.Context, agentId string, limit, offset int) (*[]model.CardInfo, int, error) {
	cards := new([]model.CardInfo)
	var count int
	resp := repository.BillingSqlClient.GetDB().NewSelect().Model(cards).
		TableExpr("cc_card AS c").
		ColumnExpr("c.id, c.creationdate, c.firstusedate, c.expirationdate, c.enableexpire, c.expiredays, c.username, c.useralias, c.credit, c.activated, c.status, c.lastuse, c.creditlimit, c.id_group, c.tariff").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId).Order("c.id ASC").
		Limit(limit).Offset(offset)
	count, err := resp.ScanAndCount(ctx)
	if err == sql.ErrNoRows {
		return cards, 0, nil
	} else if err != nil {
		return nil, count, err

	}
	return cards, count, nil
}

func (repo *CardRepository) UpdateCCardCreditOfAgent(ctx context.Context, id string, credit float64) (bool, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewUpdate().Model(&model.CardInfo{}).
		Set("credit = ?", credit).
		Where("id = ?", id).
		Exec(ctx)

	if affected, _ := resp.RowsAffected(); affected == -1 {
		return false, errors.New("update failed")
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *CardRepository) UpdateCCardStatusOfAgent(ctx context.Context, id string, status int) (bool, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewUpdate().Model((*model.Card)(nil)).
		Set("status = ?", status).
		Where("id = ?", id).
		Exec(ctx)
	if affected, _ := resp.RowsAffected(); affected == -1 {
		return false, errors.New("update failed")

	} else if err != nil {
		return false, err

	}
	return true, nil
}

func (repo *CardRepository) GetCardOfAgentById(ctx context.Context, agentId, id string) (*model.CardInfo, error) {
	card := new(model.CardInfo)
	err := repository.BillingSqlClient.GetDB().NewSelect().Model(card).
		TableExpr("cc_card AS c").
		ColumnExpr("c.id, c.creationdate, c.firstusedate, c.expirationdate, c.enableexpire, c.expiredays, c.username, c.useralias, c.credit, c.activated, c.status, c.lastuse, c.creditlimit, c.id_group, c.tariff").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Join("JOIN cc_callerid ccid ON ccid.id_cc_card = c.id").
		Where("cg.id_agent = ?", agentId).
		Where("c.id = ? OR c.username = ? OR ccid.cid = ?", id, id, id).
		Limit(1).
		Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return card, nil
}
func (repo *CardRepository) CreateCard(ctx context.Context, card model.Card) (model.Card, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(&card).Exec(ctx)
	if affected, _ := resp.RowsAffected(); affected == -1 {
		return card, errors.New("create card failed")

	} else if err != nil {
		return card, err
	}

	return card, nil
}
func (repo *CardRepository) CreateCardTransaction(ctx context.Context, card model.Card) (model.Card, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(&card).Exec(ctx)

	if affected, _ := resp.RowsAffected(); affected == -1 {
		return card, errors.New("create card failed")

	} else if err != nil {
		return card, err
	}

	return card, nil
}
