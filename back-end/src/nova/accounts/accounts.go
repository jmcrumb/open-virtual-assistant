package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

func getAccountByID(c *gin.Context) {
	id := c.Param("id")

	var account database.Account
	database.DB.Table("account").Where("id = ?", id).First(&account)
	if account.ID != "" {
		c.JSON(http.StatusOK, account)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "account not found"})
}

func getProfileByID(c *gin.Context) {
	id := c.Param("id")

	var profile database.Profile
	database.DB.Table("account").Where("id = ?", id).First(&profile)
	if profile.AccountID != "" {
		c.JSON(http.StatusOK, profile)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "profile not found"})
}

func postAccount(c *gin.Context) {
	var body database.NewAccount
	var accountResult database.Account

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarhsall request body")
		return
	}

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
		c.String(http.StatusBadRequest, "unable to unmarhsall request body")
		return
	}

	database.DB.Table("account").Where("id = ?", body.ID).Select("FirstName", "Email", "LastName", "Password").Updates(&body)
	c.Status(http.StatusCreated)
}

func putProfile(c *gin.Context) {
	var body database.Profile

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarhsall request body")
		return
	}

	database.DB.Table("profle").Where("account_id = ?", body.AccountID).Select("Bio", "Photo").Updates(&body)
	c.Status(http.StatusCreated)
}

func Route(router *gin.RouterGroup) {
	router.GET("/:id", getAccountByID)
	router.POST("/", postAccount)
	router.PUT("/", putAccount)

	router.GET("/profile/:id", getProfileByID)
	router.PUT("/profile", putProfile)
}
