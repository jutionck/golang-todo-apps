package model

import "github.com/gin-gonic/gin"

type RequestLog struct {
	Method     string
	StatusCode int
	ClientIP   string
	Path       string
	UserAgent  string
}

type ResponseWriter struct {
	gin.ResponseWriter
	status int
	body   []byte
}

type ResponseLog struct {
	StatusCode   int
	ResponseBody string
}

func NewResponseLog(rw gin.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: rw}
}

func (r *ResponseWriter) Write(data []byte) (int, error) {
	r.body = append(r.body, data...)
	return r.ResponseWriter.Write(data)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *ResponseWriter) Body() string {
	return string(r.body)
}

func (r *ResponseWriter) Status() int {
	return r.status
}
