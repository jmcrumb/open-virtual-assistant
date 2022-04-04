package reviews

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

func getReviews(c *gin.Context) {
	pluginID := c.Param("plugin")
	reviews := []database.Review{}

	// check plugin existance
	var plugin database.Plugin
	database.DB.Table("plugin").Where("id = ?", pluginID).First(&plugin)
	if plugin.ID == "" {
		c.String(http.StatusBadRequest, "invalid plugin ID")
		return
	}

	database.DB.Table("review").Where("plugin = ?", pluginID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}
func postReview(c *gin.Context) {
	var body database.NewReview
	var reviewResult database.Review

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	if err := database.DB.Table("review").Select("Account", "Plugin", "Rating", "Content").Create(&body).Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid account or plugin id provided: {account: %q, plugin: %q}", body.Account, body.Plugin))
		return
	}

	// this might lead to some duplicate reviews while only one is ever updated / used
	database.DB.Table("review").Where("plugin = ? AND account = ? AND content = ?", body.Plugin, body.Account, body.Content).First(&reviewResult)
	c.JSON(http.StatusCreated, reviewResult)
}
func putReview(c *gin.Context) {
	var body database.Review

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	database.DB.Table("review").Where("id = ?", body.ID).Select("Rating", "Content").Updates(&body)
	c.Status(http.StatusCreated)
}
func deleteReview(c *gin.Context) {
	id := c.Param("id")

	var review database.Review
	database.DB.Table("review").Where("id = ?", id).Delete(&review)

	c.Status(http.StatusNoContent)
}
func getReview(c *gin.Context) {
	id := c.Param("id")

	var review database.Review
	database.DB.Table("review").Where("id = ?", id).First(&review)

	if review.ID == "" {
		c.String(http.StatusBadRequest, "invalid review ID")
		return
	}

	c.JSON(http.StatusOK, review)
}

func Route(router *gin.RouterGroup) {
	router.GET("/:plugin", getReviews)
	router.POST("/", postReview)
	router.PUT("/", putReview)
	router.DELETE("/:plugin/:id", deleteReview)
	router.GET("/:plugin/:id", getReview)
}
