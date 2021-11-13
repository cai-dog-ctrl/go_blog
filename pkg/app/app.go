package app

import (
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// ToErrorResponse 传入错误信息并响应给前端
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}  //将错误码和错误信息写入响应体中
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)   //响应错误
}
