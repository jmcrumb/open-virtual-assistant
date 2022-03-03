package plugins

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

func postPlugin(c *gin.Context) {

}
func putPlugin(c *gin.Context) {

}
func deletePlugin(c *gin.Context) {

}
func Route(router *gin.RouterGroup) {

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