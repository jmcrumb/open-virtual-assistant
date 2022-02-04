package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/nlp"
)

type User struct {
	Name string
}

func main() {
	router := gin.Default()

	// route NLP api
	group := router.Group("/nlp")
	nlp.Route(group)

	// route account database

	// route plugin store api

	// m := autocert.Manager{
	// 	Prompt: autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("localhost"), // absolutely need to use this once CA is used
	// }

	router.SetTrustedProxies([]string{"localhost"})
	router.RunTLS(":443", "server.crt", "server.key")
	// log.Fatal(autotls.RunWithManager(router, &m))
}
