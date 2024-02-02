// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {

// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

package main

import (
	"bwastartupecho/app"
	"bwastartupecho/auth"
	"bwastartupecho/campaign"
	"bwastartupecho/controller"
	"bwastartupecho/handler"
	"bwastartupecho/middleware"
	"bwastartupecho/payment"
	"bwastartupecho/transaction"
	"bwastartupecho/user"

	"github.com/labstack/echo"
)

func main() {
	// db := app.NewDB()
	// client := db.DbName().Client()
	// session, err := client.StartSession()
	// if err != nil {
	// 	return err
	// }
	// defer session.EndSession(ctx)
	// u := user.User{
	// 	ID:             "62f6145a1bc5cf4750bb4590",
	// 	Name:           "mamaaa",
	// 	Occupation:     "mamaaa",
	// 	Email:          "mamaa@gmail.comaa",
	// 	PasswordHash:   "mamaaa",
	// 	AvatarFileName: "mamaaa",
	// 	Role:           "mamaaa",
	// 	// Token:          sql.NullString{},
	// 	// CreatedAt: time.Now(),
	// 	// UpdatedAt: time.Now(),
	// }
	// input := user.RegisterUser{
	// 	Name:       "madang sik",
	// 	Occupation: "madang sik",
	// 	Email:      "madangsik@gmail.com",
	// 	Password:   "madang wae",
	// }
	// ctx := context.Background()

	// uu := user.NewRepository()
	// us := user.NewService(db, uu)
	// usr, err := us.RegisterUser(ctx, input)
	// fmt.Println(usr, err, "madang sik")
	// uuu, err := uu.UpdateUser(ctx, db, u)
	// // uuu, err := uu.SaveUser(ctx, db, u)
	// // uuu, err := uu.FindByEmail(ctx, db, "maamaa@gmail.com")
	// // uuu, err := uu.FindByID(ctx, db, "62f5f922b290fbc821667b49")
	// fmt.Println(uuu.ID, err, uuu, "\n", u)
	// u.FindByEmail(ctx, db,)
	// type Address struct {
	// 	Street string
	// 	City   string
	// 	State  string
	// }
	// type ss struct {
	// 	Makan  string
	// 	Minnum int
	// 	Add    Address
	// }
	// type ff struct {
	// 	Makan string
	// }
	// cvv := bson.D{}
	// cvvv := bson.D{}
	// var result ss
	// // tr := ss{
	// // 	Makan:  "sssss",
	// // 	Minnum: 40,
	// // 	Add:    Address{"1www Lakewood Way", "Elwood City", "PA"},
	// // }
	// tr := ff{
	// 	Makan: "sssss",
	// }
	// rt, err := bson.Marshal(tr)
	// err = bson.Unmarshal(rt, &cvv)
	// hh, err := db.CampaignCollection().InsertOne(ctx, cvv)
	// defer db.CloseDB()
	// fmt.Println(bson.D{{"type", "Oolong"}}, rt, err, string(rt), cvv, "llll", hh)

	// // type Student struct {
	// // 	FirstName string
	// // 	LastName  string
	// // 	Address   Address
	// // 	Age       int
	// // }
	// type Student struct {
	// 	FirstName string `bson:"first_name,omitempty"`
	// 	LastName  string `bson:"last_name,omitempty"`
	// 	// Address   Address
	// 	Address Address `bson:"inline"`
	// 	Age     int
	// }

	// // _id
	// // 62f5d28c71b75528ebaf0845
	// // var aaa string
	// var bsond bson.D
	// address1 := Address{"1 Lakewood Way", "Elwooooood City", "PAll"}
	// student1 := Student{FirstName: "9090090Arthur", Address: address1, Age: 8}
	// qq, err := bson.Marshal(student1)
	// err = bson.Unmarshal(qq, &bsond)
	// // hhh, err := db.CampaignCollection().InsertOne(ctx, student1)
	// id, err := primitive.ObjectIDFromHex("62f5d15212f9a595fa8db5d5")
	// fmt.Println(id, "iiiiiiii", err)
	// update, err := db.CampaignCollection().UpdateByID(ctx, id, bson.D{{Key: "$set", Value: student1}})
	// if err == nil {
	// 	fmt.Println(id, "iiiiiiii", err, update.MatchedCount, "\n", update.ModifiedCount, "  ", update.UpsertedID, " ", update, "dpppp", result, cvvv)
	// }
	// id, err = primitive.ObjectIDFromHex("62f5d28c71b75528ebaf0845")
	// fmt.Println(id, "iiiiiiii", err, " ", update, "dpppp", result, cvvv)
	// filter := bson.D{{Key: "_id", Value: id}}
	// resultr := db.CampaignCollection().FindOne(ctx, filter)
	// fmt.Println(resultr, resultr.Decode(&student1), "tttt", student1)
	// // eee := db.CampaignCollection().FindOne(ctx, tr).Decode(&result)
	// // eeee := db.CampaignCollection().FindOne(ctx, tr).Decode(&cvvv)
	// // fmt.Println("llll", hhh.InsertedID, eee, eeee, result, "dddddddddddd\n", cvvv, "\n", hhh.InsertedID.(primitive.ObjectID).Hex(), "\n", update)
	// // https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/#struct-tags
	// // https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
	// // https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID.Hex

	// autuser := auth.Newjwtservice()
	// aa, bb := autuser.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1MTA3NzYsImp0aSI6IjIwIn0.ek-ngmpRvOX-WC1uPApiyk7X_8iCcJiU3SAOR1Yi4R0")
	// aa, bb := autuser.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1MTExOTEsImp0aSI6IjYyZjkwNmYxNDU0MGUxMjJkMDQzNDhiNSJ9.p4x5e3qnFSxnb1QZlYhW3uRWv5vbLtCEPTdbAaaSnXM")

	// log.Println(aa, bb, "dddeded")
	autuser := auth.Newjwtservice()
	ps := payment.NewPaymentService()
	db := app.NewDB()

	uu := user.NewRepository()
	us := user.NewService(db, uu)
	uh := handler.NewUserHandler(us, autuser)

	cr := campaign.NewRepository()
	cs := campaign.NewService(cr, db)
	ch := handler.NewCampaignHandler(cs)

	tr := transaction.NewRepository()
	ts := transaction.NewServiceTransaction(tr, cr, ps, db)
	th := handler.NewTransactionHandler(ts)

	// ct := transaction.CreateTransactionInput{
	// 	Amount:     100000,
	// 	CampaignID: "62f9f4961f6db684bc69fdc3",
	// 	User: user.User{
	// 		ID: "62f906f14540e122d04348b5",
	// 	},
	// }

	// ccc := transaction.GetCampaignTransactionsInput{
	// 	ID: "62f9f4961f6db684bc69fdc3",
	// 	User: user.User{
	// 		// ID: "62f906f14540e122d04348b5",
	// 		ID: "62fa00be6d724e5fe60d78c3",
	// 	},
	// }
	// eeee := transaction.TransactionNotificationInput{
	// 	TransactionStatus: "deny",
	// 	OrderID:           "630302443d6d8638f77bfa41",
	// 	PaymentType:       "",
	// 	FraudStatus:       "",
	// }
	// rr := transaction.Transaction{
	// 	ID:         "630302443d6d8638f77bfa41",
	// 	CampaignID: "62f9f4961f6db684bc69fdc3",
	// 	UserID:     "62fa00be6d724e5fe60d78c3",
	// 	Amount:     105900,
	// 	Status:     "paid",
	// 	Code:       "pending_aja",
	// 	PaymentURL: "",
	// 	// User:       user.User{},
	// 	// Campaign:   campaign.Campaign{},
	// 	// CreatedAt:  time.Time{},
	// 	// UpdatedAt:  time.Time{},
	// }
	// ctx := context.Background()

	// tts, err := ts.CreateTransaction(ctx, ct)
	// tts, err := ts.GetTransactionByCampaignID(ctx, ccc)
	// tts, err := ts.GetTransactionByUserID(ctx, "62fa00be6d724e5fe60d78c3")
	// err := ts.ProcessPayment(ctx, eeee)
	// fmt.Println("\n\n", "tts", "\n", err, ct, "ctx bos", ccc)
	// xx, err := tr.SaveTransaction(ctx, db, rr)
	// xx, err := tr.Update(ctx, db, rr)
	// xx, err := tr.GetByUserId(ctx, db, "62fa00be6d724e5fe60d78c3")
	// xx, err := tr.GetByCampaignID(ctx, db, "62f9f4961f6db684bc69fdc3")
	// xx, err := tr.GetByTransactionId(ctx, db, "63030334c2a0dcb7a4edb51a")

	// log.Println(xx, " madang sik", err)
	// fmt.Println(rr)
	// xx, err := cr.FindAll(ctx, db)
	// xx, err := cr.FindById(ctx, db, "62f9f4961f6db684bc69fdc3")
	// xx, err := cr.FindById(ctx, db, "62f9f503e9b0f18ef833843a")
	// xx, err := cr.FindByUserId(ctx, db, "62f906f14540e122d04348b5")

	// c := campaign.Campaign{
	// 	ID:               "62f9f4961f6db684bc69fdc3",
	// 	UserID:           "62f906f14540e122d04348b5",
	// 	Name:             "tttt",
	// 	ShortDescription: "eeerrrrrr",
	// 	Description:      "eeee",
	// 	Perks:            "eee",
	// 	BackerCount:      0,
	// 	GoalAmount:       0,
	// 	CurrentAmount:    0,
	// 	Slug:             "",
	// 	// CreatedAt:        time.Time{},
	// 	// UpdatedAt:        time.Time{},
	// 	// CampaignImages:   []campaign.CampaignImage{},
	// 	// User:             user.User{},
	// }
	// ci := campaign.CampaignImage{
	// 	// ID:         "",
	// 	CampaignID: "62f9f4961f6db684bc69fdc3",
	// 	FileName:   "cccffffffcccc",
	// 	IsPrimary:  1,
	// 	// CreatedAt:  time.Time{},
	// 	// UpdatedAt:  time.Time{},
	// }
	// input := campaign.CreateCampaignInput{
	// 	Name:             "madang wae sik",
	// 	ShortDescription: "podo lek wae",
	// 	Description:      "ngelih",
	// 	GoalAmount:       1000000,
	// 	Perks:            "perrkkss",
	// 	User: user.User{
	// 		ID:         "62f906f14540e122d04348b5",
	// 		Name:       "hjhjhj",
	// 		Occupation: "okoo",
	// 		Email:      "eee@gmail.com",
	// 		// PasswordHash:   "",
	// 		AvatarFileName: "sdsds",
	// 		Role:           "ssss",
	// 		// Token:          sql.NullString{},
	// 		CreatedAt: time.Now(),
	// 		UpdatedAt: time.Now(),
	// 	},
	// }

	// inputimage := campaign.CreateCampaignImageInput{
	// 	CampaignID: "62fcb488f6acc1bd17a43a33",
	// 	IsPrimary:  true,
	// 	// User:       user.User{},
	// }

	// inputdetail := campaign.GetCampaignDetailInput{
	// 	ID: "62f9f4961f6db684bc69fdc3",
	// }
	// // xx, err := cs.CreateCampaign(ctx, input)
	// // xx, err := cs.SaveCampaignImage(ctx, inputimage, "madang sik")
	// // xx, err := cs.UpdateCampaign(ctx, inputdetail, input)
	// // xx, err := cs.GetCampaignById(ctx, inputdetail)
	// xx, err := cs.FindCampaigns(ctx, "t t")
	// // xx, err := cr.UpdateCampaign(ctx, db, c)

	// // xx, err := cr.MarkAllImagesAsNonPrimary(ctx, db, "62f9f4961f6db684bc69fdc3")
	// // xx, err := cr.SaveCampaign(ctx, db, c)
	// // xx, err := cr.SaveImage(ctx, db, ci)
	// fmt.Println("xx,rr\n", xx, "\n\n", err, c, ci, "\n", input, "\n", inputimage, "\n", inputdetail)
	mid := middleware.NewAutChecker(autuser, us)
	e := echo.New()
	e.Use(middleware.PanicHandler)

	cu := controller.NewUserHandler(e, uh, autuser, us, mid)
	cc := controller.NewCampaignHandler(cu, ch, mid)
	ct := controller.NewTransactionHandler(cc, th, mid)
	e.Logger.Fatal(ct.Start(":3000"))

}
