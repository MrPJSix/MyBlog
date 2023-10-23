package v1

import (
	"github.com/gin-gonic/gin"
	dto2 "myblog.backend/dto/request"
	dto "myblog.backend/dto/response"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

/* ====================================== */

type INotificationController interface {
	GetUnReadNotifsByReciver(c *gin.Context)
	GetReadNotifsByReciver(c *gin.Context)
	MarkAsReadNotifs(c *gin.Context)
	DeletReadNotifs(c *gin.Context)
}

type NotificationController struct {
	notifService *service.NotificationService
}

func NewNotificationController() *NotificationController {
	notifService := service.NewNotificationService()
	return &NotificationController{notifService}
}

/* ====================================== */

func (nc *NotificationController) GetUnReadNotifsByReciver(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	receiverID := c.MustGet("user_id").(uint)

	notifs, total, code := nc.notifService.GetUnReadNotifsByReciver(receiverID, pageSize, pageNum)

	responseData := dto.NotificationSliceToResponse(notifs)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (nc *NotificationController) GetReadNotifsByReciver(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	receiverID := c.MustGet("user_id").(uint)

	notifs, total, code := nc.notifService.GetReadNotifsByReciver(receiverID, pageSize, pageNum)

	responseData := dto.NotificationSliceToResponse(notifs)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (nc *NotificationController) MarkAsReadNotifs(c *gin.Context) {
	var notifReq dto2.NotificationRequest
	var code int
	err := c.ShouldBindJSON(&notifReq)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = nc.notifService.MarkAsReadNotifs(notifReq.NotificationIDs)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (nc *NotificationController) DeletReadNotifs(c *gin.Context) {
	var notifReq dto2.NotificationRequest
	var code int
	err := c.ShouldBindJSON(&notifReq)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = nc.notifService.DeletReadNotifs(notifReq.NotificationIDs)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
