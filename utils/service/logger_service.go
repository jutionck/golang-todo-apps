package service

import (
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerService interface {
	InitialLoggerFile() error
	ReqLogInfo(requestLog model.RequestLog)
	ReqLogError(requestLog model.RequestLog)
	ResLogInfo(responseLog model.ResponseLog)
	ResLogError(responseLog model.ResponseLog)
}

type loggerUtil struct {
	cfg config.FileConfig
	log *logrus.Logger
}

func (l *loggerUtil) InitialLoggerFile() error {
	file, err := os.OpenFile(l.cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	} else {
		l.log = logrus.New()
		l.log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		l.log.Out = file
	}
	return nil
}

func (l *loggerUtil) ReqLogError(requestLog model.RequestLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Request Log",
	}).Error(requestLog)
}

func (l *loggerUtil) ReqLogInfo(requestLog model.RequestLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Request Log",
	}).Info(requestLog)
}

func (l *loggerUtil) ResLogError(responseLog model.ResponseLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Response Log",
	}).Error(responseLog)
}

func (l *loggerUtil) ResLogInfo(responseLog model.ResponseLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Response Log",
	}).Info(responseLog)
}

func NewLoggerService(cfg config.FileConfig) LoggerService {
	return &loggerUtil{cfg: cfg}
}
