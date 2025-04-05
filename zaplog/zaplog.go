// Package zaplog
// 在本包内创建 zap 的日志实例，以控制台输出为主，文件同步输出为可选。
// 外部调用直接使用 zaplog.Info() 、 zaplog.Erro() 等方法
package zaplog

import "go.uber.org/zap"

var log *zap.SugaredLogger

// CreateLogger 在本包内创建 zap 的日志实例，外部调用直接使用 zaplog.Info() 、 zaplog.Erro() 等方法
func CreateLogger(cfg *Config) (err error) {
	log, err = NewLogger(cfg)
	return
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Sync() {
	log.Sync()
}

func IsCreated() bool {
	return log != nil
}
