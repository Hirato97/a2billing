package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"errors"
)

type SipBuddiesRepository struct {
}

var SipBuddiesRepo SipBuddiesRepository

func NewSipBuddiesRepository() SipBuddiesRepository {
	repo := SipBuddiesRepository{}
	return repo
}
func (repo *SipBuddiesRepository) CreateSipBuddies(ctx context.Context, sipBuddies *model.SipBuddies) error {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(sipBuddies).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return errors.New("create sipBuddies failed")
	}
	return nil
}
func (repo *SipBuddiesRepository) CreateSipBuddiesTransaction(ctx context.Context, sipBuddies *model.SipBuddies) error {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(sipBuddies).Exec(ctx)

	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return errors.New("create sipBuddies failed")
	}
	return nil
}
