package Route

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//test := r.Group("/test")
	//{
	//	test.GET("/ping", testController.Ping)
	//}

	//test2 := r.Group("/test2")
	//test2.Use(Middleware.CORS(Middleware.CORSOptions{Origin: basicConfig.ServerCorsOrigin}))

	//r.GET("/redirect_test", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	//})

	return r
}
