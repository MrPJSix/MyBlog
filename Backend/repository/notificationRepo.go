package repository

import (
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type INotificationRepo interface {
	GetUnReadListByReciver(receiverID uint) ([]model.Notification, int)
	GetReadListByReciver(receiverID uint) ([]model.Notification, int)
	MarkAsRead(notificationIDs []uint) int
}

type NotificationRepo struct {
}

func NewNotifiCationRepo() *NotificationRepo {
	return &NotificationRepo{}
}

/* ====================================== */

func (nr *NotificationRepo) GetUnReadListByReciver(receiverID uint) ([]model.Notification, int) {
	var notifications []model.Notification
	err := db.Where("receiver_id = ? AND is_read = 0", receiverID).Find(&notifications).Error
	if err != nil {
		log.Println("获取未读消息失败！", err)
		return nil, errmsg.ERROR
	}
	return notifications, errmsg.SUCCESS
}

func (nr *NotificationRepo) GetReadListByReciver(receiverID uint) ([]model.Notification, int) {
	var notifications []model.Notification
	err := db.Where("receiver_id = ? AND is_read = 1", receiverID).Find(&notifications).Error
	if err != nil {
		log.Println("获取已读消息失败！", err)
		return nil, errmsg.ERROR
	}
	return notifications, errmsg.SUCCESS
}

func (nr *NotificationRepo) MarkAsRead(notificationIDs []uint) int {
	err := db.Model(model.Notification{}).Where("id IN ?", notificationIDs).Updates(map[string]interface{}{"read": 1})
	if err != nil {
		log.Println("标记为已读失败！", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
