// Package zaplog
// 创建 zap 的日志实例，以控制台输出为主，文件同步输出为可选。
// 支持文件输出(分割文件)+控制台输出、输出颜色、输出类型（fmt/json）

package zaplog

import (
	"os"

	"github.com/gzylg/kits/errs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger 创建一个日志zap.SugaredLogger，并支持同时输出到文件
func NewLogger(cfg *Config) (logger *zap.SugaredLogger, err error) {
	encoder := getEncode(cfg.Colorful, cfg.PrintType) // 获取日志输出编码
	if _, ok := logLevel[cfg.LogLv]; !ok {
		return nil, errs.New("zaplog：请配置正确的日志等级")
	}

	//* 根据 BothPrint 设置 writeSyncer
	var writeSyncer zapcore.WriteSyncer
	if cfg.BothPrint {
		// 同时输出
		writeSyncer, err = getLogWriter( // 日志文件配置 文件位置和切割
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
	} else {
		// 仅控制台输出
		writeSyncer = zapcore.AddSync(os.Stdout)
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel[cfg.LogLv])

	zapOpt := []zap.Option{}
	if !cfg.DisableStacktrace {
		zapOpt = append(zapOpt, zap.AddStacktrace(zapcore.ErrorLevel))
	}
	if cfg.Caller {
		zapOpt = append(zapOpt, zap.AddCaller(), zap.AddCallerSkip(cfg.CallerSkip))
	}

	var l *zap.Logger
	if len(zapOpt) > 0 {
		l = zap.New(core, zapOpt...)
	} else {
		l = zap.New(core)
	}

	return l.Sugar(), nil
}
