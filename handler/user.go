package handler

import (
	"bwastartupecho/auth"
	"bwastartupecho/exception"
	"bwastartupecho/helper"
	mid "bwastartupecho/middleware"
	"bwastartupecho/user"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type UserHandler interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
	CheckEmail(c echo.Context) error
	UploadAvatar(c echo.Context) error
	FetchUser(c echo.Context) error
}

type userHandler struct {
	userService user.Service
	authUser    auth.Service
}

func NewUserHandler(u user.Service, a auth.Service) UserHandler {
	return &userHandler{u, a}
}

// func NewUserHandler(u user.Service, a auth.Service) *userHandler {
// 	return &userHandler{u, a}
// }
// func deleteUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }

func (h *userHandler) RegisterUser(c echo.Context) error {
	var input user.RegisterUser
	// c.Request().se\\

	// helper.ReadFromRequestBody(request, &input)
	err := c.Bind(&input)
	helper.PanicIfError(err, "error in binding register user input hanlser")
	ctx := c.Request().Context()
	// con, erdr := context.WithCancel(ctx)
	// ctx := request.Context()
	inputUser, err := h.userService.RegisterUser(ctx, input)
	helper.PanicIfError(err, "error in register user handler")

	token, err := h.authUser.GenerateJWTToken(inputUser.ID)
	helper.PanicIfError(err, " erro in create token resgister user")
	sqlToken := sql.NullString{
		String: token,
		Valid:  true,
	}
	inputUser.Token = sqlToken

	userF := user.Formatuser(inputUser, sqlToken.String)
	res := helper.APIResponse("suscess", http.StatusOK, "ok", userF)
	// helper.WriteToResponseBody(writer, res)
	return c.JSON(http.StatusOK, res)
}

func (h *userHandler) LoginUser(c echo.Context) error {
	var input user.LoginInput

	// helper.ReadFromRequestBody(request, &input)
	c.Bind(&input)
	userLogin, err := h.userService.LoginUser(c.Request().Context(), input)
	helper.PanicIfError(err, " errrpor in login handler")
	token, err := h.authUser.GenerateJWTToken(userLogin.ID)
	helper.PanicIfError(err, " erro in create token resgister user")
	sqlToken := sql.NullString{
		String: token,
		Valid:  true,
	}
	userLogin.Token = sqlToken

	userF := user.Formatuser(userLogin, sqlToken.String)
	res := helper.APIResponse("suscess", http.StatusOK, "ok", userF)
	// helper.WriteToResponseBody(writer, res)/
	return c.JSON(http.StatusOK, res)

}

func (h *userHandler) CheckEmail(c echo.Context) error {
	var input user.CheckEmailInput
	// helper.ReadFromRequestBody(r, &input)
	c.Bind(&input)
	b, err := h.userService.CheckEmailAvailable(c.Request().Context(), input)
	exception.PanicIfNotFound(err, "error in finding emai")
	res := helper.APIResponse("suscess", http.StatusOK, "ok", b)
	// helper.WriteToResponseBody(w, res)
	return c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c echo.Context) error {

	// struct_ctx_intf := r.Context().Value(mid.Contectkey)
	struct_ctx_intf := c.Get(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	user_id := struct_context.CurrentUser.ID
	// file, fileHeader, err := c.Request().FormFile("file")
	fileHeader, err := c.FormFile("file")
	// fmt.Println(file, fileHeader, err)
	helper.PanicIfError(err, " error in get file upload")

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	path := fmt.Sprintf("images/%s-%s", user_id, fileHeader.Filename)
	fileDestinatin, err := os.Create(path)
	helper.PanicIfError(err, " error in create path destination")
	defer fileDestinatin.Close()

	fileSize, err := io.Copy(fileDestinatin, file)
	helper.PanicIfError(err, " errror in copy file into destinatio")
	fmt.Println(fileSize, " this is file size")

	userAvatar, err := h.userService.SaveAvatar(c.Request().Context(), user_id, fileDestinatin.Name())
	helper.PanicIfError(err, "error in save avater handler")
	log.Println(userAvatar)

	data := make(map[string]bool)
	data["is_uploaded"] = true
	respponse := helper.APIResponse("sucess", http.StatusOK, "success", data)

	return c.JSON(http.StatusOK, respponse)
	// helper.WriteToResponseBody(w, respponse)
}
func (h *userHandler) FetchUser(c echo.Context) error {
	// id := 2
	struct_ctx_intf := c.Get(mid.Contectkey)
	// struct_ctx_intf := r.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	id := struct_context.CurrentUser.ID
	us, err := h.userService.GetUserById(c.Request().Context(), id)
	helper.PanicIfError(err, " error in get current user handler")
	formatedUser := user.Formatuser(us, "")
	res := helper.APIResponse("susccess", http.StatusOK, "success", formatedUser)
	// helper.WriteToResponseBody(w, res)
	return c.JSON(http.StatusOK, res)
}

// func (h *userHandler) RegisterUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	var input user.RegisterUser
// 	helper.ReadFromRequestBody(request, &input)
// 	userRegistered, err := h.userService.RegisterUser(request.Context(), input)
// 	helper.PanicIfError(err, " erroor register user handler")
// 	jwtToken, err := h.authUser.GenerateJWTToken(userRegistered.ID)
// 	helper.PanicIfError(err, "error generate token")

// 	val := sql.NullString{
// 		String: jwtToken,
// 		Valid:  true,
// 	}
// 	userRegistered.Token = val
// 	formatedUser := user.Formatuser(userRegistered, val.String)
// 	response := helper.APIResponse("success register user", http.StatusOK, "success", formatedUser)
// 	helper.WriteToResponseBody(writer, response)

// }

// func (h *userHandler) LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

// 	var input user.LoginInput
// 	helper.ReadFromRequestBody(request, &input)
// 	userRegistered, err := h.userService.LoginUser(request.Context(), input)
// 	helper.PanicIfError(err, "error in login handler")

// 	jwtToken, err := h.authUser.GenerateJWTToken(userRegistered.ID)
// 	helper.PanicIfError(err, "error generate token")

// 	val := sql.NullString{
// 		String: jwtToken,
// 		Valid:  true,
// 	}
// 	userRegistered.Token = val
// 	formatedUser := user.Formatuser(userRegistered, val.String)
// 	repsonse := helper.APIResponse("success login", http.StatusAccepted, "success", formatedUser)
// 	helper.WriteToResponseBody(writer, repsonse)
// 	// writer.Write(repsonse)

// }

// func (h *userHandler) CheckEmail(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
// 	var input user.CheckEmailInput
// 	helper.ReadFromRequestBody(r, &input)
// 	available, err := h.userService.CheckEmailAvailable(r.Context(), input)
// 	helper.PanicIfError(err, "error handler check email")
// 	data := map[string]bool{"is_available": available}

// 	respnse := helper.APIResponse("email already registered", http.StatusOK, "success", data)
// 	helper.WriteToResponseBody(w, respnse)

// }

// func (h *userHandler) UploadAvatar(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
// 	fmt.Println(r.Context(), "cocococooc", r.Context().Value(middleware.Contectkey))
// 	struct_context_interface := r.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in upload avatar handler")
// 	}

// 	// userid := r.Context().Value(middleware.Contectkey)
// 	// fmt.Println(userid, "user id")
// 	// user_id, ok := userid.(int)

// 	// if !ok {
// 	// 	helper.PanicIfError(errors.New("error in convert user id"), "error in conver user id in upload avatar handler")
// 	// }

// 	user_id := struct_context.User_id
// 	// user_id := int(user_id_f)
// 	// request.ParseMultipartForm(32 << 20)
// 	file, fileHeader, err := r.FormFile("file")
// 	if err != nil {
// 		panic(err)
// 	}
// 	path := fmt.Sprintf("images/%d-%s", user_id, fileHeader.Filename)
// 	// fileDestination, err := os.Create("./images/" + strconv.Itoa(user_id) + "-" + fileHeader.Filename)
// 	fileDestination, err := os.Create(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = io.Copy(fileDestination, file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(fileDestination, fileDestination.Name(), fileHeader.Filename, "mudang wae")
// 	// user, err := h.userService.SaveAvatar(r.Context(), user_id, fileDestination.Name())
// 	_, err = h.userService.SaveAvatar(r.Context(), user_id, fileDestination.Name())
// 	if err != nil {
// 		panic(err)
// 	}
// 	data := make(map[string]bool)
// 	data["is_uploaded"] = true
// 	fmt.Println(fileDestination, fileDestination.Name(), fileHeader.Filename, "mudang wae")
// 	respnse := helper.APIResponse("upload image success", http.StatusOK, "success", data)
// 	helper.WriteToResponseBody(w, respnse)
// }

// func (h *userHandler) FetchUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
// 	struct_context_interface := r.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in upload avatar handler")
// 	}

// 	// userid := r.Context().Value(middleware.Contectkey)
// 	// fmt.Println(userid, "user id")
// 	// user_id, ok := userid.(int)

// 	// if !ok {
// 	// 	helper.PanicIfError(errors.New("error in convert user id"), "error in conver user id in upload avatar handler")
// 	// }

// 	currentUser := struct_context.CurrentUser

// 	formatter := user.Formatuser(currentUser, "")

// 	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)
// 	helper.WriteToResponseBody(w, response)
// }

// // func (h *userHandler) FetchUser(c *gin.Context) {

// // 	currentUser := c.MustGet("curresntuser").(user.User)

// // 	formatter := user.Formatuser(currentUser, "")

// // 	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

// // 	c.JSON(http.StatusOK, response)

// // }
