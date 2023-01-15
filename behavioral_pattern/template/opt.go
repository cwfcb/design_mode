package template

import (
	"math/rand"
	"time"
)

type IOtp interface {
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type Otp struct {
	iOtp IOtp
}

// 定义一个「算法」框架，具体算法由 IOtp 实现
func (o *Otp) GenAndSendOTP(optLength int) error {
	otp := o.genRandomOTP(optLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

func (o *Otp) genRandomOTP(otpLength int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterLength := len(letters)
	randomRune := make([]rune, otpLength)
	for i := 0; i < otpLength; i++ {
		randomRune[i] = letters[rand.Intn(letterLength)]
	}
	return string(randomRune)
}
