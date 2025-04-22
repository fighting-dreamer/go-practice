package repo

import (
	"context"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2" // Alias to avoid conflict if needed

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (*OTPRepo, *miniredis.Miniredis, func()) {
	t.Helper() // Marks this function as a test helper

	mr, err := miniredis.Run()
	require.NoError(t, err, "Failed to start miniredis")

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(), // Connect to the miniredis instance
	})

	// Optional: Ping to ensure connection is okay, though miniredis usually just works
	_, err = client.Ping().Result()
	require.NoError(t, err, "Failed to ping miniredis")

	repo := NewOTPRepo(client)

	// Cleanup function to close miniredis and the client connection
	cleanup := func() {
		// It's good practice to close the client, though often not strictly
		// necessary with miniredis as closing the server disconnects clients.
		_ = client.Close() // Ignore error on close for simplicity in test
		mr.Close()
	}

	return repo, mr, cleanup
}

func TestOTPRepo_SaveOTP(t *testing.T) {
	repo, mr, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	identifier := "user:123:email"
	otp := "987654"
	expiration := 5 * time.Minute

	err := repo.SaveOTP(ctx, identifier, otp, expiration)
	assert.NoError(t, err, "SaveOTP should not return an error on success")

	// Verify directly with miniredis
	savedOTP, err := mr.Get(identifier)
	assert.NoError(t, err, "miniredis should have the key after SaveOTP")
	assert.Equal(t, otp, savedOTP, "Saved OTP in redis does not match")

	// Verify TTL (Time To Live) is set correctly
	ttl := mr.TTL(identifier)
	// Allow a small delta for processing time
	assert.InDelta(t, expiration.Seconds(), ttl.Seconds(), 1, "TTL is not set correctly")
}

func TestOTPRepo_GetOTP(t *testing.T) {
	repo, mr, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	identifierExists := "user:123:email"
	otpExists := "987654"
	otpExpired := "123456"
	identifierNonExistent := "user:nonexistent12345"
	identifierExpired := "user:expired:12345"
	expiration := time.Minute

	// Case 1 : OTP exists
	t.Run("OTP exists", func(t *testing.T) {
		err := repo.SaveOTP(ctx, identifierExists, otpExists, expiration)
		assert.NoError(t, err, "No error in saving OTP")

		retrievedOTP, err := repo.GetOTP(ctx, identifierExists)
		assert.NoError(t, err, "No error in retrieving OTP")
		assert.Equal(t, otpExists, retrievedOTP, "Retrieved OTP does not match")
	})
	// Case 2 : OTP non existent
	t.Run("OTP non existent", func(t *testing.T) {
		retrievedOTP, err := repo.GetOTP(ctx, identifierNonExistent)
		assert.Error(t, err, "Expected error for non-existent OTP")
		assert.Equal(t, redis.Nil, err, "Error should be redis.Nil for non-existent OTP")
		assert.Empty(t, retrievedOTP, "Retrieved OTP should be empty for non-existent OTP")
	})
	// Case 3 : OTP expired
	t.Run("OTP expired", func(t *testing.T) {
		shortExpiration := 10 * time.Millisecond
		err := repo.SaveOTP(ctx, identifierExpired, otpExpired, shortExpiration)
		assert.NoError(t, err, "No error in saving expired OTP")
		mr.FastForward(shortExpiration + time.Millisecond)

		retrievedOTP, err := repo.GetOTP(ctx, identifierExpired)
		assert.Error(t, err, "Expected error for expired OTP")
		assert.Equal(t, redis.Nil, err, "Error should be redis.Nil for expired OTP")
		assert.Empty(t, retrievedOTP, "Retrieved OTP should be empty for expired OTP")
	})
}
