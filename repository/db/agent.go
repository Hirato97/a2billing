package db

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/repository"
	"context"
	"database/sql"
	"errors"
)

type AgentRepository struct {
}

func NewAgentRepository() AgentRepository {
	repo := AgentRepository{}
	if resp, err := repository.BillingSqlClient.GetDB().
		NewAddColumn().
		Model((*model.Agent)(nil)).
		IfNotExists().
		ColumnExpr("level enum('superadmin','admin','user','agent') NOT NULL DEFAULT 'admin'").
		Exec(context.Background()); err != nil {
		log.Fatal("User", "NewUser", err)
	} else {
		log.Info("User", "NewUser", resp)
	}
	return repo
}

var AgentRepo AgentRepository

func (repo *AgentRepository) GetAgentByApiKey(ctx context.Context, apiKey string) (*model.Agent, error) {
	Agent := new(model.Agent)
	resp := repository.BillingSqlClient.GetDB().NewSelect().Model(Agent).Where("api_key = ?", apiKey).Limit(1)
	err := resp.Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, errors.New("agent not found")

	} else if err != nil {
		return nil, err
	}
	return Agent, nil
}

func (repo *AgentRepository) GetGroupIdById(ctx context.Context, id string) (int, error) {
	groupId := 0
	query := repository.BillingSqlClient.GetDB().NewSelect().
		Table("cc_card_group").
		Column("id").
		Where("id_agent = ?", id)
	err := query.Scan(ctx, &groupId)
	if err == sql.ErrNoRows {
		return 0, errors.New("group id not found")
	} else if err != nil {
		return 0, err
	}
	return groupId, nil
}
