package routers

import (
	"github.com/linuxxiaoyu/go-gin-example/middleware/jwt"
	"github.com/linuxxiaoyu/go-gin-example/routers/api"
	v1 "github.com/linuxxiaoyu/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"
)

// Init new a gin engine, then return it.
func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// tag
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// asticle
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
