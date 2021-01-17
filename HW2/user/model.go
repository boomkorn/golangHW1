package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

type Users []User

// Request model
type ReqGetUserByName struct {
	Name string `bson:"name" json:"name"`
}
type ReqPostUser struct {
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
}
