package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang_demo/pkg/setting"
	v1 "golang_demo/routers/api/v1"
	"golang_demo/routers/api/v1/client"
)

func InitRouter() *gin.Engine {
	routes := gin.New()

	routes.Use(gin.Logger())
	routes.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := routes.Group("/v1")
	{
		apiv1.GET("/member/:member_id", v1.GetMember)
		apiv1.GET("/members/:page", v1.GetMembers)
		apiv1.POST("/member", v1.AddMember)
		apiv1.PUT("/member/:member_id", v1.EditMember)
		apiv1.DELETE("/member/:member_id", v1.DeleteMember)
	}

	apiv1.GET("/test/:message", client.Hello)
	routes.GET("/ws", client.Ws)
	routes.GET("/", client.Html)

	return routes
}
