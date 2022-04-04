package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"errors"
)

type IaxBuddiesRepository struct {
}

var IaxBuddiesRepo IaxBuddiesRepository

func NewIaxBuddiesRepository() IaxBuddiesRepository {
	repo := IaxBuddiesRepository{}
	return repo
}
func (repo *IaxBuddiesRepository) CreateIaxBuddies(ctx context.Context, iaxBuddies model.IaxBuddies) (model.IaxBuddies, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(&iaxBuddies).Exec(ctx)
	if affected, _ := resp.RowsAffected(); affected == -1 {
		return iaxBuddies, errors.New("create iaxBuddies failed")

	} else if err != nil {
		return iaxBuddies, err

	}
	return iaxBuddies, nil
}
func (repo *IaxBuddiesRepository) CreateIaxBuddiesTransaction(ctx context.Context, iaxBuddies model.IaxBuddies) (model.IaxBuddies, error) {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(&iaxBuddies).Exec(ctx)
	if affected, _ := resp.RowsAffected(); affected == -1 {
		return iaxBuddies, errors.New("create iaxBuddies failed")

	} else if err != nil {
		return iaxBuddies, err

	}
	return iaxBuddies, nil
}
