package app

import (
	"github.com/chen2eric/blog-service/global"
	"github.com/chen2eric/blog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	page_size := convert.StrTo(c.Query("page_size")).MustInt()
	if page_size <= 0 {
		return global.AppSetting.DefaultPageSize
	} else if page_size > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return page_size
}

func GetPageOffset(page, page_size int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * page_size
	}
	return result
}
