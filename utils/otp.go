package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() string {
	const digits = "0123456789"
	otp := make([]byte, 6)

	for i := range otp {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		otp[i] = digits[n.Int64()]
	}
	return string(otp)
}
