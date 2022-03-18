package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/auth"
	"github.com/jmcrumb/nova/database"
)

func getAccountByID(c *gin.Context) {
	id := c.Param("id")

	// if id != auth.GetMiddlewareAuthenticatedAccountID(c) {
	// 	c.String(http.StatusUnauthorized, "Permission Denied")
	// 	return
	// }

	var account database.Account
	database.DB.Table("account").Where("id = ?", id).First(&account)
	if account.ID != "" {
		c.JSON(http.StatusOK, account)
		return
	}

	c.String(http.StatusNotFound, fmt.Sprintf("no account found with id: %v", id))
}

func getProfileByID(c *gin.Context) {
	id := c.Param("id")

	var profile database.Profile
	database.DB.Table("profile").Where("account_id = ?", id).First(&profile)
	if profile.AccountID != "" {
		c.JSON(http.StatusOK, profile)
		return
	}

	c.String(http.StatusNotFound, fmt.Sprintf("no profile found with id: %v", id))
}

func postAccount(c *gin.Context) {
	var body database.NewAccount
	var accountResult database.Account

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	// Secure password
	body.Password = auth.HashPassword(body.Password)

	// do a regex match on the email to make sure it's valid
	database.DB.Table("account").Create(&body)
	database.DB.Table("account").Where("email = ?", body.Email).First(&accountResult)
	database.DB.Table("profile").Create(&database.Profile{
		AccountID: accountResult.ID,
	})

	c.JSON(http.StatusCreated, accountResult)
}

func putAccount(c *gin.Context) {
	var body database.Account

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	err := database.DB.Table("account").Where("id = ?", body.ID).Select("FirstName", "Email", "LastName").Updates(&body).Error
	if err != nil {
		c.String(http.StatusBadRequest, "invalid email provided")
		return
	}
	c.Status(http.StatusCreated)
}

func putAccountPassword(c *gin.Context) {
	var body database.UpdatePassword
	var account database.Account

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	accountDB := database.DB.Table("account").Where("id = ?", body.AccountID)
	if err := accountDB.First(&account).Error; err != nil {
		c.String(http.StatusBadRequest, "incorrect information")
		return
	}

	err := auth.ResetPassword(account, body.OldPassword, body.NewPassword)
	if err == nil {
		accountDB.Update("password", body.NewPassword)
		c.Status(http.StatusCreated)
		return
	}
	c.String(http.StatusBadRequest, "incorrect information")
}

func putProfile(c *gin.Context) {
	var body database.Profile

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	database.DB.Table("profile").Where("account_id = ?", body.AccountID).Select("Bio", "Photo").Updates(&body)
	c.Status(http.StatusCreated)
}

func Route(router *gin.RouterGroup) {
	// router.GET("/:id", middleware.AuthorizeJWT(), getAccountByID)
	router.GET("/:id", getAccountByID)
	router.POST("/", postAccount)
	router.PUT("/", putAccount)
	router.PUT("/reset-password", putAccountPassword)

	router.GET("/profiles/:id", getProfileByID)
	router.PUT("/profiles", putProfile)
}
