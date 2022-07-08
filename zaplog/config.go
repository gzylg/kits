package zaplog

import (
	"go.uber.org/zap/zapcore"
)

// getEncoder 编码器(如何写入日志)
type PrintType string

const (
	PrintTypeFmt  PrintType = "fmt"
	PrintTypeJson PrintType = "json"
)

var logLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

type Config struct {
	BothPrint         bool // 是否需要控制台与文件同时输出
	Colorful          bool // 是否彩色打印
	Caller            bool
	CallerSkip        int       // 如果Caller为true时，当前值有效
	DisableStacktrace bool      // 是否输出堆栈信息。 为true时关闭输出。 当输出等级为Error的日志时，同时输出堆栈信息
	LogLv             string    // 日志等级（debug、info、warn、error）
	PrintType         PrintType // 打印模式

	LogPath    string // 参数 logPath: 日志文件保存目录
	Logfile    string // 参数 logfile: 日志文件保存文件名
	MaxSize    int    // 参数 maxSize: 单个日志文件最大尺寸，单位：Mb
	MaxBackups int    // 参数 maxBackups: 日志备份文件最多数量
	MaxAge     int    // 参数 maxAge: 日志最长保留时间，单位: 天 (day)
}
