package reviews

import "github.com/gin-gonic/gin"

func getReviews(c *gin.Context) {

}
func postReview(c *gin.Context) {

}
func putReview(c *gin.Context) {

}
func deleteReview(c *gin.Context) {

}
func getReview(c *gin.Context) {

}

func Route(router *gin.RouterGroup) {
	router.GET("/:plugin")
	router.POST("/:plugin")
	router.PUT("/:plugin")
	router.DELETE("/:plugin/:id")
	router.GET("/:plugin/:id")
}
