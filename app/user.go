package app

import (
	"errors"
	"fmt"

	"hoiLightningTalk/domain"

)

func SignIn(id string, username string,callbackText string, userRepo UserRepository){
	userRepo.SaveUser(domain.User{Id:username,Username:username,SlackId:id,CallbackText:callbackText})
}

func GetUserUrl(userId string, userRepo UserRepository) (string,error){
	user, err := userRepo.GetUser(userId)


	if err != nil{
		return "",errors.New(fmt.Sprintf("%v user doesn't exists",userId))
	}

	if user.CallbackText == ""{
		return "",nil
	}
	callbackText := fmt.Sprintf("Callback from <@%v>: %v",user.SlackId,user.CallbackText)
	return callbackText,nil;

}