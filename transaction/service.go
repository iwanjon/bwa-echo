package transaction

import (
	"bwastartupecho/app"
	"bwastartupecho/campaign"
	"bwastartupecho/exception"
	"bwastartupecho/helper"
	"bwastartupecho/payment"
	"context"
	"errors"
	"fmt"
)

type service struct {
	repo               Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
	db                 app.MongoDatabase
}

type Service interface {
	GetTransactionByCampaignID(ctx context.Context, input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(ctx context.Context, UserID string) ([]Transaction, error)
	CreateTransaction(ctx context.Context, input CreateTransactionInput) (Transaction, error)
	ProcessPayment(ctx context.Context, trans TransactionNotificationInput) error
}

func NewServiceTransaction(r Repository, campaignRepository campaign.Repository, payservice payment.Service, db app.MongoDatabase) Service {
	return &service{r, campaignRepository, payservice, db}
}

func (s *service) GetTransactionByCampaignID(ctx context.Context, input GetCampaignTransactionsInput) ([]Transaction, error) {

	t, err := s.campaignRepository.FindById(ctx, s.db, input.ID)
	helper.PanicIfError(err, " error in finding campaign by id service trans")
	if t.UserID != input.User.ID {
		exception.PanicIfNotOwner(errors.New("error not owner"), " not owner of transaction")
	}

	ts, err := s.repo.GetByCampaignID(ctx, s.db, input.ID)
	helper.PanicIfError(err, " error in get transaction by campaign id, service")

	return ts, nil
}

func (s *service) GetTransactionByUserID(ctx context.Context, UserID string) ([]Transaction, error) {
	t, err := s.repo.GetByUserId(ctx, s.db, UserID)
	helper.PanicIfError(err, " error in find trasn by user id service")
	return t, nil

}

func (s *service) CreateTransaction(ctx context.Context, input CreateTransactionInput) (Transaction, error) {

	trasn := Transaction{
		// ID:         "",
		CampaignID: input.CampaignID,
		UserID:     input.User.ID,
		Amount:     input.Amount,
		Status:     "pending",
		Code:       fmt.Sprintf("%s-%s", input.User.ID, input.CampaignID),
		PaymentURL: "url",
		// User:       user.User{},
		// Campaign:   campaign.Campaign{},
		// CreatedAt:  time.Time{},
		// UpdatedAt:  time.Time{},
	}
	trans, err := s.repo.SaveTransaction(ctx, s.db, trasn)
	helper.PanicIfError(err, " error in save trab=ntsaction create transaction serice")

	p := payment.Transaction{
		ID:     trans.ID,
		Amount: trans.Amount,
	}
	url, err := s.paymentService.GetPaymentUrl(p, input.User)
	helper.PanicIfError(err, "error in creat epayment url")
	trans.PaymentURL = url

	updatedTrans, err := s.repo.Update(ctx, s.db, trans)

	helper.PanicIfError(err, " error in update new transaction url")
	return updatedTrans, nil
}

func (s *service) ProcessPayment(ctx context.Context, input TransactionNotificationInput) error {
	// 	trans_id, err := strconv.Atoi(input.OrderID)
	// 	helper.PanicIfError(err, " erro in conv order id service ")
	// 	// fmt.Println(trans_id)
	// 	transaction, err := s.repo.GetByTransactionId(ctx, s.db, trans_id)
	// 	helper.PanicIfError(err, " error in get transaction by id service")
	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
	// 		transaction.Status = "paid"
	// 	} else if input.TransactionStatus == "settlement" {
	// 		transaction.Status = "paid"
	// 	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
	// 		transaction.Status = "cancelled"
	// 	}
	// 	tx, err := s.db.Begin()

	// 	helper.PanicIfError(err, " error in create transaction db in service trans")
	// 	defer helper.CommitOrRollback(tx)
	// 	updated, err := s.repo.Update(ctx, tx, transaction)
	// 	helper.PanicIfError(err, " error in update paid status")

	// 	c, err := s.campaignRepository.FindById(ctx, s.db, updated.CampaignID)
	// 	helper.PanicIfError(err, " error in finding campaign transaction service")

	// 	if transaction.Status == "paid" {
	// 		c.BackerCount += 1
	// 		c.CurrentAmount += updated.Amount
	// 		_, err = s.campaignRepository.UpdateCampaign(ctx, tx, c)
	// 		helper.PanicIfError(err, " error in update campaign notif transaction")
	// 	}
	// 	return nil
	transaction, err := s.repo.GetByTransactionId(ctx, s.db, input.OrderID)
	helper.PanicIfError(err, " error in get trans by trans id service")
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTrns, err := s.repo.Update(ctx, s.db, transaction)
	helper.PanicIfError(err, " error in update trnsaxtion status")

	ca, err := s.campaignRepository.FindById(ctx, s.db, updatedTrns.CampaignID)
	helper.PanicIfError(err, " error in get campaign to be updated")

	if transaction.Status == "paid" {
		ca.BackerCount += 1
		ca.CurrentAmount += updatedTrns.Amount

		_, err := s.campaignRepository.UpdateCampaign(ctx, s.db, ca)
		helper.PanicIfError(err, " error in update campaign in transaction")
	}
	return nil

}

// func (s *service) GetTransactionByCampaignID(ctx context.Context, input GetCampaignTransactionsInput) ([]Transaction, error) {

// 	campaig, err := s.campaignRepository.FindById(ctx, s.db, input.ID)
// 	helper.PanicIfError(err, " erro in getting campgin by id service transaction")

// 	if campaig.User.ID != input.User.ID {
// 		exception.PanicIfNotOwner(errors.New("error not owner"), " error in ceck owner and user ")
// 	}

// 	trans, err := s.repo.GetByCampaignID(ctx, s.db, input.ID)
// 	helper.PanicIfError(err, " error in get campaign y id service")
// 	return trans, nil

// }

// func (s *service) GetTransactionByUserID(ctx context.Context, UserID int) ([]Transaction, error) {
// 	trans, err := s.repo.GetByUserId(ctx, s.db, UserID)
// 	helper.PanicIfError(err, " error in get trans by usr id")

// 	return trans, nil
// }

// func (s *service) CreateTransaction(ctx context.Context, input CreateTransactionInput) (Transaction, error) {
// 	var tran Transaction
// 	tran.Amount = input.Amount
// 	tran.CampaignID = input.CampaignID
// 	tran.UserID = input.User.ID
// 	tran.User = input.User
// 	tran.Status = "pending"
// 	tran.PaymentURL = "url"
// 	tran.Code = fmt.Sprintf("%d-%d-%d", input.User.ID, input.CampaignID, input.Amount)

// 	tx, err := s.db.Begin()
// 	helper.PanicIfError(err, " error in create transaction save transaction service")
// 	defer helper.CommitOrRollback(tx)
// 	newtrans, err := s.repo.SaveTransaction(ctx, tx, tran)
// 	helper.PanicIfError(err, " error in save transaction in service transaction")

// 	p := payment.Transaction{
// 		ID:     newtrans.ID,
// 		Amount: newtrans.Amount,
// 	}
// 	ur, err := s.paymentService.GetPaymentUrl(p, input.User)
// 	helper.PanicIfError(err, "error in create url payment transaction service")
// 	newtrans.PaymentURL = ur

// 	updatedtran, err := s.repo.Update(ctx, tx, newtrans)
// 	helper.PanicIfError(err, " error in updated tran service")

// 	return updatedtran, nil
// }

// func (s *service) ProcessPayment(ctx context.Context, input TransactionNotificationInput) error {
// 	trans_id, err := strconv.Atoi(input.OrderID)
// 	helper.PanicIfError(err, " erro in conv order id service ")
// 	// fmt.Println(trans_id)
// 	transaction, err := s.repo.GetByTransactionId(ctx, s.db, trans_id)
// 	helper.PanicIfError(err, " error in get transaction by id service")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "settlement" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
// 		transaction.Status = "cancelled"
// 	}
// 	tx, err := s.db.Begin()

// 	helper.PanicIfError(err, " error in create transaction db in service trans")
// 	defer helper.CommitOrRollback(tx)
// 	updated, err := s.repo.Update(ctx, tx, transaction)
// 	helper.PanicIfError(err, " error in update paid status")

// 	c, err := s.campaignRepository.FindById(ctx, s.db, updated.CampaignID)
// 	helper.PanicIfError(err, " error in finding campaign transaction service")

// 	if transaction.Status == "paid" {
// 		c.BackerCount += 1
// 		c.CurrentAmount += updated.Amount
// 		_, err = s.campaignRepository.UpdateCampaign(ctx, tx, c)
// 		helper.PanicIfError(err, " error in update campaign notif transaction")
// 	}
// 	return nil

// }

// func (s *service) ProcessPayment(ctx context.Context, input TransactionNotificationInput) error {
// 	trans_id, err := strconv.Atoi(input.OrderID)
// 	helper.PanicIfError(err, " erro in conv order id service ")
// 	// fmt.Println(trans_id)
// 	transaction, err := s.repo.GetByTransactionId(ctx, trans_id)
// 	helper.PanicIfError(err, " error in get transaction by id service")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "settlement" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
// 		transaction.Status = "cancelled"
// 	}

// 	updatedtransaction, err := s.repo.Update(ctx, transaction)
// 	helper.PanicIfError(err, " error in update transaction service")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	campaign, err := s.campaignRepository.FindById(ctx, updatedtransaction.CampaignID)
// 	helper.PanicIfError(err, " error in get campaign by id")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	if transaction.Status == "paid" {
// 		campaign.BackerCount = campaign.BackerCount + 1
// 		campaign.CurrentAmount = campaign.CurrentAmount + updatedtransaction.Amount

// 		_, err := s.campaignRepository.UpdateCampaign(ctx, campaign)
// 		helper.PanicIfError(err, " error in update campaign by id service transaction")
// 		// if err != nil {
// 		// 	return err
// 		// }
// 	}

// 	return nil
// }

// func (s *service) CreateTransaction(ctx context.Context, input CreateTransactionInput) (Transaction, error) {
// 	var trans Transaction

// 	trans.Amount = input.Amount
// 	trans.User = input.User
// 	trans.CampaignID = input.CampaignID
// 	trans.Status = "pending"

// 	trans.Code = "muamama"
// 	trans.UserID = input.User.ID

// 	// fmt.Println(trans, "madang trans")
// 	newtrans, err := s.repo.SaveTransaction(ctx, trans)
// 	helper.PanicIfError(err, " error in create transaction service")
// 	// if err != nil {
// 	// 	return newtrans, err
// 	// }

// 	paymentTransaction := payment.Transaction{
// 		ID:     newtrans.ID,
// 		Amount: newtrans.Amount,
// 	}

// 	url, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
// 	// if err != nil {
// 	// 	fmt.Println("error url")
// 	// 	return newtrans, err
// 	// }
// 	helper.PanicIfError(err, " error in create payment url service")
// 	newtrans.PaymentURL = url
// 	newtranss, err := s.repo.Update(ctx, newtrans)
// 	helper.PanicIfError(err, " erro in update transaction url serviced")
// 	// if err != nil {
// 	// 	return newtranss, err
// 	// }
// 	return newtranss, nil
// }

// func (s *service) GetTransactionByUserID(ctx context.Context, UserID int) ([]Transaction, error) {
// 	var transactions []Transaction
// 	// fmt.Println("gettransbyuserid", UserID)
// 	transactions, err := s.repo.GetByUserId(ctx, UserID)
// 	helper.PanicIfError(err, " error in get transaction by user id service")
// 	// if err != nil {
// 	// 	return transactions, err
// 	// }
// 	return transactions, nil
// }

// func (s *service) GetTransactionByCampaignID(ctx context.Context, input GetCampaignTransactionsInput) ([]Transaction, error) {
// 	var transactions []Transaction
// 	campaign, err := s.campaignRepository.FindById(ctx, input.ID)
// 	helper.PanicIfError(err, " erro in get acampaign by id service")
// 	// if err != nil {
// 	// 	return transactions, err
// 	// }

// 	if campaign.UserID != input.User.ID {
// 		// fmt.Println("error not owner check")
// 		// fmt.Println("error not owner check", campaign.UserID, input.User.ID)
// 		// return transactions, errors.New("not an owner of campaign")
// 		exception.PanicIfNotOwner(errors.New("error not owner of campaign"), " error in campaign authorized ")
// 	}

// 	transactions, err = s.repo.GetByCampaignID(ctx, input.ID)
// 	helper.PanicIfError(err, " erro in get find transacstion by campaign id service")
// 	// if err != nil {
// 	// 	return transactions, err
// 	// }
// 	return transactions, nil
// }
