
package domain

type User struct{
	Id string `bson:"_id"`
	Username string `bson:"username"`
	SlackId string `bson:"SlackId"`
}
