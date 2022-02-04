package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/nlp"
)

func main() {
	router := gin.Default()

	// route NLP api
	group := router.Group("/nlp")
	nlp.Route(group)

	// route account database

	// route plugin store api

	router.Run()
}
