package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func ReferOtpCodeGenerate() string {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(899999))
	return strconv.FormatInt(nBig.Int64()+99999, 10)
}
