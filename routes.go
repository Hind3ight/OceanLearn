package main

import (
	"github.com/Hind3ight/OceanLearn/controller"
	"github.com/Hind3ight/OceanLearn/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.JWTAuthMiddleware(), controller.Info)

	return r
}
