package handler

import "nipun.io/otp-service/service"

type OTPHandler struct {
	otpService service.IOTPService
}

func NewOTPHandler(otpService service.IOTPService)*OTPHandler {
	return &OTPHandler{
		otpService: otpService,
	}
}