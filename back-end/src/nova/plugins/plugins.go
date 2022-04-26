package plugins

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/auth"
	"github.com/jmcrumb/nova/database"
	"github.com/jmcrumb/nova/middleware"
)

func postPlugin(c *gin.Context) {
	var body database.NewPlugin
	var pluginResult database.Plugin

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	auth.EnforceMiddlewareAuthentication(c, body.Publisher, func(id string) {
		if err := database.DB.Table("plugin").Create(&body).Error; err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("invalid publisher id provided: %q", body.Publisher))
			return
		}
		// this might lead to some duplicate plugins while only one is ever updated / used
		database.DB.Table("plugin").Where("source_link = ?", body.SourceLink).First(&pluginResult)

		c.JSON(http.StatusCreated, body)
	})
}

func putPlugin(c *gin.Context) {
	var body database.Plugin

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	auth.EnforceMiddlewareAuthentication(c, body.ID, func(id string) {
		database.DB.Table("plugin").Where("id = ?", body.ID).Select("SourceLink", "About").Updates(&body)
		c.Status(http.StatusCreated)
	})
}
func deletePlugin(c *gin.Context) {
	id := c.Param("id")

	auth.EnforceMiddlewareAuthentication(c, id, func(id string) {
		var plugin database.Plugin
		database.DB.Table("plugin").Where("id = ?", id).Delete(&plugin)

		c.Status(http.StatusNoContent)
	})
}

func getPlugin(c *gin.Context) {
	id := c.Param("id")

	var plugin database.Plugin
	database.DB.Table("plugin").Where("id = ?", id).First(&plugin)

	if plugin.ID == "" {
		c.String(http.StatusBadRequest, "invalid plugin ID")
		return
	}

	c.JSON(http.StatusOK, plugin)
}

func searchPlugin(c *gin.Context) {
	query := c.Param("query")

	var plugins []database.Plugin
	database.DB.Raw("SELECT * FROM plugin WHERE tsv_name @@ to_tsquery(?);", query).Find(&plugins)

	c.JSON(http.StatusOK, plugins)
}
func Route(router *gin.RouterGroup) {
	router.POST("/", middleware.AuthorizeJWT(), postPlugin)
	router.PUT("/", middleware.AuthorizeJWT(), putPlugin)
	router.DELETE("/:id", middleware.AuthorizeJWT(), deletePlugin)
	router.GET("/:id", getPlugin)
	router.GET("/search/:query", searchPlugin)
}
func postPluginReview(c *gin.Context) {
	var body database.NewReview

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarhsall request body")
		return
	}

	database.DB.Table("review").Create(&body)

	c.Status(http.StatusCreated)
}
func postPluginReport(c *gin.Context) {
	var body database.NewReport

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarhsall request body")
		return
	}

	database.DB.Table("report").Create(&body)

	c.Status(http.StatusCreated)
}