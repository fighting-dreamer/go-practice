package service

type IOTPService interface{

}

type OTPService struct {

}

func NewOTPService() IOTPService {
	return &OTPService{}
}