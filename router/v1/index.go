package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lihulab/xiaolongbaoproxy-mgt/controllers"
)

func NewRouter(g *gin.RouterGroup) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		v1.POST("/login", controllers.Login)
		v1.POST("/logout", controllers.Logout)
	}

	return router
}
