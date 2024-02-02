package handler

import (
	"bwastartupecho/helper"
	mid "bwastartupecho/middleware"
	"bwastartupecho/transaction"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type transactionHandler struct {
	Service transaction.Service
}

type TransactionHandler interface {
	CreateTransaction(c echo.Context) error
	GetUserTransactions(c echo.Context) error
	GetCampaignTransactions(c echo.Context) error
	GetNotif(c echo.Context) error
}

func NewTransactionHandler(s transaction.Service) TransactionHandler {
	return &transactionHandler{
		Service: s,
	}
}

func (h *transactionHandler) CreateTransaction(c echo.Context) error {
	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	var input transaction.CreateTransactionInput
	err := c.Bind(&input)
	helper.PanicIfError(err, " error in binding transction iput handler")

	input.User = struct_context.CurrentUser
	t, err := h.Service.CreateTransaction(c.Request().Context(), input)
	helper.PanicIfError(err, " eror in create transaction")
	tt := transaction.FormatTransaction(t)
	res := helper.APIResponse("success", http.StatusOK, "success", tt)
	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) GetUserTransactions(c echo.Context) error {
	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	ts, err := h.Service.GetTransactionByUserID(c.Request().Context(), struct_context.CurrentUser.ID)
	helper.PanicIfError(err, "error in finding transaction handler")

	ft := transaction.FormatUserTransactions(ts)
	res := helper.APIResponse("success", http.StatusOK, "success", ft)
	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) GetCampaignTransactions(c echo.Context) error {
	var input transaction.GetCampaignTransactionsInput
	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	campaignID := c.Param("campaignid")
	// helper.PanicIfError(err, " error in binding input get campaign transaction handler")
	input.ID = campaignID
	input.User = struct_context.CurrentUser
	ts, err := h.Service.GetTransactionByCampaignID(c.Request().Context(), input)
	helper.PanicIfError(err, " error in get tranaction handler campaign")

	ft := transaction.FormatCampaignTransactions(ts)
	res := helper.APIResponse("success", http.StatusOK, "success", ft)
	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) GetNotif(c echo.Context) error {
	var input transaction.TransactionNotificationInput

	err := c.Bind(&input)
	helper.PanicIfError(err, " error in binding input notif")
	err = h.Service.ProcessPayment(c.Request().Context(), input)
	helper.PanicIfError(err, " error in update notif")
	res := helper.APIResponse("success", http.StatusOK, "success", input)
	return c.JSON(http.StatusOK, res)
}

// func (h *transactionHandler) CreateTransaction(writer http.ResponseWriter, request *http.Request) {

// 	struct_ctx_intf := request.Context().Value(mid.Contectkey)
// 	struct_context, ok := struct_ctx_intf.(mid.StructUser)
// 	if !ok {
// 		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

// 	}

// 	currentUser := struct_context.CurrentUser

// 	var input transaction.CreateTransactionInput
// 	helper.ReadFromRequestBody(request, &input)
// 	input.User = currentUser
// 	newTrans, err := h.Service.CreateTransaction(request.Context(), input)
// 	helper.PanicIfError(err, "error in creaye handler transaction")
// 	transFormater := transaction.FormatTransaction(newTrans)
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", transFormater)
// 	helper.WriteToResponseBody(writer, response)
// }

// func (h *transactionHandler) GetUserTransactions(writer http.ResponseWriter, request *http.Request) {
// 	struct_ctx_intf := request.Context().Value(mid.Contectkey)
// 	struct_context, ok := struct_ctx_intf.(mid.StructUser)
// 	if !ok {
// 		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

// 	}
// 	currentUser := struct_context.CurrentUser
// 	user_id := currentUser.ID
// 	transactions, err := h.Service.GetTransactionByUserID(request.Context(), user_id)
// 	helper.PanicIfError(err, "error in get transaction by user id handler ")
// 	transFormater := transaction.FormatUserTransactions(transactions)
// 	fmt.Println(transactions[0], "fffffffffffffffffff")
// 	fmt.Println(transFormater[0], "fffffffffffff")
// 	response := helper.APIResponse("success", http.StatusOK, "success get by user id campaign", transFormater)
// 	helper.WriteToResponseBody(writer, response)
// }

// func (h *transactionHandler) GetCampaignTransactions(writer http.ResponseWriter, request *http.Request) {

// 	var input transaction.GetCampaignTransactionsInput
// 	campaign_id := chi.URLParam(request, "campaignid")
// 	log.Println("this is id of the campaign", campaign_id)
// 	// campaign_id := params.ByName("campaignid")

// 	int_campaign_id, err := strconv.Atoi(campaign_id)
// 	helper.PanicIfError(err, "error convert to int campaign id handler transaction")

// 	struct_ctx_intf := request.Context().Value(mid.Contectkey)
// 	struct_context, ok := struct_ctx_intf.(mid.StructUser)
// 	if !ok {
// 		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

// 	}

// 	currentUser := struct_context.CurrentUser
// 	input.User = currentUser
// 	input.ID = int_campaign_id

// 	trans, err := h.Service.GetTransactionByCampaignID(request.Context(), input)
// 	helper.PanicIfError(err, " errror in get trans by campaign id handler transaction")
// 	transFormater := transaction.FormatCampaignTransactions(trans)
// 	fmt.Println(trans, "dddddd", transFormater)
// 	response := helper.APIResponse("success", http.StatusOK, "success get campaigns", transFormater)
// 	helper.WriteToResponseBody(writer, response)

// }

// func (h *transactionHandler) GetNotif(writer http.ResponseWriter, request *http.Request) {
// 	var input transaction.TransactionNotificationInput
// 	helper.ReadFromRequestBody(request, &input)
// 	err := h.Service.ProcessPayment(request.Context(), input)
// 	helper.PanicIfError(err, "error in process payment, handler transaction")
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", input)
// 	helper.WriteToResponseBody(writer, response)
// }
