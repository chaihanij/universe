package utils_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/analytics-pumpchang/universe/utils"
	"testing"
)

func TestUtils_GetWeakPin6Digit(t *testing.T) {
	t.Run("Happy GetWeakPin6Digit", func(t *testing.T) {
		result := utils.GetWeakPin6Digit()

		assert.NotNil(t, result)
	})
}
