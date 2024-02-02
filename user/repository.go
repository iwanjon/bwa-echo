package user

import (
	"bwastartupecho/app"
	"bwastartupecho/exception"
	"bwastartupecho/helper"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository struct {
}

type Repository interface {
	SaveUser(ctx context.Context, db app.MongoDatabase, user User) (User, error)
	UpdateUser(ctx context.Context, db app.MongoDatabase, user User) (User, error)
	FindByEmail(ctx context.Context, db app.MongoDatabase, email string) (User, error)
	FindByID(ctx context.Context, db app.MongoDatabase, id string) (User, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) SaveUser(ctx context.Context, db app.MongoDatabase, user User) (User, error) {
	user.CreatedAt = time.Now()
	bsonuser, err := UserToBsonM(user)
	helper.PanicIfError(err, "error in create bson")
	// delete(bsonuser, "updatedat")
	// fmt.Println(bsonuser)

	result, err := db.UserCollection().InsertOne(ctx, bsonuser)
	helper.PanicIfError(err, " errror in save user repo")
	stringId := result.InsertedID.(primitive.ObjectID).Hex()
	user.ID = stringId
	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, db app.MongoDatabase, user User) (User, error) {
	user.UpdatedAt = time.Now()
	bsonuser, err := UserToBsonM(user)
	helper.PanicIfError(err, "error in create bson")
	delete(bsonuser, "createdat")
	fmt.Println(bsonuser)
	id, err := primitive.ObjectIDFromHex(user.ID)
	// bson.D{{"$set", user}}
	helper.PanicIfError(err, " eerror in create onject id repo user")
	result, err := db.UserCollection().UpdateByID(ctx, id, bson.D{{Key: "$set", Value: bsonuser}})
	helper.PanicIfError(err, " errror in save user repo")
	log.Println(result)

	return user, nil
}

func (r *repository) FindByEmail(ctx context.Context, db app.MongoDatabase, email string) (User, error) {
	var userI User
	var userb bson.M
	filter := bson.D{{Key: "email", Value: email}}
	result := db.UserCollection().FindOne(ctx, filter)
	err := result.Decode(&userb)
	// helper.PanicIfError(err, " errro in decode reult find by email repo")
	exception.PanicIfNotFound(err, "error in findnign email user repo")
	// if err != nil {
	// 	return userI, err
	// }
	fmt.Println(userb["_id"], "madanag")
	// userb["ID"] = userb["_id"].(primitive.ObjectID).Hex()
	userI, err = BsonToUser(userb)
	helper.PanicIfError(err, "eroor in bson to user ")
	// userI, err := helper.BsonToInterface(userb)
	// fmt.Println(userb, "user b \n\n", userI)
	// ne, er := userI.(User)
	// if !er {
	// 	helper.PanicIfError(err, " error inassertion user")
	// }

	// fmt.Println("neeee", userI, userb["_id"].(primitive.ObjectID).Hex(), "\n\n", userb, "\n\n", "nee")

	return userI, nil

}

func (r *repository) FindByID(ctx context.Context, db app.MongoDatabase, id string) (User, error) {
	var user User
	primId, err := primitive.ObjectIDFromHex(id)
	helper.PanicIfError(err, " errro in conver object id find by id repo")
	filter := bson.D{{Key: "_id", Value: primId}}
	result := db.UserCollection().FindOne(ctx, filter)
	err = result.Decode(&user)
	helper.PanicIfError(err, " errro in decode reult find by id repo")
	user.ID = id

	return user, nil

}
