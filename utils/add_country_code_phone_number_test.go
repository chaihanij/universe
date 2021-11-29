package utils_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"universe/utils"
)

func TestUtils_AddCountryCodePhoneNumber(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		countryCode := "TH"
		phoneNumberArray := []string{"0891234567", "+66891234567", "891234567", "089-1234567", "089 123 4567", "089-123 4567", "0 8 9 1 2 3 4 5 6 7", "0  8  9 1  2 3 4 5 6 7"}
		expectedPhoneNumber := "+66891234567"

		for _, v := range phoneNumberArray {
			phoneNumber, err := utils.AddCountryCodePhoneNumber(countryCode, v)
			assert.NotNil(t, phoneNumber)
			assert.Nil(t, err)
			assert.Equal(t, expectedPhoneNumber, *phoneNumber)
		}
	})

	t.Run("Error parameters", func(t *testing.T) {
		countryCode := ""
		phoneNumberArray := []string{""}

		for _, v := range phoneNumberArray {
			phoneNumber, err := utils.AddCountryCodePhoneNumber(countryCode, v)
			assert.Nil(t, phoneNumber)
			assert.NotNil(t, err)
		}
	})

	t.Run("Error countryCode", func(t *testing.T) {
		countryCode := "AA"
		phoneNumberArray := []string{"0891234567", "891234567"}

		for _, v := range phoneNumberArray {
			phoneNumber, err := utils.AddCountryCodePhoneNumber(countryCode, v)
			assert.Nil(t, phoneNumber)
			assert.NotNil(t, err)
		}
	})

	t.Run("Error phone", func(t *testing.T) {
		countryCode := "TH"
		phoneNumberArray := []string{"test", "test+0123", "089+0123", "0891234567+test"}

		for _, v := range phoneNumberArray {
			phoneNumber, err := utils.AddCountryCodePhoneNumber(countryCode, v)
			assert.Nil(t, phoneNumber)
			assert.NotNil(t, err)
		}
	})
}
