package errs

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("TestNew")

	t.Log(err.(*YsErr).Errorf())
}

func TestNewWithCode(t *testing.T) {
	err := NewWithCode(ERRTYPE_USER_EMAIL_IS_EXIST)

	t.Log(err.(*YsErr).Errorf())
}

func TestNewWithErr(t *testing.T) {
	err := errors.New("TestNewWithErr")
	err = NewWithErr("包装一下：", err)

	t.Log(err.(*YsErr).Errorf())
}

func TestAsYsErr(t *testing.T) {
	err := errors.New("TestAsYsErr")
	t.Log(AsYsErr(err))

	err = NewWithErr("包装一下：", err)
	t.Log(AsYsErr(err))

	err = New("TestAsYsErr")
	t.Log(AsYsErr(err))

	err = NewWithCode(ERRTYPE_USER_EMAIL_IS_EXIST)
	t.Log(AsYsErr(err))
}

func TestIsYsErr(t *testing.T) {
	e1 := errors.New("TestIsYsErr")
	e2 := New("TestIsYsErr")
	t.Log(IsYsErr(e1, e2))
}

func TestIsYsErrWithStrict(t *testing.T) {
	e1 := errors.New("TestIsYsErr")
	e2 := New("TestIsYsErr")
	t.Log(IsYsErrWithStrict(e1, e2))

	e1 = New("TestIsYsErr")
	e2 = New("TestIsYsErr")
	t.Log(IsYsErrWithStrict(e1, e2))
}
