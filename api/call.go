package api

import (
	dataUtil "a2billing-go-api/common/data"
	"a2billing-go-api/common/response"
	mdw "a2billing-go-api/middleware/auth"
	"a2billing-go-api/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CallHandler struct {
	CallService service.CallService
}

func NewCallHandler(r *gin.Engine, callService service.CallService) {
	handler := &CallHandler{
		CallService: callService,
	}
	Group := r.Group("v1/call")
	{
		Group.GET("/log", mdw.AuthMiddleware(), handler.GetCallLogs)
	}
}

func (handler *CallHandler) GetCallLogs(c *gin.Context) {
	id, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	startTime := time.Now()
	endTime := time.Now()
	if startDate != "" {
		startTime = dataUtil.ParseFromStringToTime(startDate)
		if startTime.IsZero() {
			c.JSON(response.BadRequestMsg("end_date is invalid"))
			return
		}
	} else {
		startDate = dataUtil.TimeToStringLayout(startTime, "2006-01-02") + " 00:00:00"
	}
	if endDate != "" {
		endTime = dataUtil.ParseFromStringToTime(endDate)
		if endTime.IsZero() {
			c.JSON(response.BadRequestMsg("end_date is invalid"))
			return
		}
	} else {
		endDate = dataUtil.TimeToStringLayout(startTime, "2006-01-02") + " 23:59:59"
	}
	if startTime.After(endTime) {
		c.JSON(response.BadRequestMsg("start_date must be after end_date"))
		return
	}
	if endTime.Sub(startTime) >= (time.Hour * 24 * 31) {
		c.JSON(response.BadRequestMsg("date range must be in 31 days"))
		return
	}
	source := c.Query("source")
	customerId := c.Query("customer_id")
	limit := dataUtil.ParseLimit(c.Query("limit"))
	offset := dataUtil.ParseOffset(c.Query("offset"))
	code, result := handler.CallService.GetCallLogs(&gin.Context{}, id, customerId, source, startDate, endDate, limit, offset)
	c.JSON(code, result)

}
