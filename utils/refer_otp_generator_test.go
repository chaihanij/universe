package utils_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/analytics-pumpchang/universe/utils"
	"testing"
)

func TestReferOtpCodeGenerate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		referOtp := utils.ReferOtpCodeGenerate()
		assert.NotNil(t, referOtp)
	})
}
