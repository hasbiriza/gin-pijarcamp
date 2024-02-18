package routes

import (
	"github.com/gin-gonic/gin"
)

func router() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		ping := v1.Group("ping")
		{
			ping.Get("/", contollers.PingController)
		}
	}
}
