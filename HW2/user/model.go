package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Age           int                `bson:"age" json:"age"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"password"`
	PhoneNo       string             `bson:"phoneNo" json:"phoneNo"`
	CreatedDate   time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate   time.Time          `bson:"updatedDate" json:"updatedDate"`
	LastLoginDate time.Time          `bson:"lastLoginDate" json:"lastLoginDate"`
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
type ReqGetUserByEmail struct {
	Email string `bson:"email" json:"email"`
}

// API request model
type ReqRegister struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`

	// CustomValidator struct {
	// 	validator *validator.Validate
	// }
}
type ReqLogin struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
type ReqSaveProfile struct {
	Name    string `bson:"name" json:"name"`
	Age     int    `bson:"age" json:"age"`
	PhoneNo string `bson:"phoneNo" json:"phoneNo"`
}
type ReqChangePassword struct {
	// Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// API response model
type ResponseDefault struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
type ResponseLogin struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       ReponseWithToken `json:"data"`
}
type ReponseWithToken struct {
	UserDetail ResponseUser `json:"userDetail"`
	Token      string       `json:"token"`
}
type ResponseWithUser struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       ResponseUser `json:"data"`
}
type ResponseUser struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Age           int                `bson:"age" json:"age"`
	Email         string             `bson:"email" json:"email"`
	PhoneNo       string             `bson:"phoneNo" json:"phoneNo"`
	CreatedDate   time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate   time.Time          `bson:"updatedDate" json:"updatedDate"`
	LastLoginDate time.Time          `bson:"lastLoginDate" json:"lastLoginDate"`
}
