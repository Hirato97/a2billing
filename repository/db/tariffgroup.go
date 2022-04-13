package db

import (
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"database/sql"
)

type TariffGroup struct {
}

var TariffGroupRepo TariffGroup

func NewTariffGroupRepository() TariffGroup {
	repo := TariffGroup{}
	return repo
}

func (repo *TariffGroup) GetTariffGroupById(ctx context.Context, id int64) (*model.TariffGroup, error) {
	tariffGroup := new(model.TariffGroup)
	err := repository.BillingSqlClient.GetDB().NewSelect().Model(tariffGroup).
		TableExpr("cc_tariffgroup AS tg").
		ColumnExpr("tg.id, tg.iduser, tg.tariffgroupname").
		Where("tg.id = ?", id).
		Limit(1).
		Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return tariffGroup, nil
}
