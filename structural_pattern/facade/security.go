package facade

import (
	"errors"
	"fmt"
)

// 复杂子系统的组成部分
type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
