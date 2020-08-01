package app

import (
	"fmt"

	"hoiLightningTalk/domain"
)

func SignIn(id string, username string, callbackText string, userRepo UserRepository) {
	userRepo.SaveUser(domain.User{ID: username, Username: username, AppID: id, CallbackText: callbackText})
}

func GetUserURL(userID string, userRepo UserRepository) (string, error) {
	user, err := userRepo.GetUser(userID)

	if err != nil {
		return "", fmt.Errorf("%v user doesn't exists", userID)
	}

	if user.CallbackText == "" {
		return "", nil
	}
	callbackText := fmt.Sprintf("Callback from <@%v>: %v", user.AppID, user.CallbackText)
	return callbackText, nil

}
