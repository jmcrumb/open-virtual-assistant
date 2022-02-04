package nlp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ttsBody struct {
	Text string
}

func getTTS(c *gin.Context) {
	var body ttsBody

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, body)
}

type sttBody struct {
	Audio []byte
}

func getSTT(c *gin.Context) {
	var body sttBody

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, body)
}

func Route(router *gin.RouterGroup) {
	router.GET("/text2speech", getTTS)
	router.GET("/speech2text", getSTT)
}
