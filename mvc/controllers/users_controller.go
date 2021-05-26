package controllers

import (
	"encoding/json"
	"golang-microservices-course/mvc/services"
	"golang-microservices-course/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userIdParam, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil{
		// just return the bad request to the client
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return;
	}

	user, apiErr := services.GetUser(userIdParam)
	if apiErr != nil{
		// Handle error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}