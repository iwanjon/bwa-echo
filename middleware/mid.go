package middleware

import (
	"bwastartupecho/auth"
	"bwastartupecho/exception"
	"bwastartupecho/helper"
	"bwastartupecho/user"
	"log"

	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

type StructUser struct {
	CurrentUser user.User
}

// type ContextKeyType string

const (
	Contectkey string = "userKey"
)

type authchecker struct {
	jwtService  auth.Service
	userService user.Service
}

type AuthChecker interface {
	AuthChecker(next echo.HandlerFunc) echo.HandlerFunc
}

func NewAutChecker(jwtService auth.Service, userService user.Service) AuthChecker {
	return &authchecker{jwtService, userService}
}

func (a *authchecker) AuthChecker(next echo.HandlerFunc) echo.HandlerFunc {
	// var jwtService auth.Service
	// var userService user.Service

	return func(c echo.Context) error {
		authVal := c.Request().Header.Get("Authorization")
		if !strings.Contains(authVal, "Bearer") {
			helper.PanicIfError(errors.New("error check bearer"), " error in check bearer")
		}

		arrayToken := strings.Split(authVal, " ")
		fmt.Println(arrayToken, "arrayToken")
		if len(arrayToken) != 2 {
			helper.PanicIfError(errors.New("error in spliting token "), "error split token")
		}
		jwtToken := arrayToken[1]
		log.Println(jwtToken, "token")
		tok, err := a.jwtService.ValidateToken(jwtToken)
		helper.PanicIfError(err, " error in validate token")
		log.Println(jwtToken, "token")
		if err != nil {
			log.Println(jwtToken, "inside nil token")

		}

		claim, ok := tok.Claims.(jwt.MapClaims)
		fmt.Println(claim, "claim", tok, "took", ok)
		if !ok || !tok.Valid {
			helper.PanicIfError(errors.New("erro validate token"), "error claim token")
		}
		user_id := claim["jti"]
		stringuser, ok := user_id.(string)
		fmt.Println(stringuser, "floatuser", ok, user_id)
		if !ok {
			helper.PanicIfError(errors.New("error in conver to int auth checker"), "error conver user id to int auth checker")
		}
		// intid, err := strconv.Atoi(stringuser)
		// helper.PanicIfError(err, " error in id string to int")
		// intuser := intid
		// fmt.Println(intuser)

		LoginUser, err := a.userService.GetUserById(c.Request().Context(), stringuser)
		helper.PanicIfError(err, " error in getting loigin user middleware")
		context_value := StructUser{
			CurrentUser: LoginUser,
		}
		c.Set(Contectkey, context_value)
		fmt.Println("from middleware two")
		return next(c)
	}
}
func PanicHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer exception.EchoRecover(c)
		return next(c)
	}
}
