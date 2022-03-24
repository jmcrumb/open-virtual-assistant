package plugins

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

func getPlugins(c *gin.Context) {
	c.String(http.StatusBadGateway, "this endpoint is not yet implemented")
}
func postPlugin(c *gin.Context) {
	var body database.NewPlugin
	var pluginResult database.Plugin

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	if err := database.DB.Table("plugin").Create(&body).Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid publisher id provided: %q", body.Publisher))
		return
	}
	// this might lead to some duplicate plugins while only one is ever updated / used
	database.DB.Table("plugin").Where("source_link = ?", body.SourceLink).First(&pluginResult)

	c.JSON(http.StatusCreated, body)
}
func putPlugin(c *gin.Context) {
	var body database.Plugin

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	database.DB.Table("plugin").Where("id = ?", body.ID).Select("SourceLink", "About").Updates(&body)
	c.Status(http.StatusCreated)
}
func deletePlugin(c *gin.Context) {
	id := c.Param("id")

	var plugin database.Plugin
	database.DB.Table("plugin").Where("id = ?", id).Delete(&plugin)

	c.Status(http.StatusNoContent)
}
func getPlugin(c *gin.Context) {
	c.String(http.StatusBadGateway, "this endpoint is not yet implemented")
}
func Route(router *gin.RouterGroup) {
	router.GET("/", getPlugins)
	router.POST("/", postPlugin)
	router.PUT("/", putPlugin)
	router.DELETE("/:id", deletePlugin)
	router.GET("/:id", getPlugin)
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