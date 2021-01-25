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

	register(req *ReqRegister) error
	login(req *ReqLogin) (ResponseUser, error)
	saveProfile(userId string, req *ReqSaveProfile) (ResponseUser, error)
	getProfile(userId string) (ResponseUser, error)
	updatePassword(userId string, req *ReqChangePassword) error
}

func createUserRepository(resource *db.Resource) repository {
	collection := resource.DB.Collection("user")
	repository := &userRepository{resource: resource, collection: collection}
	return repository
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

func (ur *userRepository) register(req *ReqRegister) error {
	ctx, cancel := initContext()
	defer cancel()

	// data := []byte(req.Password)

	user := User{
		Id:    primitive.NewObjectID(),
		Email: req.Email,
		// Password:      fmt.Sprintf("%x", md5.Sum(data)),
		Password:      req.Password,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	_, err := ur.collection.InsertOne(ctx, user)
	return err
}
func (ur *userRepository) login(req *ReqLogin) (ResponseUser, error) {
	ctx, cancel := initContext()
	defer cancel()

	var user = ResponseUser{}
	err := ur.collection.FindOne(ctx, bson.M{"email": req.Email, "password": req.Password}).Decode(&user)
	if err != nil {
		return ResponseUser{}, err
	}

	user.LastLoginDate = time.Now()

	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": bson.M{
			"lastLoginDate": time.Now(),
		},
	}

	_, err = ur.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ResponseUser{}, err
	}

	return ur.getProfile(user.Id.Hex())
}
func (ur *userRepository) saveProfile(userId string, req *ReqSaveProfile) (ResponseUser, error) {
	ctx, cancel := initContext()
	defer cancel()

	var user = ResponseUser{}
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return ResponseUser{}, err
	}
	err = ur.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return ResponseUser{}, err
	}

	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"age":         req.Age,
			"phoneNo":     req.PhoneNo,
			"updatedDate": time.Now(),
		},
	}

	_, err = ur.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ResponseUser{}, err
	}

	return ur.getProfile(userId)
}

func (ur *userRepository) getProfile(userId string) (ResponseUser, error) {
	ctx, cancel := initContext()
	defer cancel()

	var user = ResponseUser{}
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return ResponseUser{}, err
	}
	err = ur.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return ResponseUser{}, err
	}
	return user, nil
}

func (ur *userRepository) updatePassword(userId string, req *ReqChangePassword) error {
	ctx, cancel := initContext()
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{
		"$set": bson.M{
			"password":    req.Password,
			"updatedDate": time.Now(),
		},
	}

	_, err = ur.collection.UpdateOne(ctx, filter, update)

	return err
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
