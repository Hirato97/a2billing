package api

import (
	"a2billing-go-api/common/response"
	"a2billing-go-api/common/validator"
	mdw "a2billing-go-api/internal/middleware"
	"a2billing-go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CallerIdHandler struct {
	CallerIdService service.CallerIdService
}

func NewCallerIdHandler(r *gin.Engine, callService service.CallerIdService) {
	handler := &CallerIdHandler{
		CallerIdService: callService,
	}
	Group := r.Group("v1/callerid")
	{
		Group.POST("", mdw.AuthMiddleware(), handler.AddCallerIdToCard)
		Group.PUT(":id", mdw.AuthMiddleware(), handler.UpdateCallerIdToCard)
	}
}

func (handler *CallerIdHandler) AddCallerIdToCard(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("add-callerid.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	cardId, _ := jsonBody["user_id"].(string)
	callerId, _ := jsonBody["cid"].(string)
	code, result := handler.CallerIdService.AddCallerIdToCard(c, userId.(string), cardId, callerId)
	c.JSON(code, result)
}

func (handler *CallerIdHandler) UpdateCallerIdToCard(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	id := c.Param("id")
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("update-callerid.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	cardId, _ := jsonBody["user_id"].(string)
	if cardId == "" {
		c.JSON(response.BadRequestMsg("user_id is required"))
		return
	}
	code, result := handler.CallerIdService.UpdateCallerIdToCard(c, userId.(string), id, cardId)
	c.JSON(code, result)
}
