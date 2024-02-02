package handler

import (
	"bwastartupecho/campaign"
	"bwastartupecho/helper"
	mid "bwastartupecho/middleware"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type campaignHandler struct {
	service campaign.Service
}

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
	UpdateCampaign(c echo.Context) error
	GetCampaigns(c echo.Context) error
	GetCampaign(c echo.Context) error
	UploadCampaignImage(c echo.Context) error
}

func NewCampaignHandler(s campaign.Service) *campaignHandler {
	return &campaignHandler{s}
}

func (h *campaignHandler) CreateCampaign(c echo.Context) error {
	var ci campaign.CreateCampaignInput
	err := c.Bind(&ci)
	helper.PanicIfError(err, " error in inding camapign input")
	newca, err := h.service.CreateCampaign(c.Request().Context(), ci)
	helper.PanicIfError(err, "error in create cmpiagn handler")

	cf := campaign.FormatCampaign(newca)
	res := helper.APIResponse("success", http.StatusOK, "success", cf)
	return c.JSON(http.StatusOK, res)
}

func (h *campaignHandler) UpdateCampaign(c echo.Context) error {
	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	var inputID campaign.GetCampaignDetailInput
	var inp campaign.CreateCampaignInput

	inputID.ID = c.Param("campaignid")
	err := c.Bind(&inp)
	helper.PanicIfError(err, " error in binging update campaign input")
	inp.User = struct_context.CurrentUser
	ca, err := h.service.UpdateCampaign(c.Request().Context(), inputID, inp)

	helper.PanicIfError(err, " error in update campaign handler")

	cf := campaign.FormatCampaign(ca)
	res := helper.APIResponse("success", http.StatusOK, "success", cf)
	return c.JSON(http.StatusOK, res)
}

func (h *campaignHandler) GetCampaigns(c echo.Context) error {

	userid := c.QueryParam("userid")
	cams, err := h.service.FindCampaigns(c.Request().Context(), userid)
	helper.PanicIfError(err, "error in finding campaigns handler")

	formatedcams := campaign.FormatCampaigns(cams)
	res := helper.APIResponse("success", http.StatusOK, "success", formatedcams)
	return c.JSON(http.StatusOK, res)
}

func (h *campaignHandler) GetCampaign(c echo.Context) error {
	campaignid := c.Param("campaignid")
	input := campaign.GetCampaignDetailInput{
		ID: campaignid,
	}
	cam, err := h.service.GetCampaignById(c.Request().Context(), input)
	helper.PanicIfError(err, " error in get campaign handler ")

	formatedcam := campaign.FormatCampaign(cam)
	res := helper.APIResponse("success", http.StatusOK, "success", formatedcam)
	return c.JSON(http.StatusOK, res)
}

func (h *campaignHandler) UploadCampaignImage(c echo.Context) error {

	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	user_id := struct_context.CurrentUser.ID
	// file, fileHeader, err := c.Request().FormFile("file")
	fileHeader, err := c.FormFile("file")
	helper.PanicIfError(err, " error in get file upload")
	isPrimary := c.FormValue("is_primary")
	campaignId := c.FormValue("campaign_id")

	input := campaign.CreateCampaignImageInput{
		CampaignID: campaignId,
		IsPrimary:  false,
		User:       struct_context.CurrentUser,
	}

	if isPrimary == "1" {
		input.IsPrimary = true
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	path := fmt.Sprintf("images/%s-%s", user_id, fileHeader.Filename)
	fileDestinatin, err := os.Create(path)
	helper.PanicIfError(err, " error in create path destination")
	defer fileDestinatin.Close()

	if _, err = io.Copy(fileDestinatin, file); err != nil {
		helper.PanicIfError(err, " error in copy to path")
	}

	_, err = h.service.SaveCampaignImage(c.Request().Context(), input, path)
	helper.PanicIfError(err, " error in save campaign image handler")

	mep := map[string]bool{
		"is_uploaded": true,
	}

	// mep["is_uploaded"] = true

	res := helper.APIResponse("success", http.StatusOK, "success", mep)

	return c.JSON(http.StatusOK, res)
}
