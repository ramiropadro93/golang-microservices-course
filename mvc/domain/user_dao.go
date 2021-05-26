package domain

import (
	"fmt"
	"golang-microservices-course/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "LALA", LastName: "LALA", Email: "LALA@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError){
	user := users[userId]
	if user == nil{
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}
	return user, nil
}