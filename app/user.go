package app

import (
	"errors"
	"fmt"

	"hoiLightningTalk/domain"
	"hoiLightningTalk/db"

)

func SignIn(id string, username string,callbackText string){
	ur := db.NewUserRepository()
	ur.SaveUser(domain.User{Id:username,Username:username,SlackId:id,CallbackText:callbackText})
}

func GetUserUrl(userId string) (string,error){
	ur := db.NewUserRepository()
	user, err := ur.GetUser(userId)


	if err != nil{
		return "",errors.New(fmt.Sprintf("%v user doesn't exists",userId))
	}

	if user.CallbackText == ""{
		return "",nil
	}
	callbackText := fmt.Sprintf("Callback from <@%v>: %v",user.SlackId,user.CallbackText)
	return callbackText,nil;

}