package routers

import (
	v1 "github.com/cindyyangcaixia/gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/schools", v1.GetSchools)
	}

	return r

}
