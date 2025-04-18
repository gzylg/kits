package zlog_test

import (
	"testing"

	"github.com/gzylg/kits/zlog"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestZlog(t *testing.T) {
	_ = zlog.NewLogger(&zlog.LogConfig{
		Level:      zapcore.DebugLevel,
		Color:      true,
		ShowCaller: false,

		FileLogWarnConfig: &lumberjack.Logger{
			Filename:   "log.log",
			MaxSize:    10,
			MaxBackups: 100,
			MaxAge:     28,
			Compress:   false,
		},
		FileLogErrorConfig: &lumberjack.Logger{
			Filename:   "error.log",
			MaxSize:    10,
			MaxBackups: 100,
			MaxAge:     28,
			Compress:   false,
		},
	})
	defer zlog.Sync()

	zlog.Debug("test debug")
	zlog.Debugf("test debugf %s", "test")

	zlog.Info("test info")
	zlog.Infof("test infof %s", "test")

	zlog.Warn("test warn")
	zlog.Warnf("test warnf %s", "test")

	zlog.Error("test error")
	zlog.Errorf("test errorf %s", "test")
}
