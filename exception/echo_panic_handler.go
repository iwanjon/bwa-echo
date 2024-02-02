package exception

import (
	"bwastartupecho/helper"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func EchoRecover(e echo.Context) {
	err := recover()
	if err != nil && echovalidationErrors(e, err) {
		return
	} else if err != nil && echonotOwnerError(e, err) {
		return
	} else if err != nil && echonotFoundError(e, err) {
		return
	} else if err != nil && echointernalServerError(e, err) {
		return
	}

	return
}

func echovalidationErrors(e echo.Context, err interface{}) bool {

	exception, ok := err.(validator.ValidationErrors)
	if ok {

		res := helper.APIResponse("error in validation", http.StatusBadRequest, "error", helper.FormatValidationError(exception))

		e.JSON(http.StatusBadGateway, res)
		return true
	} else {
		return false
	}
}

func echonotOwnerError(e echo.Context, err interface{}) bool {
	exception, ok := err.(NotOwnerError)
	if ok {

		webResponse := helper.APIResponse("Not The Owner", http.StatusNotAcceptable, "error", exception.Error)
		e.JSON(http.StatusBadGateway, webResponse)
		return true
	}
	return false
}

func echonotFoundError(e echo.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok && e.Request().URL.Path == "/api/v1/email_checkers" {
		fmt.Println("error", exception)
		webResponse := helper.APIResponse("email available", http.StatusOK, "success", true)
		e.JSON(http.StatusBadGateway, webResponse)
		return true
	} else if ok && e.Request().URL.Path != "/api/v1/email_checkers" {

		webResponse := helper.APIResponse("not found error", http.StatusNotFound, "error", exception.Error)
		e.JSON(http.StatusBadGateway, webResponse)
		return true

	} else {
		return false
	}
}

func echointernalServerError(e echo.Context, err interface{}) bool {

	log.Println(err, "error inside internal server error ")
	webResponse := helper.APIResponse("internal server error", http.StatusInternalServerError, "error", err)

	e.JSON(http.StatusInternalServerError, webResponse)
	return true
}
