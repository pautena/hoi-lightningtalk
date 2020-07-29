package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"hoiLightningTalk/domain"
)

type MessageRepository struct {
	Collection *mongo.Collection 
}



func NewMessageRepository() MessageRepository {
	return MessageRepository {Collection:getDatabase().Collection("messages")};
}

func (mr MessageRepository) SaveMessage(s domain.SlackResponse) {

	insertResult, err := mr.Collection.InsertOne(context.TODO(), s)

	if err != nil {
		fmt.Sprintln("Save message error: ",err)
	}

	fmt.Println("Message had been inserted: ", insertResult)
}
