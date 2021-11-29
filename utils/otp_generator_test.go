package utils_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"universe/utils"
)

func TestGenerateOTPCode(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		otp, err := utils.GenerateOTPCode()

		assert.Nil(t, err)
		assert.NotNil(t, otp)
	})
}
