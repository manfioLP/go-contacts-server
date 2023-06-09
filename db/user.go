package db

import (
	Models "contacts-server/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(page int64, limit int64) ([]Models.User, error) {
	var users []Models.User

	filter := bson.D{}
	findOptions := options.Find()
	findOptions.SetSkip(page)
	findOptions.SetLimit(limit)

	cur, err := Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user Models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id string) (*Models.User, error) {
	var user Models.User
	err := Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *Models.User) error {
	_, err := Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *Models.User) error {
	_, err := Collection.ReplaceOne(context.TODO(), bson.M{"_id": user.ID}, user)
	if err != nil {
		return err
	}
	return nil
}
