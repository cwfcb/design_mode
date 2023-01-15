package template

import "fmt"

// Email golang 使用组合的方法实现继承，实现具体的步骤
type Email struct {
	*Otp
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending Email: %s\n", message)
	return nil
}

func (s *Email) publishMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}
