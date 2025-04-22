package repo

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type OTPRepo struct {
	redisClient *redis.Client
}

func NewOTPRepo(client *redis.Client) *OTPRepo {
	return &OTPRepo{
		redisClient: client,
	}
}

func (r *OTPRepo) SaveOTP(ctx context.Context, identifier string, otp string, expiration time.Duration) error {
	// Use the Set command with expiration (equivalent to SETEX)
	statusCmd := r.redisClient.Set(identifier, otp, expiration)
	return statusCmd.Err() // Return any error encountered during the operation
}

func (r *OTPRepo) GetOTP(ctx context.Context, identifier string) (string, error) {
	stringCmd := r.redisClient.Get(identifier)
	otp, err := stringCmd.Result()

	if err == redis.Nil {
		// Key does not exist (OTP expired or never set)
		// Return redis.Nil so the caller can specifically check for "not found"
		return "", redis.Nil
	} else if err != nil {
		// Another Redis error occurred
		return "", err
	}

	// OTP found successfully
	return otp, nil
}

func (r *OTPRepo) DeleteOTP(ctx context.Context, identifier string) error {
	intCmd := r.redisClient.Del(identifier)
	// We usually just care if the command succeeded, not how many keys were deleted (should be 0 or 1).
	return intCmd.Err()
}
