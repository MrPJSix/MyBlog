package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
)

/* ====================================== */

type INotificationService interface {
	GetUnReadNotifsByReciver(receiverID uint, pageSize, pageNum int) ([]model.Notification, int64, int)
	GetReadNotifsByReciver(receiverID uint, pageSize, pageNum int) ([]model.Notification, int64, int)
	MarkAsReadNotifs(notificationIDs []uint) int
	DeletReadNotifs(notificationIDs []uint) int
}

type NotificationService struct {
	notiRepo *repository.NotificationRepo
}

func NewNotificationService() *NotificationService {
	notiRepo := repository.NewNotificationRepo()
	return &NotificationService{notiRepo: notiRepo}
}

/* ====================================== */

func (ns *NotificationService) GetUnReadNotifsByReciver(receiverID uint, pageSize, pageNum int) ([]model.Notification, int64, int) {
	var offset int
	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}
	if pageNum <= 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	return ns.notiRepo.GetUnReadListByReciver(receiverID, pageSize, offset)
}

func (ns *NotificationService) GetReadNotifsByReciver(receiverID uint, pageSize, pageNum int) ([]model.Notification, int64, int) {
	var offset int
	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}
	if pageNum <= 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	return ns.notiRepo.GetReadListByReciver(receiverID, pageSize, offset)
}

func (ns *NotificationService) MarkAsReadNotifs(notificationIDs []uint) int {
	return ns.notiRepo.MarkAsRead(notificationIDs)
}

func (ns *NotificationService) DeletReadNotifs(notificationIDs []uint) int {
	return ns.notiRepo.DeletList(notificationIDs)
}
