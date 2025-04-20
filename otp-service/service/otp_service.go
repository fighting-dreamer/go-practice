package service

import (
	"context"
	"math/rand"

	"strconv"

	"github.com/hashicorp/go-uuid"
	"nipun.io/otp-service/repo"
)

type UserContext struct {
	// this holds info that helps the user notifier
	// to send on right address, contact number or channel.

}

type OTPOptions struct {
	OTPType    string
	Length     int
	TTLSeconds int
}

type IOTPService interface {
	generateOTP(ctx context.Context, userInfo UserContext, optOptions OTPOptions) (string, string, error)
	validateOTP(ctx context.Context, requestID string, otp string) (bool, error)
}

type OTPService struct {
	db *repo.OTPRepo
}

func NewOTPService(db *repo.OTPRepo) IOTPService {
	return &OTPService{
		db: db,
	}
}

func randomOTP(OTPOptions OTPOptions) string {
	// for now it is constant 6 char long, numeric type, with no ttl as 1 minute => 60 seconds
	otp := int64(100000 + rand.Float32()*100000)
	return strconv.FormatInt(otp, 10)
}

func (o *OTPService) generateOTP(ctx context.Context, userInfo UserContext, otpOptions OTPOptions) (string, string, error) {
	otp := randomOTP(otpOptions)
	requestID, _ := uuid.GenerateUUID()
	err := o.db.SaveOTP(ctx, requestID, otp, 60) // otpOptions.TTLSeconds = 60 seconds default
	if err != nil {
		return requestID, "", err
	}
	return requestID, otp, nil
}

func (o *OTPService) validateOTP(ctx context.Context, requestID string, otp string) (bool, error) {
	otp, err := o.db.GetOTP(ctx, requestID)
	if err != nil {
		return false, err
	}
	return true, nil
}
