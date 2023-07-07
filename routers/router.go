package routers

import (
	"github.com/cindyyangcaixia/gin-example/middlewares"
	v1 "github.com/cindyyangcaixia/gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.POST("/schools", v1.CreateSchools)
		apiv1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}

	return r

}
