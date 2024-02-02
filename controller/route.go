package controller

import (
	"bwastartupecho/auth"
	"bwastartupecho/handler"
	"bwastartupecho/middleware"
	"bwastartupecho/user"

	"github.com/labstack/echo"
)

// e.GET("/", func(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }, middleware.AuthChecker)

func NewUserHandler(r *echo.Echo, h handler.UserHandler, jwtService auth.Service, userService user.Service, midd middleware.AuthChecker) *echo.Echo {
	// midd := middleware.NewAutChecker(jwtService, userService)
	// r.Use(middleware.PanicHandler)
	// r.Handle()
	// dire := http.Dir("./images")
	// FileServer(r, "/static", dire)
	r.Static("/static", "images")
	r.POST("/api/v1/users", h.RegisterUser)
	r.POST("/api/v1/sessions", h.LoginUser)
	r.POST("/api/v1/email_checkers", h.CheckEmail)
	r.POST("/api/v1/avatars", h.UploadAvatar, midd.AuthChecker)
	// r.Post("/api/v1/avatares", h.UploadAvatar)
	r.POST("/api/v1/user/fetch", h.FetchUser)
	return r
}

// func FileServer(r chi.Router, path string, root http.FileSystem) {
// 	if strings.ContainsAny(path, "{}*") {
// 		panic("FileServer does not permit any URL parameters.")
// 	}
// 	fmt.Println("kakakasssssssk")

// 	if path != "/" && path[len(path)-1] != '/' {
// 		fmt.Println("pattt", path, len(path), path[len(path)-1])
// 		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
// 		fmt.Println("pattt", path, len(path), path[len(path)-1])
// 		path += "/"
// 		fmt.Println("pattt", path, len(path), path[len(path)-1])

// 	}
// 	path += "*"
// 	fmt.Println("kakakak")
// 	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
// 		rctx := chi.RouteContext(r.Context())
// 		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
// 		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
// 		fmt.Println(rctx, "ggg", pathPrefix, "ggg", fs, "lll", rctx.RoutePattern())
// 		fs.ServeHTTP(w, r)
// 		fmt.Println("after fs run ", fs)
// 		return
// 	})
// }

// func NewUserHandler(r *chi.Mux, h handler.UserHandler, jwtService auth.Service, userService user.Service, authModdleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {
// 	// r.Handle()
// 	dire := http.Dir("./images")
// 	FileServer(r, "/static", dire)

// 	r.Post("/api/v1/users", h.RegisterUser)
// 	r.Post("/api/v1/sessions", h.LoginUser)
// 	r.Post("/api/v1/email_checkers", h.CheckEmail)
// 	r.Post("/api/v1/avatars", authModdleware(jwtService, userService, h.UploadAvatar))
// 	// r.Post("/api/v1/avatares", h.UploadAvatar)
// 	r.Post("/api/v1/user/fetch", h.FetchUser)
// 	return r
// }

func NewCampaignHandler(r *echo.Echo, h handler.CampaignHandler, midd middleware.AuthChecker) *echo.Echo {
	// midd := middleware.NewAutChecker(jwtService, userService)
	// r.Use(middleware.PanicHandler)
	// r.Handle()
	// dire := http.Dir("./images")
	// FileServer(r, "/static", dire)
	// r.Static("/static", "assets")
	// r.POST("/api/v1/users", h.RegisterUser)
	// r.POST("/api/v1/sessions", h.LoginUser)
	// r.POST("/api/v1/email_checkers", h.CheckEmail)
	// r.POST("/api/v1/avatars", h.UploadAvatar, midd.AuthChecker)
	// // r.Post("/api/v1/avatares", h.UploadAvatar)
	// r.POST("/api/v1/user/fetch", h.FetchUser)
	// return r

	r.GET("/api/v1/campaigns", h.GetCampaigns)
	r.GET("/api/v1/campaigns/{campaignid}", h.GetCampaign)
	g := r.Group("")

	g.Use(midd.AuthChecker)

	g.POST("/api/v1/campaigns", h.CreateCampaign)
	g.PUT("/api/v1/campaigns/:campaignid", h.UpdateCampaign)
	g.POST("/api/v1/campaign-images", h.UploadCampaignImage)

	//// another way using midleware in echo
	// r.POST("/api/v1/campaigns", h.CreateCampaign, midd.AuthChecker)
	// r.PUT("/api/v1/campaigns/{campaignid}", h.UpdateCampaign, midd.AuthChecker)
	// r.POST("/api/v1/campaign-images", h.UploadCampaignImage, midd.AuthChecker)

	return r
}

// func NewCampaignHandler(r *chi.Mux, h handler.CampaignHandler, jwtService auth.Service, userService user.Service, authModdleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {
// 	// r.Handle()
// 	// dire := http.Dir("./images")
// 	// FileServer(r, "/static", dire)

// 	// r.Post("/api/v1/users", h.RegisterUser)
// 	// r.Post("/api/v1/sessions", h.LoginUser)
// 	// r.Post("/api/v1/email_checkers", h.CheckEmail)
// 	// r.Post("/api/v1/avatars", authModdleware(jwtService, userService, h.UploadAvatar))
// 	// // r.Post("/api/v1/avatares", h.UploadAvatar)
// 	// r.Post("/api/v1/user/fetch", h.FetchUser)

// 	r.Get("/api/v1/campaigns", h.GetCampaigns)
// 	r.Post("/api/v1/campaigns", authModdleware(jwtService, userService, h.CreateCampaign))
// 	r.Put("/api/v1/campaigns/{campaignid}", authModdleware(jwtService, userService, h.UpdateCampaign))
// 	r.Get("/api/v1/campaigns/{campaignid}", h.GetCampaign)
// 	r.Post("/api/v1/campaign-images", authModdleware(jwtService, userService, h.UploadCampaignImage))
// 	return r
// }

func NewTransactionHandler(r *echo.Echo, transactionHandler handler.TransactionHandler, midd middleware.AuthChecker) *echo.Echo {
	// midd := middleware.NewAutChecker(jwtService, userService)
	// r.Use(middleware.PanicHandler)
	// func NewTransactionHandler(router *httprouter.Router, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {

	r.POST("/api/v1/transactions/notification", transactionHandler.GetNotif)
	g := r.Group("")
	g.Use(midd.AuthChecker)
	g.GET("/api/v1/campaigns/:campaignid/transactions", transactionHandler.GetCampaignTransactions)
	g.GET("/api/v1/transactions", transactionHandler.GetUserTransactions)
	g.POST("/api/v1/transactions", transactionHandler.CreateTransaction)
	return r
}

// func NewTransactionHandler(r *chi.Mux, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {

// 	// func NewTransactionHandler(router *httprouter.Router, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {

// 	r.Get("/api/v1/campaigns/{campaignid}/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetCampaignTransactions))
// 	r.Get("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetUserTransactions))
// 	r.Post("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.CreateTransaction))
// 	r.Post("/api/v1/transactions/notification", transactionHandler.GetNotif)
// 	return r
// }
