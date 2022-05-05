package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/accounts"
	"github.com/jmcrumb/nova/auth"
	"github.com/jmcrumb/nova/database"
	"github.com/jmcrumb/nova/middleware"
	"github.com/jmcrumb/nova/plugins"
	"github.com/jmcrumb/nova/reports"
	"github.com/jmcrumb/nova/reviews"
)

func main() {
	database.SetupDB()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// route authorization API
	authGroup := router.Group("/auth")
	auth.Route(authGroup)

	// route account database
	accountGroup := router.Group("/account")
	accounts.Route(accountGroup)

	// route plugin store api
	pluginGroup := router.Group("/plugin")
	plugins.Route(pluginGroup)

	reviewGroup := router.Group("/review")
	reviews.Route(reviewGroup)

	// route plugin store api
	reportGroup := router.Group("/report")
	reports.Route(reportGroup)

	router.SetTrustedProxies([]string{"localhost", "127.0.0.1"})
	router.RunTLS(":443", "server.crt", "server.key")
	// log.Fatal(autotls.Run(router)) // use this instead of lines above when LetsEncrypt is configured
}
