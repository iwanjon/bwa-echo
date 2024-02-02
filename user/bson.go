package user

import (
	"bwastartupecho/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserToBsonM(ss User) (bson.M, error) {
	var b bson.M

	qq, err := bson.Marshal(ss)
	// PanicIfError(err, " error in marshal struct")
	if err != nil {
		return b, err
	}
	err = bson.Unmarshal(qq, &b)
	// PanicIfError(err, " error in conver to BSON D")
	if err != nil {
		return b, err
	}
	return b, nil
}
func BsonToUser(ss bson.M) (User, error) {
	var b User

	ss["id"] = ss["_id"].(primitive.ObjectID).Hex()
	qq, err := bson.Marshal(ss)
	// PanicIfError(err, " error in marshal struct")
	if err != nil {
		return b, err
	}
	err = bson.Unmarshal(qq, &b)
	helper.PanicIfError(err, " error in conver to BSON D")

	return b, nil
}
