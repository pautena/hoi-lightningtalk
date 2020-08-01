package app

import (
	"hoiLightningTalk/domain"
)

type UserRepository interface {
	SaveUser(user domain.User)
	GetUser(uID string) (domain.User, error)
}
