package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/accounts"
	"github.com/jmcrumb/nova/database"
	"github.com/jmcrumb/nova/nlp"
)

func main() {
	err := database.SetupDB()
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	// route NLP api
	nlpGroup := router.Group("/nlp")
	nlp.Route(nlpGroup)

	// route account database
	accountGroup := router.Group("/account")
	accounts.Route(accountGroup)

	// route plugin store api

	router.SetTrustedProxies([]string{"localhost"})
	router.RunTLS(":443", "server.crt", "server.key")
	// log.Fatal(autotls.Run(router)) // use this instead of lines above when LetsEncrypt is configured
}
