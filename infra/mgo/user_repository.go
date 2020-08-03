package mgo

import (
	"context"
	"fmt"

	"hoiLightningTalk/app"
	"hoiLightningTalk/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository() app.UserRepository {
	return UserMongoRepository{Collection: getDatabase().Collection("users")}
}

func (ur UserMongoRepository) SaveUser(p domain.User) {

	filter := bson.D{primitive.E{Key: "_id", Value: p.ID}}
	options := options.Replace().SetUpsert(true)
	insertResult, err := ur.Collection.ReplaceOne(context.TODO(), filter, p, options)

	if err != nil {
		fmt.Sprintln("Save user error: ", err)
	}

	fmt.Println("User had been inserted: ", insertResult)
}

func (ur UserMongoRepository) GetUser(userID string) (domain.User, error) {
	var result domain.User
	filter := bson.D{primitive.E{Key: "_id", Value: userID}}
	err := ur.Collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}
