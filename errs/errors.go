package errs

import (
	"fmt"
	"runtime"
)

type YsErr struct {
	Code int    // 错误代码
	Msg  string // 错误信息
	Func string // 错误处函数名+行号
	File string // 错误处文件名
}

func (e *YsErr) caller() {
	if funcName, file, line, ok := runtime.Caller(2); ok {
		e.Func = fmt.Sprintf("%s:%d", runtime.FuncForPC(funcName).Name(), line)
		e.File = file
	}
}

func New(msg string) error {
	e := &YsErr{Code: -1, Msg: msg}
	e.caller()
	return e
}

func NewWithErr(prefixStr string, err error) error {
	e := &YsErr{Code: -1, Msg: prefixStr + err.Error()}
	e.caller()
	return e
}

func NewWithCode(code ErrType) error {
	e := &YsErr{Code: int(code), Msg: msgList[code]}
	e.caller()
	return e
}

// As 判断一个error是否为此处自定义实现的err
func AsYsErr(e error) bool {
	_, ok := e.(*YsErr)
	return ok
}

// IsYsErr 判断错误内容是否相同
func IsYsErr(e1 error, e2 error) bool {
	return e1.Error() == e2.Error()
}

// IsYsErr 判断错误内容是否相同，判断e1的msg，是否为传入的errType msg相同
func IsYsErrWithErrType(e1 error, errType ErrType) bool {
	msg := GetErrMessage(errType)
	return msg == e1.Error()
}

// ISYsErrWithStrict 判断两个错误内容是否相同
// 先通过 AsYsErr 检查，通过后判断两个错误内容的code和msg是否相同
func IsYsErrWithStrict(e1 error, e2 error) bool {
	if !AsYsErr(e1) || !AsYsErr(e2) {
		return false
	}
	return e1.(*YsErr).Code == e2.(*YsErr).Code && e1.(*YsErr).Msg == e2.(*YsErr).Msg
}

func (e *YsErr) Error() string {
	return e.Msg
}

func (e *YsErr) Errorf() string {
	return fmt.Sprintf(
		"code:%d, msg:%s, func:%s, file:%s",
		e.Code,
		e.Msg,
		e.Func,
		e.File,
	)
}
