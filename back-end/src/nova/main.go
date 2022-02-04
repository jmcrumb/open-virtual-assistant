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

	router.SetTrustedProxies([]string{"localhost"})
	router.RunTLS(":443", "server.crt", "server.key")
	// log.Fatal(autotls.Run(router)) // use this instead of lines above when LetsEncrypt is configured
}
