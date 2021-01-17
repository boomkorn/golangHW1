package user

import (
	"app/db"
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	resource   *db.Resource
	collection *mongo.Collection
}

type repository interface {
	getAllUser() (Users, error)
	getByName(req ReqGetUserByName) (Users, error)
	postUser(req ReqPostUser) (Users, error)
}

func createUserRepository(resource *db.Resource) repository {
	collection := resource.DB.Collection("user")
	repository := &userRepository{resource: resource, collection: collection}
	return repository
}

func (ur *userRepository) getAllUser() (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepository) getByName(req ReqGetUserByName) (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{"name": req.Name})
	if err != nil {
		return Users{}, err
	}
	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepository) postUser(req ReqPostUser) (Users, error) {
	user := User{
		Id:   primitive.NewObjectID(),
		Name: req.Name,
		Age:  req.Age,
	}

	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	_, err := ur.collection.InsertOne(ctx, user)

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}
	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
