package util

import (
	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Page(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
