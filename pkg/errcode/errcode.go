package errcode

import (
	"fmt"
	"net/http"
)

// Error 统一错误模型
type Error struct {
	code    int      `json:"code"`    // 错误码
	msg     string   `json:"msg"`     // 错误信息
	details []string `json:"details"` // 错误详情
}

// codes 建立错误码和错误信息的映射关系
var codes = map[int]string{}

// NewError 创建 Error 统一错误实例
func NewError(code int, msg string) *Error {
	// 判断 codes 中是否含有 code 信息
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	// 如果不存在 code 信息
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

// GetError 获取错误信息
func (e *Error) GetError() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.GetCode(), e.GetMessage())
}

// GetCode 获取错误码
func (e *Error) GetCode() int {
	return e.code
}

// GetMessage 获取错误消息
func (e *Error) GetMessage() string {
	return e.msg
}

// MessagePrintf 格式化错误信息
func (e *Error) MessagePrintf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

// GetDetails 获取错误详情
func (e *Error) GetDetails() []string {
	return e.details
}

// WithDetails 添加额外的错误信息
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(details, d)
	}
	return e
}

// StatusCode 通过错码获取 HTTP 状态码
func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK // 200 OK
	case ServerError.GetCode():
		return http.StatusInternalServerError // 500 服务器内部错误
	case InvalidParams.GetCode():
		return http.StatusBadRequest // 400 错误的请求
	case UnauthorizedAuthNotExist.GetCode():
		fallthrough
	case UnauthorizedTokenError.GetCode():
		fallthrough
	case UnauthorizedTokenGenerate.GetCode():
		fallthrough
	case UnauthorizedTokenTimeout.GetCode():
		return http.StatusUnauthorized // 401 没有权限
	case TooManyRequests.GetCode():
		return http.StatusTooManyRequests // 429 请求太多
	}
	// 500 服务器内部错误
	return http.StatusInternalServerError
}
