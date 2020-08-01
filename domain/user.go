package domain

type User struct {
	ID           string `bson:"_id"`
	Username     string `bson:"username"`
	AppID        string `bson:"AppID"`
	CallbackText string `bson:"callbackText"`
}
