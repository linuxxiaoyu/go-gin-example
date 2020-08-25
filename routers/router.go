package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"
)

// Init new a gin engine, then return it.
func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	return r
}
