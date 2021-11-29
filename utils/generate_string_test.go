package utils_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/analytics-pumpchang/universe/utils"
	"testing"
)

func TestUtils_GenerateRandomString(t *testing.T) {
	t.Run("Happy GenerateString", func(t *testing.T) {
		expectLength := 5
		result := utils.GenerateRandomString(expectLength)
		assert.NotNil(t, result)
		assert.Equal(t, expectLength, len(result))
	})
}
