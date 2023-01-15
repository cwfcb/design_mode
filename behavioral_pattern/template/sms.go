package template

import "fmt"

// Sms golang 使用组合的方法实现继承，实现具体的步骤
type Sms struct {
	*Otp
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending Sms: %s\n", message)
	return nil
}

func (s *Sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}
