package plugins

import (
	"github.com/gin-gonic/gin"
)

func postPlugin(c *gin.Context) {

}
func putPlugin(c *gin.Context) {

}
func deletePlugin(c *gin.Context) {

}

func Route(router *gin.RouterGroup) {
	router.POST("/plugins/:name", postPlugin)
	router.PUT("/plugins/:name", putPlugin)
	router.DELETE("/plugins/:name", deletePlugin)
}
