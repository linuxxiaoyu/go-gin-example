package routers

import (
	"net/http"

	_ "github.com/linuxxiaoyu/go-gin-example/docs"
	"github.com/linuxxiaoyu/go-gin-example/middleware/jwt"
	"github.com/linuxxiaoyu/go-gin-example/pkg/export"
	"github.com/linuxxiaoyu/go-gin-example/pkg/upload"
	"github.com/linuxxiaoyu/go-gin-example/routers/api"
	v1 "github.com/linuxxiaoyu/go-gin-example/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"
)

// Init new a gin engine, then return it.
func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// tag
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)

		// asticle
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
