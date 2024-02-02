package user

import (
	"bwastartupecho/app"
	"bwastartupecho/helper"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(ctx context.Context, input RegisterUser) (User, error)
	LoginUser(ctx context.Context, input LoginInput) (User, error)
	CheckEmailAvailable(ctx context.Context, input CheckEmailInput) (bool, error)
	SaveAvatar(ctx context.Context, id string, filelocation string) (User, error)
	GetUserById(ctx context.Context, id string) (User, error)
}

type service struct {
	db   app.MongoDatabase
	repo Repository
	// Validate *validator.Validate
}

func NewService(db app.MongoDatabase, repo Repository) Service {
	return &service{db, repo}
}

func (s *service) RegisterUser(ctx context.Context, input RegisterUser) (User, error) {
	u := User{
		// ID:             "",
		Name:         input.Name,
		Occupation:   input.Occupation,
		Email:        input.Email,
		PasswordHash: input.Password,
		// AvatarFileName: "",
		// Role:           "",
		// Token:          sql.NullString{},
		// CreatedAt:      time.Time{},
		// UpdatedAt:      time.Time{},
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	helper.PanicIfError(err, "error in create hash password register user service")
	u.PasswordHash = string(bytes)

	registerUser, err := s.repo.SaveUser(ctx, s.db, u)
	helper.PanicIfError(err, " error in register user service")

	return registerUser, nil
}

func (s *service) LoginUser(ctx context.Context, input LoginInput) (User, error) {

	u, err := s.repo.FindByEmail(ctx, s.db, input.Email)
	helper.PanicIfError(err, "error in finding user by email login service")
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password))
	helper.PanicIfError(err, "error in compare pasword login service")
	return u, nil

}

func (s *service) CheckEmailAvailable(ctx context.Context, input CheckEmailInput) (bool, error) {
	_, err := s.repo.FindByEmail(ctx, s.db, input.Email)
	// exception.PanicIfNotFound(err, "error in finding email user service")
	if err != nil {
		return true, nil
	}
	return false, nil

}

func (s *service) SaveAvatar(ctx context.Context, id string, filelocation string) (User, error) {
	var u User
	u, err := s.repo.FindByID(ctx, s.db, id)
	helper.PanicIfError(err, "error in find user by id update service user")
	u.AvatarFileName = filelocation
	upuser, err := s.repo.UpdateUser(ctx, s.db, u)
	helper.PanicIfError(err, "error in update user")
	return upuser, nil
}

func (s *service) GetUserById(ctx context.Context, id string) (User, error) {
	u, err := s.repo.FindByID(ctx, s.db, id)
	helper.PanicIfError(err, " errror in fidning user by id")
	return u, nil
}

// ...

// func main() {
// 	// ...

// 	wc := writeconcern.New(writeconcern.W(1))
// 	rc := readconcern.Snapshot()
// 	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

// 	session, err := client.StartSession()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.EndSession(context.Background())

// 	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
// 		if err = session.StartTransaction(txnOpts); err != nil {
// 			return err
// 		}
// 		result, err := episodesCollection.InsertOne(
// 			sessionContext,
// 			Episode{
// 				Title:    "A Transaction Episode for the Ages",
// 				Duration: 15,
// 			},
// 		)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(result.InsertedID)
// 		result, err = episodesCollection.InsertOne(
// 			sessionContext,
// 			Episode{
// 				Title:    "Transactions for All",
// 				Duration: 1,
// 			},
// 		)
// 		if err != nil {
// 			return err
// 		}
// 		if err = session.CommitTransaction(sessionContext); err != nil {
// 			return err
// 		}
// 		fmt.Println(result.InsertedID)
// 		return nil
// 	})
// 	if err != nil {
// 		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
// 			panic(abortErr)
// 		}
// 		panic(err)
// 	}
// }
