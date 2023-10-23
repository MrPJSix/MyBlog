package repository

import (
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type INotificationRepo interface {
	GetUnReadListByReciver(receiverID uint, pageSize, offset int) ([]model.Notification, int64, int)
	GetReadListByReciver(receiverID uint, pageSize, offset int) ([]model.Notification, int64, int)
	MarkAsRead(notificationIDs []uint) int
	DeletReadList(notificationIDs []uint) int
}

type NotificationRepo struct {
}

func NewNotificationRepo() *NotificationRepo {
	return &NotificationRepo{}
}

/* ====================================== */

func (nr *NotificationRepo) GetUnReadListByReciver(receiverID uint, pageSize, offset int) ([]model.Notification, int64, int) {
	var notifications []model.Notification
	var total int64
	err := db.Model(&model.Notification{}).
		Preload("Sender").Preload("Comment").Preload("Reply").
		Where("receiver_id = ? AND is_read = 0", receiverID).
		Order("create_at DESC").
		Limit(pageSize).Offset(offset).
		Count(&total).
		Find(&notifications).Error
	if err != nil {
		log.Println("获取未读消息失败！", err)
		return nil, 0, errmsg.ERROR
	}
	return notifications, total, errmsg.SUCCESS
}

func (nr *NotificationRepo) GetReadListByReciver(receiverID uint, pageSize, offset int) ([]model.Notification, int64, int) {
	var notifications []model.Notification
	var total int64
	err := db.Model(&model.Notification{}).
		Preload("Sender").Preload("Comment").Preload("Reply").
		Where("receiver_id = ? AND is_read = 1", receiverID).
		Order("create_at DESC").
		Limit(pageSize).Offset(offset).
		Count(&total).
		Find(&notifications).Error
	if err != nil {
		log.Println("获取已读消息失败！", err)
		return nil, 0, errmsg.ERROR
	}
	return notifications, total, errmsg.SUCCESS
}

func (nr *NotificationRepo) MarkAsRead(notificationIDs []uint) int {
	err := db.Model(model.Notification{}).
		Where("id IN ?", notificationIDs).
		Updates(map[string]interface{}{"is_read": 1}).Error
	if err != nil {
		log.Println("标记为已读失败！", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (nr *NotificationRepo) DeletList(notificationIDs []uint) int {
	err := db.Delete(model.Notification{}, notificationIDs).
		Where("is_read = 1").Error
	if err != nil {
		log.Println("批量删除已读消息失败！", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
