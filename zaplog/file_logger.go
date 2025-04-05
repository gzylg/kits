// Package zaplog
// 创建 zap 的文件日志实例，以文件输出为主，控制台同步输出为可选。
// 支持文件输出(分割文件)+控制台输出、输出颜色、输出类型（fmt/json）

package zaplog

import (
	"os"
	"path/filepath"

	"github.com/gzylg/kits/errs"
	"github.com/gzylg/kits/file"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFileLogger 创建一个写出到文件的zap.SugaredLogger，并支持同时输出到控制台
func NewFileLogger(cfg *Config) (logger *zap.SugaredLogger, err error) {
	writeSyncer, err := getLogWriter( // 日志文件配置 文件位置和切割
		cfg.BothPrint,
		cfg.LogPath,
		cfg.Logfile,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)

	if err != nil {
		return nil, err
	}
	encoder := getEncode(cfg.Colorful, cfg.PrintType) // 获取日志输出编码

	if _, ok := logLevel[cfg.LogLv]; !ok {
		return nil, errs.New("zaplog：请配置正确的日志等级")
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel[cfg.LogLv])

	zapOpt := []zap.Option{}
	zapOpt = append(zapOpt, zap.AddStacktrace(zapcore.ErrorLevel))
	if cfg.Caller {
		zapOpt = append(zapOpt, zap.AddCaller())
	}
	l := zap.New(core, zapOpt...)

	return l.Sugar(), nil
}

func getEncode(colorful bool, printType PrintType) zapcore.Encoder {
	//* 配置 EncoderConfig
	encodeCfg := zap.NewProductionEncoderConfig()
	// 自定义时间格式
	encodeCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	if colorful {
		// 使用彩色输出
		encodeCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encodeCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	if printType == PrintTypeFmt {
		return zapcore.NewConsoleEncoder(encodeCfg) // 以logfmt格式写入
	}

	return zapcore.NewJSONEncoder(encodeCfg) // 以json格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(needStdPrint bool, logPath, logfile string, maxSize, maxBackups, maxAge int) (zapcore.WriteSyncer, error) {
	// 判断日志路径是否存在，如果不存在就创建
	if !file.Exists(logPath) {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, logfile), // 日志文件路径
		MaxSize:    maxSize,                         // 单个日志文件最大多少 mb
		MaxBackups: maxBackups,                      // 日志备份文件最多数量
		MaxAge:     maxAge,                          // 日志最长保留时间，单位: 天 (day)
		Compress:   true,                            // 是否压缩日志
	}

	if needStdPrint {
		// 文件、控制台 同时输出
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger)), nil
}
