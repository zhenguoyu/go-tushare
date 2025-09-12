package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestCli(t *testing.T) {
	// 测试代码
	err := errors.New("error--error")
	err2 := fmt.Errorf("error: %w", err)
	err3 := fmt.Errorf("error2: %w", err2)
	t.Log("errs3", err3)
}
