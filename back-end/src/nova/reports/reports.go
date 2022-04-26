package reports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/auth"
	"github.com/jmcrumb/nova/database"
	"github.com/jmcrumb/nova/middleware"
)

func getReports(c *gin.Context) {
	auth.EnforceMiddlewareAuthenticatedAdmin(c, func() {
		pluginID := c.Param("plugin")
		reports := []database.Report{}

		// check plugin existance
		var plugin database.Plugin
		database.DB.Table("plugin").Where("id = ?", pluginID).First(&plugin)
		if plugin.ID == "" {
			c.String(http.StatusBadRequest, "invalid plugin ID")
			return
		}

		database.DB.Table("report").Where("plugin = ?", pluginID).Find(&reports)
		c.JSON(http.StatusOK, reports)
	})
}
func postReport(c *gin.Context) {
	var body database.NewReport
	var reportResult database.Report

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	auth.EnforceMiddlewareAuthentication(c, body.Account, func(id string) {
		if err := database.DB.Table("report").Select("Account", "Plugin", "Content").Create(&body).Error; err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("invalid account or plugin id provided: {account: %q, plugin: %q}", id, body.Plugin))
			return
		}

		// this might lead to some duplicate reports while only one is ever updated / used
		database.DB.Table("report").Where("plugin = ? AND account = ? AND content = ?", body.Plugin, body.Account, body.Content).First(&reportResult)
		c.JSON(http.StatusCreated, reportResult)
	})
}
func putReport(c *gin.Context) {
	var body database.Report

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "unable to unmarshall request body")
		return
	}

	auth.EnforceMiddlewareAuthentication(c, body.Account, func(id string) {
		database.DB.Table("report").Where("id = ?", body.ID).Select("Content").Updates(&body)
		c.Status(http.StatusCreated)
	})
}
func deleteReport(c *gin.Context) {
	id := c.Param("id")

	auth.EnforceMiddlewareAuthentication(c, id, func(id string) {
		var report database.Report
		database.DB.Table("report").Where("id = ?", id).Delete(&report)

		c.Status(http.StatusNoContent)
	})
}
func getReport(c *gin.Context) {
	auth.EnforceMiddlewareAuthenticatedAdmin(c, func() {
		id := c.Param("id")

		var report database.Report
		database.DB.Table("report").Where("id = ?", id).First(&report)

		if report.ID == "" {
			c.String(http.StatusBadRequest, "invalid report ID")
			return
		}

		c.JSON(http.StatusOK, report)
	})
}

func Route(router *gin.RouterGroup) {
	router.GET("/:plugin", middleware.AuthorizeJWT(), getReports)
	router.POST("/", middleware.AuthorizeJWT(), postReport)
	router.PUT("/", middleware.AuthorizeJWT(), putReport)
	router.DELETE("/:plugin/:id", middleware.AuthorizeJWT(), deleteReport)
	router.GET("/:plugin/:id", middleware.AuthorizeJWT(), getReport)
}
