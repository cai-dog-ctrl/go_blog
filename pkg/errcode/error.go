package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code int `json:"code"`  //错误码
	msg string `json:"msg"`  //错误信息
	detail []string `json:"detail"`  //错误细节
}
var codes =map[int]string{}

func NewError(code int,msg string)*Error{
	if _,ok:=codes[code];ok{
		panic(fmt.Sprintf("错误码：%d已经存在，请跟换一个",code))
	}
	codes[code]=msg
	return &Error{code:code,msg: msg}
}

func (e *Error)Error()string{
	return fmt.Sprintf("错误码：%d，错误信息：%s",e.code,e.msg)
}

func (e *Error)Code()int{
	return e.code
}

func (e *Error)Msg()string{
	return e.msg
}

func (e *Error)Msgf(args []interface{})string{
	return fmt.Sprintf(e.msg,args...)
}

func (e *Error)Details()[]string{
	return e.detail
}

// WithDetails 将错误的详细信息写入到Error体中
func (e *Error)WithDetails(details ...string)*Error{
	newError:=*e
	newError.detail=details
	for _,v:=range details{
		newError.detail=append(newError.detail,v)
	}
	return &newError
}
var a int

// StatusCode 错误处理公共方法
func (e *Error)StatusCode()int{
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
