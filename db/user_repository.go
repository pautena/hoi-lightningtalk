package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hoiLightningTalk/domain"
)

/*
Mongodb models
*/

type UserRepository struct {
	Collection *mongo.Collection 
}



func NewUserRepository() UserRepository {
	return UserRepository {Collection:getDatabase().Collection("users")};
}

func (ur UserRepository) SaveUser(p domain.User) {
	
	filter := bson.D{{"_id", p.Id}}
	options := options.Replace().SetUpsert(true)
	insertResult, err := ur.Collection.ReplaceOne(context.TODO(),filter, p,options)

	if err != nil {
		fmt.Sprintln("Save user error: ",err)
	}

	fmt.Println("User had been inserted: ", insertResult)
}
