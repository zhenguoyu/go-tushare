package gotushare

import (
	"fmt"
)

type Error struct {
	Code    int
	Message string
	Op      string
}

func (e Error) Error() string {
	return fmt.Sprintf("tushare error (code=%d, op=%s): %s", e.Code, e.Op, e.Message)
}

func NewError(code int, message, op string) error {
	return Error{
		Code:    code,
		Message: message,
		Op:      op,
	}
}
