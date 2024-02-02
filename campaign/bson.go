package campaign

import (
	"bwastartupecho/app"
	"bwastartupecho/helper"
	"bwastartupecho/user"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonMToCampign(b bson.M, c *Campaign) error {
	// var cam Campaign

	b["id"] = b["_id"].(primitive.ObjectID).Hex()
	bByte, err := bson.Marshal(b)
	helper.PanicIfError(err, " error in byting bson campaign")

	err = bson.Unmarshal(bByte, c)
	helper.PanicIfError(err, " error in ubmarshal bson campaign")
	return nil

}
func BsonMToCampignImage(b bson.M, c *CampaignImage) error {
	// var cam Campaign

	b["id"] = b["_id"].(primitive.ObjectID).Hex()
	bByte, err := bson.Marshal(b)
	helper.PanicIfError(err, " error in byting bson campaign")

	err = bson.Unmarshal(bByte, c)
	helper.PanicIfError(err, " error in ubmarshal bson campaign")
	return nil

}
func CampaignToBsonM(b *bson.M, c Campaign) error {
	bByte, err := bson.Marshal(c)
	helper.PanicIfError(err, " error in byting bson campaign")

	err = bson.Unmarshal(bByte, b)
	helper.PanicIfError(err, " error in ubmarshal bson campaign")
	return nil
}

// func (repo *repository) FindsAll(ctx context.Context, db app.MongoDatabase) ([]Campaign, error) {
// 	var cam Campaign
// 	var camss []Campaign
// 	var cc []bson.M
// 	cams, err := db.CampaignCollection().Find(ctx, bson.D{})
// 	helper.PanicIfError(err, "error in findind all campaign repo")
// 	err = cams.All(ctx, &cc)
// 	helper.PanicIfError(err, "error in unmarsal findind all campaign repo")

// 	for _, x := range cc {
// 		var us user.User
// 		var cis []CampaignImage
// 		var ci CampaignImage
// 		var cim []bson.M

// 		err = BsonMToCampign(x, &cam)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		objUserId, err := primitive.ObjectIDFromHex(cam.UserID)
// 		if err != nil {
// 			log.Println("user id is not valid", cam.UserID)
// 			continue
// 		}
// 		userFIlter := bson.D{{Key: "_id", Value: objUserId}}
// 		u := db.UserCollection().FindOne(ctx, userFIlter)
// 		err = u.Decode(&us)
// 		if err != nil {
// 			log.Println("eror in decode user", us)
// 			continue
// 		}
// 		ciFIlter := bson.D{{Key: "_id", Value: x["_id"]}}
// 		icurs, err := db.CampaignImageCollection().Find(ctx, ciFIlter)
// 		if err != nil {
// 			cam.User = us
// 			camss = append(camss, cam)
// 			continue
// 		}
// 		err = icurs.All(ctx, &cim)
// 		helper.PanicIfError(err, " error in unmasrhal cim")
// 		for _, cims := range cim {

// 			err := BsonMToCampignImage(cims, &ci)
// 			if err != nil {
// 				continue
// 			}
// 			cis = append(cis, ci)
// 		}
// 		cam.CampaignImages = cis
// 		cam.User = us
// 		camss = append(camss, cam)

// 	}
// 	fmt.Println(cams, "\n\n camss", camss, "\n\n", cam, "\n\n", cc)
// 	return camss, nil

// }

func campaignExtractor(ctx context.Context, db app.MongoDatabase, cam Campaign, x bson.M) (Campaign, error) {
	var us user.User
	var cis []CampaignImage
	var ci CampaignImage
	var cim []bson.M
	err := BsonMToCampign(x, &cam)
	if err != nil {
		log.Println(err)
		return cam, err
	}

	log.Println(cam.UserID, "madang sik")
	objUserId, err := primitive.ObjectIDFromHex(cam.UserID)
	if err != nil {
		log.Println("user id is not valid", cam.UserID)
		return cam, err
	}
	userFIlter := bson.D{{Key: "_id", Value: objUserId}}
	u := db.UserCollection().FindOne(ctx, userFIlter)
	err = u.Decode(&us)
	if err != nil {
		log.Println("eror in decode user", us)
		return cam, err
	}
	fmt.Printf("var1 = %T\n", x["_id"])
	fmt.Println("var1 ", x["_id"], cam.ID)
	ciFIlter := bson.D{{Key: "campaignid", Value: cam.ID}}
	icurs, err := db.CampaignImageCollection().Find(ctx, ciFIlter)
	fmt.Println(icurs, err, " madang sikbos ")
	if err == nil {
		err = icurs.All(ctx, &cim)
		fmt.Println(err, " madang si    kbos ")
		if err == nil {
			// helper.PanicIfError(err, " error in unmasrhal cim")
			// fmt.Println(cis, cim, "\n\n madang neh ", ci)
			for _, cims := range cim {
				fmt.Println("\n\n madang neh ", cims)

				err := BsonMToCampignImage(cims, &ci)
				if err != nil {
					fmt.Println("lewat kah?")
					continue
				}
				cis = append(cis, ci)
			}
		}

	}

	cam.CampaignImages = cis
	cam.User = us
	fmt.Println(cam, "cammmmmmmmera ")
	return cam, nil
}

// var us user.User
// var cis []CampaignImage
// var ci CampaignImage
// var cim []bson.M

// err = BsonMToCampign(x, &cam)
// if err != nil {
// 	log.Println(err)
// 	continue
// }
// objUserId, err := primitive.ObjectIDFromHex(cam.UserID)
// if err != nil {
// 	log.Println("user id is not valid", cam.UserID)
// 	continue
// }
// userFIlter := bson.D{{Key: "_id", Value: objUserId}}
// u := db.UserCollection().FindOne(ctx, userFIlter)
// err = u.Decode(&us)
// if err != nil {
// 	log.Println("eror in decode user", us)
// 	continue
// }
// ciFIlter := bson.D{{Key: "_id", Value: x["_id"]}}
// icurs, err := db.CampaignImageCollection().Find(ctx, ciFIlter)
// if err != nil {
// 	cam.User = us
// 	camss = append(camss, cam)
// 	continue
// }
// err = icurs.All(ctx, &cim)
// helper.PanicIfError(err, " error in unmasrhal cim")
// for _, cims := range cim {

// 	err := BsonMToCampignImage(cims, &ci)
// 	if err != nil {
// 		continue
// 	}
// 	cis = append(cis, ci)
// }
// cam.CampaignImages = cis
// cam.User = us
