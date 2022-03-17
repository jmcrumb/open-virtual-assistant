package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	jWtService JWTService
}

func LoginHandler(jWtService JWTService) LoginController {
	return &loginController{
		jWtService: jWtService,
	}
}

func AuthenticateUser(credential *LoginCredentials) (bool, database.Account) {
	var account database.Account

	database.DB.Table("account").Where("email = ?", credential.Email).First(&account)
	credential.Id = account.ID
	return (account.ID != "" && ComparePassword(account, credential.Password)), account
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated, account := AuthenticateUser(&credential)
	if isUserAuthenticated {
		database.DB.Table("account").Where("id = ?", account.ID).UpdateColumn("last_login", time.Now())
		return controller.jWtService.GenerateToken(account.ID)
	}
	return ""
}
