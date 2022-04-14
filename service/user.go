package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/response"
	"a2billing-go-api/middleware/auth/goauth"
	"a2billing-go-api/repository/db"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AgentService struct {
}

func NewAgentService() AgentService {
	return AgentService{}
}
func (service *AgentService) GenerateTokenByApiKey(ctx context.Context, apiKey string, isRefresh bool) (int, interface{}) {
	Agent, err := db.AgentRepo.GetAgentByApiKey(ctx, apiKey)
	if err != nil {
		log.Error(err)
		return response.NotFound()
	}
	if Agent == nil {
		log.Error(err)
		return response.NotFound()
	}

	clientAuth := goauth.AuthClient{
		ClientId: apiKey,
		UserId:   fmt.Sprintf("%d", Agent.ID),
		UserData: map[string]string{},
	}
	client, err := goauth.GoAuthClient.ClientCredential(clientAuth, false)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg(err.Error())
	}
	token := gin.H{
		"client_id":     client.ClientId,
		"user_id":       client.UserId,
		"token":         client.Token,
		"refresh_token": client.RefreshToken,
		"expired_in":    client.ExpiredIn,
		"token_type":    client.TokenType,
	}
	return response.NewOKResponse(token)
}
