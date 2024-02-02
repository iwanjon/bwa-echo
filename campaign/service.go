package campaign

import (
	"bwastartupecho/app"
	"bwastartupecho/exception"
	"bwastartupecho/helper"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Service interface {
	FindCampaigns(ctx context.Context, userId string) ([]Campaign, error)
	GetCampaignById(ctx context.Context, input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(ctx context.Context, input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(ctx context.Context, inputID GetCampaignDetailInput, inputparam CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(ctx context.Context, input CreateCampaignImageInput, filelocation string) (CampaignImage, error)
}

type service struct {
	repo Repository
	db   app.MongoDatabase
	// Validate *validator.Validate
}

func NewService(repo Repository, db app.MongoDatabase) *service {
	return &service{repo, db}
}

func (s *service) FindCampaigns(ctx context.Context, userId string) ([]Campaign, error) {
	if strings.TrimSpace(userId) != "" {
		campaigns, err := s.repo.FindByUserId(ctx, s.db, userId)
		helper.PanicIfError(err, "error in get campaign by userid")
		return campaigns, nil

	}
	campaigns, err := s.repo.FindAll(ctx, s.db)
	helper.PanicIfError(err, " error in get campaigns")

	return campaigns, nil
}

func (s *service) GetCampaignById(ctx context.Context, input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repo.FindById(ctx, s.db, input.ID)
	helper.PanicIfError(err, "error in find campaign by id service")

	return campaign, nil
}

func (s *service) CreateCampaign(ctx context.Context, input CreateCampaignInput) (Campaign, error) {

	cam := Campaign{
		// ID:               "",
		UserID:           input.User.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		BackerCount:      0,
		GoalAmount:       input.GoalAmount,
		CurrentAmount:    0,
		Slug:             fmt.Sprintf("%s-%s", input.User.ID, input.Name),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		// CampaignImages:   []CampaignImage{},
		// User:             user.User{},
	}

	camp, err := s.repo.SaveCampaign(ctx, s.db, cam)
	helper.PanicIfError(err, " error in save campaign")
	return camp, nil
}

func (s *service) UpdateCampaign(ctx context.Context, inputID GetCampaignDetailInput, inputparam CreateCampaignInput) (Campaign, error) {

	cam := Campaign{
		ID: inputID.ID,
		// UserID:           "",
		Name:             inputparam.Name,
		ShortDescription: inputparam.ShortDescription,
		Description:      inputparam.Description,
		Perks:            inputparam.Perks,
		// BackerCount:      0,
		GoalAmount: inputparam.GoalAmount,
		// CurrentAmount:    "",
		Slug: fmt.Sprintf("%s-%s", inputparam.User.ID, inputparam.Name),
		// CreatedAt:        time.Time{},
		UpdatedAt: time.Now(),
		// CampaignImages:   []CampaignImage{},
		User: inputparam.User,
	}

	ca, err := s.repo.FindById(ctx, s.db, inputID.ID)
	helper.PanicIfError(err, " error in get campiang by id service")

	if ca.UserID != inputparam.User.ID {
		exception.PanicIfNotOwner(errors.New("error in compare owner"), " eror not owner")
	}

	updated, err := s.repo.UpdateCampaign(ctx, s.db, cam)
	helper.PanicIfError(err, " error in updated campaign")
	return updated, nil
}

func (s *service) SaveCampaignImage(ctx context.Context, input CreateCampaignImageInput, filelocation string) (CampaignImage, error) {
	ci := CampaignImage{
		// ID:         "",
		CampaignID: input.CampaignID,
		FileName:   filelocation,
		IsPrimary:  0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	ca, err := s.repo.FindById(ctx, s.db, input.CampaignID)
	helper.PanicIfError(err, " error in finding campaign service")

	if ca.UserID != input.User.ID {
		exception.PanicIfNotOwner(errors.New("eror not owner "), " error not the owner")
	}
	if input.IsPrimary {
		ci.IsPrimary = 1
	}

	sci, err := s.repo.SaveImage(ctx, s.db, ci)
	helper.PanicIfError(err, " error in save imaeg")

	return sci, nil
}
