package services

import (
	"github.com/jinzhu/gorm"
	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/utils"
)

type BaseService struct {
}

var Error *utils.HttpError

var Achievement AchievementService = AchievementService{}
var User UserService = UserService{}
var Comment CommentService = CommentService{}
var Response ResponseService = ResponseService{}
var Request RequestService = RequestService{}
var Tag TagService = TagService{}
var Transaction TransactionService = TransactionService{}
var Notification NotificationService = NotificationService{}

func (service BaseService) GetDB() *gorm.DB {
	return api.GetContext().DB
}