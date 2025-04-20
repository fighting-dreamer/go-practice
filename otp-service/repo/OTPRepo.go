package repo

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type OTPRepo struct {
	prefix      string
	redisClient *redis.Client
}

func NewOTPRepo(prefix string, client *redis.Client) *OTPRepo {
	return &OTPRepo{
		prefix:      prefix,
		redisClient: client,
	}
}

func (r *OTPRepo) SaveOTP(ctx context.Context, identifier string, otp string, expiration time.Duration) error {
	// Use the Set command with expiration (equivalent to SETEX)
	key := r.prefix + ":" + identifier
	statusCmd := r.redisClient.Set(key, otp, expiration)
	return statusCmd.Err() // Return any error encountered during the operation
}

func (r *OTPRepo) GetOTP(ctx context.Context, identifier string) (string, error) {
	key := r.prefix + ":" + identifier
	stringCmd := r.redisClient.Get(key)
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
	key := r.prefix + ":" + identifier
	intCmd := r.redisClient.Del(key)
	// We usually just care if the command succeeded, not how many keys were deleted (should be 0 or 1).
	return intCmd.Err()
}
