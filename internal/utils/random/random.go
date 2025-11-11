package random

import "math/rand"

// GenerateRandomOTP generates a 4-digit OTP code.
func GenerateOTPWithLength(length int) string {
	const chars = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
