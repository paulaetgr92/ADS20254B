package middleware

import (
	"crypto/rand"
	"math/big"
	"os"
)

const charsetNumber = "0123456789"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetSignatureString() string {
	return os.Getenv("TOKEN_SIGNATURE")
}

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsetNumber))))
		if err != nil {
			return "", err
		}
		result[i] = charsetNumber[num.Int64()]
	}
	return string(result), nil
}

func GenerateRandomNumberContractString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}
