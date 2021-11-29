package utils_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"universe/utils"
)

func TestUtils_GetWeakPin6Digit(t *testing.T) {
	t.Run("Happy GetWeakPin6Digit", func(t *testing.T) {
		result := utils.GetWeakPin6Digit()

		assert.NotNil(t, result)
	})
}
