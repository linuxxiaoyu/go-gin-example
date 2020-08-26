package routers

import (
	v1 "github.com/linuxxiaoyu/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"
)

// Init new a gin engine, then return it.
func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
