package routers

import (
	v1 "github.com/cindyyangcaixia/ApplicationOfNEMT/controller/v1"
	_ "github.com/cindyyangcaixia/ApplicationOfNEMT/docs"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	controller := v1.NewController()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// docs.SwaggerInfo.Title = "Application of National College Entrance Examination"
	// docs.SwaggerInfo.Description = "This application is for college candidates to fill in their volunteers as a reference"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.BasePath = "/api/v1"

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.POST("/schools", controller.CreateSchool)
		apiv1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
