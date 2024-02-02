package transaction

import (
	"bwastartupecho/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TransactionToBsonM(ss Transaction) (bson.M, error) {
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

func BsonToTransaction(ss bson.M) (Transaction, error) {
	var b Transaction

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
