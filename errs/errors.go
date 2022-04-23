package errs

import (
	"fmt"
	"runtime"
)

type YsErr struct {
	code     int     // 错误代码
	msg      string  // 错误信息
	funcName uintptr // 错误处函数名
	file     string  // 错误处文件名
	line     int     // 错误所在行
}

func New(msg string) error {
	e := &YsErr{msg: msg}
	if funcName, file, line, ok := runtime.Caller(1); ok {
		e.funcName = funcName
		e.file = file
		e.line = line
	}

	return e
}

func NewWithErr(prefixStr string, err error) error {
	e := &YsErr{msg: prefixStr + err.Error()}
	if funcName, file, line, ok := runtime.Caller(1); ok {
		e.funcName = funcName
		e.file = file
		e.line = line
	}

	return e
}

// As 判断一个error是否为此处自定义实现的err
func AsYsErr(e error) bool {
	_, ok := e.(*YsErr)
	return ok
}

func (e *YsErr) Error() string {
	return e.msg
}

func (e *YsErr) Errorf() string {
	return fmt.Sprintf(
		"msg:%s, func:%s:%v, file:%s",
		e.msg,
		runtime.FuncForPC(e.funcName).Name(), e.line,
		e.file,
	)
}
