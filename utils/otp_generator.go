package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateOTPCode() (string, error) {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(8999))
	return strconv.FormatInt(nBig.Int64()+999, 10), nil
}
