package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"github.com/jutionck/golang-todo-apps/utils/service"
)

type LogMiddleware interface {
	Logger() gin.HandlerFunc
}

type logMiddleware struct {
	loggerService service.LoggerService
}

func (l *logMiddleware) Logger() gin.HandlerFunc {
	err := l.loggerService.InitialLoggerFile()
	if err != nil {
		panic(err)
	}

	return func(ctx *gin.Context) {
		responseWriter := model.NewResponseLog(ctx.Writer)
		requestLog := model.RequestLog{
			Method:     ctx.Request.Method,
			StatusCode: ctx.Writer.Status(),
			ClientIP:   ctx.ClientIP(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		ctx.Writer = responseWriter
		ctx.Next()

		responseLog := model.ResponseLog{
			StatusCode:   responseWriter.Status(),
			ResponseBody: responseWriter.Body(),
		}

		switch {
		case ctx.Writer.Status() >= 400:
			l.loggerService.ReqLogError(requestLog)
			l.loggerService.ResLogError(responseLog)
		default:
			l.loggerService.ReqLogInfo(requestLog)
			l.loggerService.ResLogInfo(responseLog)
		}
	}
}

func NewLogMiddleware(loggerService service.LoggerService) LogMiddleware {
	return &logMiddleware{loggerService: loggerService}
}
