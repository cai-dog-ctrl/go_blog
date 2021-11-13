// Package app 分页
package app

import (
	"blog/global"
	"blog/pkg/convert"
	"github.com/gin-gonic/gin"
)

// GetPage 返回页数
func GetPage(c *gin.Context)int{
	page:=convert.StrTo(c.Query("Page")).MustInt()
	if page<=0{
		return 1
	}
	return page
}

// GetPageSize 返回分页大小
func GetPageSize(c *gin.Context)int{
	pageSize:=convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize<=0{
		return global.AppSetting.DefaultPageSize
	}
	if pageSize >global.AppSetting.MaxPageSize{
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

// GeyPageOffset 返回分页偏移量
func GeyPageOffset (page ,pageSize int)int {
	result := 0
	if page >0{
		result=(page-1)*pageSize
	}
	return result
}


