package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"errors"
)

type SystemLogRepository struct {
}

func NewSystemLogRepository() SystemLogRepository {
	repo := SystemLogRepository{}
	return repo
}

var SystemLogRepo SystemLogRepository

func (repo *SystemLogRepository) CreateLog(ctx context.Context, systemLog *model.SystemLog) error {
	resp, err := repository.BillingSqlClient.GetDB().NewInsert().Model(systemLog).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected == -1 {
		return errors.New("create failed")

	}

	return nil
}
