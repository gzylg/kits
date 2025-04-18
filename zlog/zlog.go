package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	Level      zapcore.Level `json:"level"`      // 日志级别
	Color      bool          `json:"color"`      // 是否彩色
	ShowCaller bool          `json:"showCaller"` // 是否显示调用者

	// 文件日志相关配置
	FileLogWarnConfig  *lumberjack.Logger // <=warn级别日志文件配置，为nil时不启用；Level高于WarnLevel时，<=warn级别文件日志自动关闭
	FileLogErrorConfig *lumberjack.Logger // >=error级别日志文件配置，为nil时不启用
}

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func NewLogger(cfg *LogConfig) *zap.SugaredLogger {
	var cores []zapcore.Core

	// 如果Level高于WarnLevel，则关闭<=warn级别文件日志
	if cfg.Level > zapcore.WarnLevel {
		cfg.FileLogWarnConfig = nil
	}

	// 初始化Console Core
	cores = append(cores, zapcore.NewCore(newConsoleEncoder(cfg), zapcore.AddSync(os.Stdout), zap.DebugLevel))

	// 初始化 <=warn 级别 Core
	warnFileWriterSyncer := newWarnFileWriterSyncer(cfg)
	warnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool { return l <= zap.WarnLevel })
	if warnFileWriterSyncer != nil {
		cores = append(cores, zapcore.NewCore(newJsonEncoder(), warnFileWriterSyncer, warnLevel))
	}

	// 初始化 >=error 级别 Core
	errorFileWriterSyncer := newErrorFileWriterSyncer(cfg)
	errorLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool { return l >= zap.ErrorLevel })
	if errorFileWriterSyncer != nil {
		cores = append(cores, zapcore.NewCore(newJsonEncoder(), errorFileWriterSyncer, errorLevel))
	}

	// 初始化Logger
	if cfg.ShowCaller {
		logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		logger = zap.New(zapcore.NewTee(cores...))
	}

	sugar = logger.Sugar()
	return sugar
}

func Sync() {
	_ = sugar.Sync()
	_ = logger.Sync()
}

// newConsoleEncoder 创建一个Console编码器
func newConsoleEncoder(cfg *LogConfig) zapcore.Encoder {
	zapConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg", // 关键字段‌:ml-citation{ref="4" data="citationList"}
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), //指定时间格式
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// 设置日志级别编码器是否彩色
	if cfg.Color {
		zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		zapConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	return zapcore.NewConsoleEncoder(zapConfig)
}

// newJsonEncoder 创建一个json编码器
func newJsonEncoder() zapcore.Encoder {
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return zapcore.NewJSONEncoder(zapConfig)
}

func newWarnFileWriterSyncer(cfg *LogConfig) zapcore.WriteSyncer {
	if cfg.FileLogWarnConfig == nil || cfg.FileLogWarnConfig.Filename == "" {
		return nil
	}

	// 创建 <=warn 级别日志文件配置
	warnLumberIO := cfg.FileLogWarnConfig
	if warnLumberIO.MaxSize == 0 {
		warnLumberIO.MaxSize = 10
	}
	if warnLumberIO.MaxBackups == 0 {
		warnLumberIO.MaxBackups = 100
	}
	if warnLumberIO.MaxAge == 0 {
		warnLumberIO.MaxAge = 10
	}

	return zapcore.AddSync(warnLumberIO)
}

func newErrorFileWriterSyncer(cfg *LogConfig) zapcore.WriteSyncer {
	if cfg.FileLogErrorConfig == nil || cfg.FileLogErrorConfig.Filename == "" {
		return nil
	}

	// 创建 >=error 级别日志文件配置
	errorLumberIO := cfg.FileLogErrorConfig
	if errorLumberIO.MaxSize == 0 {
		errorLumberIO.MaxSize = 10
	}
	if errorLumberIO.MaxBackups == 0 {
		errorLumberIO.MaxBackups = 100
	}
	if errorLumberIO.MaxAge == 0 {
		errorLumberIO.MaxAge = 10
	}

	return zapcore.AddSync(errorLumberIO)
}

// ---------------------------------- 输出封装 ----------------------------------

func Debug(msg string, fields ...zap.Field) { logger.Debug(msg, fields...) }

func Info(msg string, fields ...zap.Field) { logger.Info(msg, fields...) }

func Warn(msg string, fields ...zap.Field) { logger.Warn(msg, fields...) }

func Error(msg string, fields ...zap.Field) { logger.Error(msg, fields...) }

func DPanic(msg string, fields ...zap.Field) { logger.DPanic(msg, fields...) }

func Panic(msg string, fields ...zap.Field) { logger.Panic(msg, fields...) }

func Fatal(msg string, fields ...zap.Field) { logger.Fatal(msg, fields...) }

func Debugf(template string, args ...any) { sugar.Debugf(template, args...) }

func Infof(template string, args ...any) { sugar.Infof(template, args...) }

func Warnf(template string, args ...any) { sugar.Warnf(template, args...) }

func Errorf(template string, args ...any) { sugar.Errorf(template, args...) }

func DPanicf(template string, args ...any) { sugar.DPanicf(template, args...) }

func Panicf(template string, args ...any) { sugar.Panicf(template, args...) }

func Fatalf(template string, args ...any) { sugar.Fatalf(template, args...) }
