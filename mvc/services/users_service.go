package services

import (
	"golang-microservices-course/mvc/domain"
	"golang-microservices-course/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}
