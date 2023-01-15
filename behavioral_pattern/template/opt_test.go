package template

import "testing"

func TestOtp(t *testing.T) {
	sms := &Sms{}
	email := &Email{}

	otp := &Otp{}
	/*
		golang 不支持抽象类；
		所以不能直接使用子类，要将子类注入到父类中；
		父类中将接口作为其 field
	*/
	otp.iOtp = sms
	err := otp.GenAndSendOTP(4)
	if err != nil {
		t.Fatal(err)
	}

	otp.iOtp = email
	err = otp.GenAndSendOTP(4)
	if err != nil {
		t.Fatal(err)
	}
}
