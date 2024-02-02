package campaign

import (
	"bwastartupecho/app"
	"bwastartupecho/helper"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	FindAll(ctx context.Context, db app.MongoDatabase) ([]Campaign, error)
	FindByUserId(ctx context.Context, db app.MongoDatabase, userId string) ([]Campaign, error)
	FindById(ctx context.Context, db app.MongoDatabase, campaignId string) (Campaign, error)
	SaveCampaign(ctx context.Context, db app.MongoDatabase, campaign Campaign) (Campaign, error)
	UpdateCampaign(ctx context.Context, db app.MongoDatabase, campaign Campaign) (Campaign, error)
	SaveImage(ctx context.Context, db app.MongoDatabase, campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(ctx context.Context, db app.MongoDatabase, campaignId string) (bool, error)
}

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (repo *repository) FindAll(ctx context.Context, db app.MongoDatabase) ([]Campaign, error) {
	var cam Campaign
	var camss []Campaign
	var cc []bson.M
	cams, err := db.CampaignCollection().Find(ctx, bson.D{})
	helper.PanicIfError(err, "error in findind all campaign repo")
	err = cams.All(ctx, &cc)
	helper.PanicIfError(err, "error in unmarsal findind all campaign repo")
	// aa, err := bson.Marshal(cc)
	// err = bson.Unmarshal(aa, &camss)
	// helper.PanicIfError(err, "error in unmarsal findind all campaign repo")
	for _, x := range cc {
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
		ca, err := campaignExtractor(ctx, db, cam, x)
		if err != nil {
			continue
		}
		camss = append(camss, ca)

	}
	fmt.Println(cams, "\n\n camss", camss, "\n\n", cam, "\n\n", cc)
	return camss, nil

}

func (repo *repository) SaveCampaign(ctx context.Context, db app.MongoDatabase, campaign Campaign) (Campaign, error) {
	campaign.CreatedAt = time.Now()
	id, err := db.CampaignCollection().InsertOne(ctx, campaign)
	helper.PanicIfError(err, " error in insert capaign ")
	stringId := id.InsertedID.(primitive.ObjectID).Hex()
	campaign.ID = stringId
	return campaign, nil
}
func (repo *repository) SaveImage(ctx context.Context, db app.MongoDatabase, campaignImage CampaignImage) (CampaignImage, error) {
	campaignImage.CreatedAt = time.Now()
	id, err := db.CampaignImageCollection().InsertOne(ctx, campaignImage)
	helper.PanicIfError(err, " error in insert capaignimage ")
	stringId := id.InsertedID.(primitive.ObjectID).Hex()
	campaignImage.ID = stringId
	return campaignImage, nil

}

func (repo *repository) FindByUserId(ctx context.Context, db app.MongoDatabase, userId string) ([]Campaign, error) {
	var bsonm []bson.M
	var cam Campaign
	var camps []Campaign

	// objUserId, err := primitive.ObjectIDFromHex(userId)
	// helper.PanicIfError(err, " error invalid object id")
	// userFIlter := bson.D{{Key: "_id", Value: objUserId}}
	// filter := bson.D{{"_id", userId}}
	filter := bson.D{{Key: "userid", Value: userId}}
	cur, err := db.CampaignCollection().Find(ctx, filter)
	helper.PanicIfError(err, " error in find by user id campaign repo")

	err = cur.All(ctx, &bsonm)
	helper.PanicIfError(err, " error in extract to bsonm campaign repo")

	// log.Println(bsonm, "bsom")
	for _, x := range bsonm {

		camp, err := campaignExtractor(ctx, db, cam, x)
		// helper.PanicIfError(err, " error in extract campaign bby user id")
		if err != nil {
			continue
		}
		camps = append(camps, camp)
	}

	return camps, nil

}

func (repo *repository) FindById(ctx context.Context, db app.MongoDatabase, campaignId string) (Campaign, error) {
	var campaign Campaign
	var bm bson.M
	objid, err := primitive.ObjectIDFromHex(campaignId)
	helper.PanicIfError(err, " erro in convert ot object id")
	filter := bson.D{{Key: "_id", Value: objid}}
	c := db.CampaignCollection().FindOne(ctx, filter)
	err = c.Decode(&bm)
	helper.PanicIfError(err, " error in decode campaign ")
	campaign, err = campaignExtractor(ctx, db, campaign, bm)
	helper.PanicIfError(err, " erro in campign extrator ")
	campaign.ID = campaignId
	return campaign, nil

}

func (repo *repository) UpdateCampaign(ctx context.Context, db app.MongoDatabase, campaign Campaign) (Campaign, error) {
	objid, err := primitive.ObjectIDFromHex(campaign.ID)
	helper.PanicIfError(err, " erro in convert ot object id")

	var bsonuser bson.M
	campaign.UpdatedAt = time.Now()
	filter := bson.D{{Key: "_id", Value: objid}}

	err = CampaignToBsonM(&bsonuser, campaign)
	helper.PanicIfError(err, " erro in convert to bson M")

	delete(bsonuser, "createdat")
	delete(bsonuser, "userid")

	ss := db.CampaignCollection().FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: bsonuser}})
	helper.PanicIfError(ss.Err(), " error in updateing campaign")

	return campaign, nil
}

func (repo *repository) MarkAllImagesAsNonPrimary(ctx context.Context, db app.MongoDatabase, campaignId string) (bool, error) {
	// objid, err := primitive.ObjectIDFromHex(campaignId)
	// helper.PanicIfError(err, " erro in convert ot object id")

	filter := bson.D{{Key: "campaignid", Value: campaignId}}
	update := bson.M{"$set": bson.M{"isprimary": 0}}
	dd, err := db.CampaignImageCollection().UpdateMany(ctx, filter, update)
	helper.PanicIfError(err, " erro in update many")
	if dd.ModifiedCount == 0 {
		log.Println("no campaign image update to non primary")
	}
	fmt.Printf("dd: %v\n", dd)
	return true, nil
}
