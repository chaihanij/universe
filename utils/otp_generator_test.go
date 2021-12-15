package utils_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/gridwhizth/universe/utils"
	"testing"
)

func TestGenerateOTPCode(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		otp, err := utils.GenerateOTPCode()

		assert.Nil(t, err)
		assert.NotNil(t, otp)
	})
}
