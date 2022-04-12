package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository/db"
	"context"
	"net/http"
	"strconv"
	"time"
)

type CallerIdService struct {
}

func NewCallerIdService() CallerIdService {
	return CallerIdService{}
}

func (service *CallerIdService) AddCallerIdToCard(ctx context.Context, agentId, cardId, cid string) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, cardId)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if card == nil {
		return response.BadRequestMsg("user_id is not exists")
	}

	callerIdRes, err := db.CallerIdRepo.GetCallerIdByCid(ctx, agentId, cid)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerIdRes != nil {
		return response.BadRequestMsg("caller_id is already exists")
	}
	callerId := model.CallerId{Cid: cid, IDCcCard: card.ID, Activated: "t"}
	if err := db.CallerIdRepo.CreateCallerId(ctx, &callerId); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerId.ID < 1 {
		return response.ServiceUnavailableMsg("create customer invalid")
	} else {
		userId, _ := strconv.Atoi(agentId)
		db.SystemLogRepo.CreateLog(ctx, &model.SystemLog{
			Iduser:       userId,
			Loglevel:     1,
			Action:       "API EXECUTE",
			Description:  "Create callerid " + cid + " card_id : " + cardId,
			Tablename:    "cc_callerid",
			Creationdate: time.Now(),
			Ipaddress:    "API",
			Pagename:     "Caller API",
			Agent:        int64(userId),
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"cid":     cid,
		"user_id": cardId,
	})
}

func (service *CallerIdService) UpdateCallerIdToCard(ctx context.Context, agentId, cid, cardId string) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, cardId)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if card == nil {
		return response.BadRequestMsg("user_id is not exists")
	}

	callerId, err := db.CallerIdRepo.GetCallerIdByCid(ctx, agentId, cid)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerId == nil {
		return response.BadRequestMsg("caller_id is not exists")
	}

	isUpdated, err := db.CallerIdRepo.UpdateCallerIdToCard(ctx, int(callerId.ID), int(card.ID))
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("update callerid invalid")
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		db.SystemLogRepo.CreateLog(ctx, &model.SystemLog{
			Iduser:       userId,
			Loglevel:     1,
			Action:       "API EXECUTE",
			Description:  "Update callerid " + cid + " card_id : " + cardId,
			Tablename:    "cc_callerid",
			Creationdate: time.Now(),
			Ipaddress:    "API",
			Pagename:     "Caller API",
			Agent:        int64(userId),
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"cid":     cid,
		"user_id": cardId,
	})
}
