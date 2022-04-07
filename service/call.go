package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository/db"
	"context"
)

type CallService struct {
}

func NewCallService() CallService {
	return CallService{}
}

func (service *CallService) GetCallLogs(ctx context.Context, agentId, cardId, source, fromDate, toDate string, limit, offset int) (int, interface{}) {
	callLogs, total, err := db.CallRepo.GetCallLogs(ctx, agentId, cardId, source, fromDate, toDate, limit, offset)
	if err != nil {
		log.Error(err)
	}
	return response.NewBaseResponsePagination(callLogs, limit, offset, int(total))
}
