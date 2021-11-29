package utils_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/analytics-pumpchang/universe/utils"
	"testing"
)

func TestUtils_CensorPhoneNumber(t *testing.T) {
	t.Run("Happy replaceDigit is 0", func(t *testing.T) {
		phoneNumber := "0891234567"
		expectedPhoneNumber := "0891234567"
		phoneNumberCensor := utils.CensorPhoneNumber(phoneNumber, 0)

		assert.NotNil(t, phoneNumberCensor)
		assert.Equal(t, expectedPhoneNumber, phoneNumberCensor)
	})

	t.Run("Happy", func(t *testing.T) {
		phoneNumber := "0891234567"
		expectedPhoneNumber := "0891234XXX"
		phoneNumberCensor := utils.CensorPhoneNumber(phoneNumber, 3)

		assert.NotNil(t, phoneNumberCensor)
		assert.Equal(t, expectedPhoneNumber, phoneNumberCensor)
	})

	t.Run("Happy replaceDigit More Than PhoneNumber", func(t *testing.T) {
		phoneNumber := "0891234567"
		expectedPhoneNumber := "XXXXXXXXXX"
		phoneNumberCensor := utils.CensorPhoneNumber(phoneNumber, 99)

		assert.NotNil(t, phoneNumberCensor)
		assert.Equal(t, expectedPhoneNumber, phoneNumberCensor)
	})

	t.Run("Happy replaceDigit Less Than 0", func(t *testing.T) {
		phoneNumber := "0891234567"
		expectedPhoneNumber := "0891234567"
		phoneNumberCensor := utils.CensorPhoneNumber(phoneNumber, -1)

		assert.NotNil(t, phoneNumberCensor)
		assert.Equal(t, expectedPhoneNumber, phoneNumberCensor)
	})
}
