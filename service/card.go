package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository/db"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type CardService struct {
}

func NewCardService() CardService {
	return CardService{}
}

func (service *CardService) GetCardsOfAgent(ctx context.Context, agentId string, limit, offset int) (int, interface{}) {
	cards, total, err := db.CardRepo.GetCardsOfAgent(ctx, agentId, limit, offset)
	if err != nil {
		log.Error(err)
	}
	return response.NewBaseResponsePagination(cards, limit, offset, int(total))
}

func (service *CardService) GetCardsOfAgentById(ctx context.Context, agentId string, id string) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, id)
	if err != nil {
		log.Error(err)
	}
	if card == nil {
		return response.NotFound()
	}
	return response.NewResponse(http.StatusOK, card)
}

func (service *CardService) UpdateCardCreditOfAgent(ctx context.Context, agentId, id string, credit float64) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, id)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if card == nil {
		return response.NotFound()
	}
	isUpdated, err := db.CardRepo.UpdateCCardCreditOfAgent(ctx, fmt.Sprintf("%v", card.ID), credit)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		db.SystemLogRepo.CreateLog(ctx, model.SystemLog{
			Iduser:       userId,
			Loglevel:     1,
			Action:       "API EXECUTE",
			Description:  "Update card " + id + " credit : " + fmt.Sprintf("%v", credit),
			Tablename:    "cc_card",
			Creationdate: time.Now(),
			Ipaddress:    "API",
			Pagename:     "Customer API",
			Agent:        int64(userId),
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"credit":  credit,
	})
}
func (service *CardService) AddCardCreditOfAgent(ctx context.Context, agentId, id string, credit float64) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, id)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if card == nil {
		return response.NotFound()
	}

	credit = card.Credit + credit
	isUpdated, err := db.CardRepo.UpdateCCardCreditOfAgent(ctx, fmt.Sprintf("%v", card.ID), credit)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		db.SystemLogRepo.CreateLog(ctx, model.SystemLog{
			Iduser:      userId,
			Loglevel:    1,
			Action:      "API EXECUTE",
			Description: "Update card " + id + " credit : " + fmt.Sprintf("%v", credit),

			Tablename:    "cc_card",
			Creationdate: time.Now(),
			Ipaddress:    "API",
			Pagename:     "Customer API",
			Agent:        int64(userId),
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"credit":  credit,
	})
}

func (service *CardService) UpdateCardStatusOfAgent(ctx context.Context, agentId, id string, status int) (int, interface{}) {
	card, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, id)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if card == nil {
		return response.NotFound()
	}

	isUpdated, err := db.CardRepo.UpdateCCardStatusOfAgent(ctx, fmt.Sprintf("%v", card.ID), status)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		db.SystemLogRepo.CreateLog(ctx, model.SystemLog{
			Iduser:       userId,
			Loglevel:     1,
			Action:       "API EXECUTE",
			Description:  "Update card " + id + " status : " + fmt.Sprintf("%v", status),
			Tablename:    "cc_card",
			Creationdate: time.Now(),
			Ipaddress:    "API",
			Pagename:     "Customer API",
			Agent:        int64(userId),
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"status":  status,
	})
}

func (service *CardService) CreateCardAndSip(ctx context.Context, agentId string, card model.Card, cid string) (int, interface{}) {
	cardRes, err := db.CardRepo.GetCardOfAgentById(ctx, agentId, card.Username)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if cardRes != nil {
		return response.BadRequestMsg("username is already exists")
	}
	if groupId, err := db.AgentRepo.GetGroupIdById(ctx, agentId); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if groupId < 1 {
		return response.BadRequestMsg("please check group configuration")
	} else {
		card.IDGroup = groupId
	}
	if callerId, err := db.CallerIdRepo.GetCallerIdByCid(ctx, agentId, cid); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerId != nil {
		return response.BadRequestMsg("caller_id is already exists")
	}
	if tariffGroup, err := db.TariffGroupRepo.GetTariffGroupById(ctx, card.Tariff); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if tariffGroup == nil {
		return response.BadRequestMsg("call_plan is not exists")
	}

	err = db.CardRepo.CreateCardTransaction(ctx, &card)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	}
	result := map[string]interface{}{
		"message":   "successfully",
		"id":        int(card.ID),
		"username":  card.Username,
		"password":  card.Uipass,
		"status":    card.Status,
		"call_plan": card.Tariff,
	}
	if callerId, err := db.CallerIdRepo.CreateCallerIdTransaction(ctx, model.CallerId{Cid: cid, IDCcCard: card.ID, Activated: "t"}); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("create customer invalid")
	} else {
		result["cid"] = callerId.Cid
	}
	if card.SipBuddy == 1 {
		sipBuddies := model.SipBuddies{
			IDCcCard:       int(card.ID),
			Name:           card.Username,
			Accountcode:    card.Username,
			Regexten:       card.Username,
			Amaflags:       "billing",
			Canreinvite:    "YES",
			Context:        "a2billing",
			Dtmfmode:       "RFC2833",
			Host:           "dynamic",
			Nat:            "no",
			Qualify:        "no",
			Secret:         card.Uipass,
			Type:           "friend",
			Username:       card.Username,
			Disallow:       "ALL",
			Allow:          "ulaw,alaw,gsm,g729",
			Regseconds:     0,
			Cancallforward: "yes",
			Rtpkeepalive:   "0",
		}
		sipBuddies, err := db.SipBuddiesRepo.CreateSipBuddiesTransaction(ctx, sipBuddies)
		if err != nil {
			log.Error(err)
			return response.ServiceUnavailableMsg("create customer invalid")
		}
		result["sip"] = "created"
	}
	if card.IaxBuddy == 1 {
		iaxBuddies := model.IaxBuddies{
			IDCcCard:    int(card.ID),
			Name:        card.Username,
			Accountcode: card.Username,
			Regexten:    card.Username,
			Amaflags:    "billing",
			Context:     "a2billing",
			Host:        "dynamic",
			Qualify:     "no",
			Secret:      card.Uipass,
			Type:        "friend",
			Username:    card.Username,
			Disallow:    "",
			Allow:       "ulaw,alaw,gsm,g729",
			Regseconds:  0,
			Trunk:       "no",
		}
		iaxBuddies, err := db.IaxBuddiesRepo.CreateIaxBuddiesTransaction(ctx, iaxBuddies)
		if err != nil {
			log.Error(err)
			return response.ServiceUnavailableMsg("create customer invalid")
		}
		result["iax"] = "created"
	}

	return response.NewResponse(http.StatusOK, result)
}
