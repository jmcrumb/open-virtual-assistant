package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/auth"
	"github.com/jmcrumb/nova/database"
	"github.com/jmcrumb/nova/middleware"
)

func getAccountByID(c *gin.Context) {
	id := c.Param("id")

	auth.EnforceMiddlewareAuthentication(c, id, func(id string) {
		var account database.Account
		database.DB.Table("account").Where("id = ?", id).First(&account)
		if account.ID != "" {
			c.JSON(http.StatusOK, account)
			return
		}
		c.String(http.StatusNotFound, fmt.Sprintf("no account found with id: %v", id))
		return
	})
}

func getProfileByID(c *gin.Context) {
	id := c.Param("id")

	var profile database.Profile
	database.DB.Table("profile").Where("account_id = ?", id).First(&profile)
	if profile.AccountID != "" {
		var account database.Account
		database.DB.Table("account").Where("id = ?", id).First(&account)

		var pubProfile database.PublicProfile
		pubProfile.AccountID = profile.AccountID
		pubProfile.Bio = profile.Bio
		pubProfile.Photo = profile.Photo
		pubProfile.FirstName = account.FirstName
		pubProfile.LastName = account.LastName

		c.JSON(http.StatusOK, pubProfile)
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

	// TODO: protect against DB returning pre-existing account

	c.JSON(http.StatusCreated, accountResult)
}

func putAccount(c *gin.Context) {
	var body database.Account

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	auth.EnforceMiddlewareAuthentication(c, body.ID, func(id string) {
		err := database.DB.Table("account").Where("id = ?", id).Select("FirstName", "Email", "LastName").Updates(&body).Error
		if err != nil {
			c.String(http.StatusBadRequest, "invalid email provided")
			return
		}
		c.Status(http.StatusCreated)
		return
	})
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

	auth.EnforceMiddlewareAuthentication(c, body.AccountID, func(id string) {
		err := auth.ResetPassword(account, body.OldPassword, body.NewPassword)
		if err == nil {
			accountDB.Update("password", body.NewPassword)
			c.Status(http.StatusCreated)
			return
		}
		c.String(http.StatusBadRequest, "incorrect information")
	})
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
	router.GET("/:id", getAccountByID)
	router.POST("/", postAccount)
	router.PUT("/", middleware.AuthorizeJWT(), putAccount)
	router.PUT("/reset-password", middleware.AuthorizeJWT(), putAccountPassword)

	router.GET("/profile/:id", getProfileByID)
	router.PUT("/profile", putProfile)
}
