package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type LoginService interface {
// 	LoginUser(email string, password string) bool
// }
// type loginInformation struct {
// 	email    string
// 	password string
// }

//Login credential
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Id       string
}

// func StaticLoginService() LoginService {
// 	return &loginInformation{
// 		email:    "jcrumb@sandiego.edu",
// 		password: "usdnova",
// 	}
// }

func Route(router *gin.RouterGroup) {
	// var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	// var loginController LoginController = LoginHandler(loginService, jwtService)
	var loginController LoginController = LoginHandler(jwtService)

	router.POST("/login", func(ctx *gin.Context) {
		token, accountId := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token":      token,
				"account_id": accountId,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}
